package stream

import (
	"encoding/json"
	"fmt"
	"schwab-client-go/internal"
	"schwab-client-go/internal/stream/model"
	"schwab-client-go/util"
	"strings"
	"time"

	tmodel "schwab-client-go/internal/trader/model"

	"github.com/sacOO7/gowebsocket"
)

type TDStream struct {
	stop             bool
	user             tmodel.UserPreference
	socket           *gowebsocket.Socket
	isConnected      bool
	isLoggingEnabled bool

	requestidindex   int64
	requestqueue     map[string]TDStreamOnResponseFunc
	requestqueuesent map[string]model.TDWSRequest

	OnResponse   func(message string)
	OnError      func(err error)
	OnCheck      func(message string)
	OnConnect    func()
	OnDisconnect func()
}

type TDStreamOnResponseFunc struct {
	onResponse func(message string)
}

func (t *TDStream) EnableLogging() {
	t.isLoggingEnabled = true
}

func (t *TDStream) Init(userpref tmodel.UserPreference) {
	t.requestidindex = 0
	t.requestqueue = make(map[string]TDStreamOnResponseFunc)
	t.requestqueuesent = make(map[string]model.TDWSRequest)
	t.user = userpref

	t.login()
}

func (t *TDStream) Dispose() {
	t.isConnected = false
	t.stop = true
	logoutreq := model.TDWSRequest{
		Service:   "ADMIN",
		Command:   "LOGOUT",
		Requestid: t.nextRequestID(),
	}

	var o model.TDWSResponse[model.TDWSResponse_General]
	t.getCommandResponse(logoutreq, &o)
	if t.isLoggingEnabled {
		util.Log(fmt.Sprintf("OnLogout:%s", util.Serialize(o)))
	}
}

func (t *TDStream) nextRequestID() string {
	t.requestidindex = t.requestidindex + int64(1)
	return fmt.Sprintf("%d", t.requestidindex)
}

func (t *TDStream) login() {
	tmpsocket := gowebsocket.New(t.user.StreamerInfo[0].StreamerSocketUrl)
	tmpsocket.Connect()
	t.socket = &tmpsocket
	/*
		if t.isLoggingEnabled {
			t.socket.EnableLogging()
		}
	*/

	t.socket.OnTextMessage = t.onResponse
	t.socket.OnBinaryMessage = func(message []byte, socket gowebsocket.Socket) {
		t.onResponse(string(message), socket)
	}
	t.socket.OnConnectError = func(err error, socket gowebsocket.Socket) {
		if t.isLoggingEnabled {
			util.Log(fmt.Sprintf("OnConnectError:%s", err.Error()))
		}
		if t.OnError != nil {
			t.OnError(err)
		}
	}
	t.socket.OnConnected = func(socket gowebsocket.Socket) {
		if t.isLoggingEnabled {
			util.Log("stream connection success")
		}
		if t.OnConnect != nil {
			t.OnConnect()
		}
	}
	t.socket.OnPingReceived = func(data string, socket gowebsocket.Socket) {
		if t.OnCheck != nil {
			t.OnCheck(data)
		}
	}
	t.socket.OnPongReceived = func(data string, socket gowebsocket.Socket) {
		if t.OnCheck != nil {
			t.OnCheck(data)
		}
	}
	t.socket.OnDisconnected = func(err error, socket gowebsocket.Socket) {
		if t.isConnected {
			t.isConnected = false
			t.login()
			return
		}

		if err != nil {
			util.OnError(err)
		}
		if t.isLoggingEnabled {
			util.Log("stream disconnected")
		}
		for _, v := range t.requestqueue {
			v.onResponse(`{"response":[{"content":{"code":22,"msg":"stream disconnected"}}]}`)
		}
		if t.OnDisconnect != nil {
			t.OnDisconnect()
		}
	}

	loginreq := model.TDWSRequest{
		Service: "ADMIN",
		Command: "LOGIN",
		Parameters: map[string]string{
			"Authorization":          internal.Token.GetToken(),
			"SchwabClientChannel":    t.user.StreamerInfo[0].SchwabClientChannel,
			"SchwabClientFunctionId": t.user.StreamerInfo[0].SchwabClientFunctionId,
		},
	}

	if t.isLoggingEnabled {
		util.Log(fmt.Sprintf("connecting to %s with stream", t.user.StreamerInfo[0].StreamerSocketUrl))
	}
	var o model.TDWSResponse[model.TDWSResponse_General]
	t.getCommandResponse(loginreq, &o)
	if t.isLoggingEnabled {
		util.Log(fmt.Sprintf("OnLogin:%s", util.Serialize(o)))
	}
	t.isConnected = true
}

func (t *TDStream) sendCommand(request model.TDWSRequest, onresponse func(message string)) {
	request.Requestid = t.nextRequestID()
	callbackid := request.Requestid
	t.requestqueue[callbackid] = TDStreamOnResponseFunc{
		onResponse: onresponse,
	}
	t.requestqueuesent[callbackid] = request

	if request.Command == "SUBS" || request.Command == "ADD" {
		for _, k := range strings.Split(request.Parameters["keys"], ",") {
			callbackid = fmt.Sprintf("%s:%s", request.Service, k)
			t.requestqueue[callbackid] = TDStreamOnResponseFunc{
				onResponse: onresponse,
			}
			t.requestqueuesent[callbackid] = request
		}
	}

	if t.isLoggingEnabled {
		util.Log(fmt.Sprintf("Sending, %s %s", request.Command, request.Service))
	}

	request.SchwabClientCustomerId = t.user.StreamerInfo[0].SchwabClientCustomerId
	request.SchwabClientCorrelId = t.user.StreamerInfo[0].SchwabClientCorrelId

	reqs := model.TDWSRequests{}
	reqs.Requests = append(reqs.Requests, request)

	reqB, _ := json.Marshal(request)
	t.socket.SendText(string(reqB))
}

func (t *TDStream) getCommandResponse(request model.TDWSRequest, out interface{}) {
	gotresponse := false
	var response string

	request.Requestid = t.nextRequestID()
	callbackid := request.Requestid
	t.requestqueue[callbackid] = TDStreamOnResponseFunc{
		onResponse: func(message string) {
			response = message
			gotresponse = true
		},
	}
	t.requestqueuesent[callbackid] = request

	if request.Command == "SUBS" || request.Command == "ADD" {
		for _, k := range strings.Split(request.Parameters["keys"], ",") {
			callbackid = fmt.Sprintf("%s:%s", request.Service, k)
			t.requestqueue[callbackid] = TDStreamOnResponseFunc{
				onResponse: func(message string) {
					response = message
					gotresponse = true
				},
			}
			t.requestqueuesent[callbackid] = request
		}
	}

	if t.isLoggingEnabled {
		util.Log(fmt.Sprintf("Sending, %s %s", request.Command, request.Service))
	}

	request.SchwabClientCustomerId = t.user.StreamerInfo[0].SchwabClientCustomerId
	request.SchwabClientCorrelId = t.user.StreamerInfo[0].SchwabClientCorrelId

	reqs := model.TDWSRequests{}
	reqs.Requests = append(reqs.Requests, request)
	t.socket.SendText(util.Serialize(reqs))

	// Process Response
	for !gotresponse && !t.stop {
		if t.isLoggingEnabled {
			util.Log(fmt.Sprintf("Waiting for %s - %s", request.Service, request.Command))
		}
		time.Sleep(1000 * time.Millisecond)
	}

	err := util.Deserialize(response, out)
	if err != nil {
		util.OnError(err)
	}

	if t.isLoggingEnabled {
		util.Log(fmt.Sprintf("Received message - %s", response))
	}
}

func (t *TDStream) onResponse(message string, socket gowebsocket.Socket) {
	var q model.TDWSResponse[model.TDWSResponse_General]
	b := []byte(message)
	err := json.Unmarshal(b, &q)

	if err != nil {
		util.OnError(err)
	}

	requestid := ""
	for _, r := range q.Response {
		if requestid != "" {
			break
		}
		requestid = r.Requestid
		if requestid == "" && r.Content != nil {
			switch v := r.Content.(type) {
			case nil:
				break
			case map[string]interface{}:
				if v["0"] != nil {
					requestid = v["0"].(string)
				}
			case []interface{}:
				if len(v) > 0 && v[0].(map[string]interface{})["0"] != nil {
					requestid = v[0].(map[string]interface{})["0"].(string)
				}
			}
		}
		if f, ok := t.requestqueue[requestid]; ok && f.onResponse != nil {
			f.onResponse(message)
		}
		if t.OnResponse != nil {
			t.OnResponse(message)
		}
	}

	if requestid == "" {
		var h model.TDWSHeartBeat
		err := json.Unmarshal(b, &h)
		if err == nil && len(h.Notify) > 0 {
			t.isConnected = true
			if t.isLoggingEnabled {
				util.Log(fmt.Sprintf("TD Stream Heartbeat, %s", h.Notify[0].Heartbeat))
			}
			return
		}
	}

	var s model.TDWSResponse_L1_Root
	err = json.Unmarshal(b, &s)

	if err != nil {
		if t.isLoggingEnabled {
			util.Log("Unidentified Response")
			util.Log(message)
		}
	}

	for _, row := range s.Data {
		for _, item := range row.Content {
			callbackid := fmt.Sprintf("%s:%s", row.Service, item["key"].(string))
			if f, ok := t.requestqueue[callbackid]; ok && f.onResponse != nil {
				f.onResponse(util.Serialize(item))
			}
		}
	}
}

func (t *TDStream) GetFuturesOptionBook(symbol string) (error, string) {
	req := model.TDWSRequest{
		Service: "FUTURES_OPTIONS_BOOK",
		Command: "GET",
		Parameters: map[string]string{
			"keys": symbol,
		},
	}
	var rv interface{}
	t.getCommandResponse(req, &rv)
	return nil, util.Serialize(rv)
}

var subscriptionQueue map[string]func(err error, quote interface{})

func (t *TDStream) getSubscriptionCommand() string {
	return "ADD"
}

func (t *TDStream) getSubscription(
	service string,
	symbol string,
	fields string,
	onmessage func(err error, quote interface{})) {
	req := model.TDWSRequest{
		Service: service,
		Command: t.getSubscriptionCommand(),
		Parameters: map[string]string{
			"keys":   symbol,
			"fields": fields,
		},
	}
	if t.isLoggingEnabled {
		util.Log(fmt.Sprintf("Subscribing to %s, %s", service, symbol))
	}

	t.sendCommand(req, func(message string) {
		if t.isLoggingEnabled {
			util.Log(message)
		}
		var resp map[string]any
		b := []byte(message)
		err := json.Unmarshal(b, &resp)

		if err != nil {
			for _, onmessage := range subscriptionQueue {
				onmessage(err, nil)
			}
			return
		}

		onmessage(nil, &resp)
	})
}

func (t *TDStream) GetFuturesSub(symbol string, onmessage func(err error, quote *model.TDWSResponse_L1_Content_Futures)) {
	t.getSubscription(
		"LEVELONE_FUTURES",
		symbol,
		t.numberCSV(40),
		func(err error, quote interface{}) {
			var rv model.TDWSResponse_L1_Content_Futures
			if quote != nil {
				util.Clone(quote, &rv)
			}
			onmessage(err, &rv)
		})
}

func (t *TDStream) GetFuturesOptionSub(symbol string, onmessage func(err error, quote *model.TDWSResponse_L1_Content_FuturesOption)) {
	t.getSubscription(
		"LEVELONE_FUTURES_OPTIONS",
		symbol,
		t.numberCSV(35),
		func(err error, quote interface{}) {
			var rv model.TDWSResponse_L1_Content_FuturesOption
			if quote != nil {
				util.Clone(quote, &rv)
			}
			onmessage(err, &rv)
		})
}

func (t *TDStream) GetFuturesHistory(symbol, periodType, period, frequencyType string, frequency int) []model.Quote {
	rv := make([]model.Quote, 0)

	req := model.TDWSRequest{
		Service: "CHART_HISTORY_FUTURES",
		Command: "GET",
		Parameters: map[string]string{
			"symbol":    symbol,
			"frequency": fmt.Sprintf("%s%d", frequencyType, frequency),
			"period":    periodType + period,
		},
	}

	if t.isLoggingEnabled {
		util.Log("Fetching Futures History")
	}
	var q model.TDWSHistoryResponse
	t.getCommandResponse(req, &q)

	if q.Snapshots != nil && len(q.Snapshots) > 0 &&
		q.Snapshots[0].Content != nil && len(q.Snapshots[0].Content) > 0 &&
		q.Snapshots[0].Content[0].Quotes != nil && len(q.Snapshots[0].Content[0].Quotes) > 0 {
		for _, q := range q.Snapshots[0].Content[0].Quotes {
			rv = append(rv, model.Quote{
				Open:         q.Open,
				High:         q.High,
				Low:          q.Low,
				Close:        q.Close,
				Volume:       q.Volume,
				DateTimeEpoc: q.DateTime,

				DateTime: util.EpocToTime(q.DateTime),
			})
		}
	}

	return rv
}

func (t *TDStream) numberCSV(count int) string {
	var strvalues []string
	for i := 0; i <= count; i++ {
		strvalues = append(strvalues, fmt.Sprintf("%d", i))
	}
	return strings.Join(strvalues, ",")
}
