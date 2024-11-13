package internal

import (
	"bytes"
	"io"
	"net/http"

	"github.com/Noob-Trading-Inc/schwab-client-go/util"
)

type api struct{}

var API = &api{}

func (c *api) Get(url string, response any) (err error) {
	var req *http.Request
	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		util.OnError(err)
		return
	}

	_, err = c.Do(req, response)
	return
}

func (c *api) Post(url string, request any, response any) (location string, err error) {
	var req *http.Request
	req, err = http.NewRequest("POST", url, bytes.NewBuffer([]byte(util.Serialize(request))))
	if err != nil {
		util.OnError(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	var res *http.Response
	res, err = c.Do(req, response)
	location = res.Header.Get("Location")

	return
}

func (c *api) Put(url string, request any, response any) (location string, err error) {
	var req *http.Request
	req, err = http.NewRequest("PUT", url, bytes.NewBuffer([]byte(util.Serialize(request))))
	if err != nil {
		util.OnError(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	var res *http.Response
	res, err = c.Do(req, response)
	location = res.Header.Get("Location")
	return
}

func (c *api) Delete(url string, response any) (err error) {
	var req *http.Request
	req, err = http.NewRequest("DELETE", url, nil)
	if err != nil {
		util.OnError(err)
		return
	}

	_, err = c.Do(req, response)
	return
}

func (c *api) Do(req *http.Request, response any) (res *http.Response, err error) {
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", Token.GetTokenForHeader())

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

	if res.ContentLength == 0 || len(responsestr) == 0 || response == nil {
		return
	}

	util.Deserialize(responsestr, response)
	return
}
