package internal

import (
	"io"
	"net/http"

	"github.com/Noob-Trading-Inc/schwab-client-go/util"
)

type api struct{}

var API = &api{}

func (c *api) Execute(url string, response any) (err error) {
	var req *http.Request
	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		util.OnError(err)
		return
	}

	req.Header = http.Header{
		"accept":        {"application/json"},
		"Authorization": {Token.GetTokenForHeader()},
	}

	var res *http.Response
	httpclient := http.Client{}
	res, err = httpclient.Do(req)
	if err != nil {
		util.OnError(err)
		return
	}
	defer res.Body.Close()
	responseBytes, err := io.ReadAll(res.Body)
	if err != nil {
		util.OnError(err)
		return
	}

	responsestr := string(responseBytes)
	if res.StatusCode < 200 || res.StatusCode > 299 {
		err = NewApiError(responsestr)
		util.OnError(err)
		return
	}

	util.Deserialize(responsestr, response)
	return
}
