package internal

import (
	"fmt"
	"io"
	"net/http"
	"schwab-client-go/util"
)

type api struct{}

var API = &api{}

func (c *api) Execute(url string, response any) (err error) {
	var req *http.Request
	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		util.Util.OnError(err)
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
		util.Util.OnError(err)
		return
	}
	defer res.Body.Close()
	responseBytes, err := io.ReadAll(res.Body)
	if err != nil {
		util.Util.OnError(err)
		return
	}

	if res.StatusCode == 401 {
		err = fmt.Errorf("401 Unauthorized")
		return
	}
	responsestr := string(responseBytes)
	if res.StatusCode < 200 || res.StatusCode > 299 {
		err = fmt.Errorf(fmt.Sprintf("%d: %s", res.StatusCode, responsestr))
		return
	}

	util.Util.FromJson(responsestr, response)
	return
}
