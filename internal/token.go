package internal

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"schwab-client-go/util"
	"time"
)

var (
	clientID     = ""
	clientSecret = ""
	redirectURL  = ""
)

type responseToken struct {
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	IDToken      string `json:"id_token"`
}

type token struct {
	BearerToken        string
	BearerTokenExpiry  time.Time
	RefreshToken       string
	RefreshTokenExpiry time.Time
}

var Token = &token{}

func (c *token) GetToken() string {
	clientID = os.Getenv("schwab_appkey")
	clientSecret = os.Getenv("schwab_secret")
	redirectURL = os.Getenv("schwab_redirecturl")

	err := util.Util.FromFile(os.Getenv("schwab_tokenpath"), c)
	if err != nil || c.BearerToken == "" || time.Now().After(c.RefreshTokenExpiry) {
		err = c.LoadNewToken()

		if err != nil {
			util.Util.OnError(err)
			return ""
		}
	}

	if time.Now().After(c.BearerTokenExpiry) {
		payload := fmt.Sprintf("grant_type=refresh_token&client_id=%s&client_secret=%s&refresh_token=%s", clientID, clientSecret, c.RefreshToken)
		err = c.fetchTokens(payload)
		if err != nil {
			util.Util.OnError(err)
			return ""
		}
	}

	return c.BearerToken
}

func (c *token) LoadNewToken() error {
	// 1 : Start local call back Server
	authCodeEncoded := ""
	http.HandleFunc("/schwab/callback", func(w http.ResponseWriter, r *http.Request) {
		// Use the code to get a token.
		authCodeEncoded = r.URL.Query().Get("code")
	})
	go http.ListenAndServe(":9999", nil)

	// 2 : OAuth - Authorization Code
	util.Util.OpenBrowser(fmt.Sprintf("%s?client_id=%s&redirect_uri=%s", Endpoints.Auth, clientID, redirectURL))
	util.Util.Log("Waiting fot Auth Code...")
	for authCodeEncoded == "" {
		time.Sleep(1 * time.Second)
	}
	authCode, err := url.QueryUnescape(authCodeEncoded)
	if util.Util.OnError(err) != nil {
		return err
	}

	// 3 : Get Refresh, Bearer Tokens
	payload := fmt.Sprintf("grant_type=authorization_code&code=%s&redirect_uri=%s", authCode, redirectURL)
	return c.fetchTokens(payload)
}

func (c *token) fetchTokens(payload string) error {
	authHeader := fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", clientID, clientSecret))))
	req, err := http.NewRequest("POST", Endpoints.Token, bytes.NewBuffer([]byte(payload)))
	if util.Util.OnError(err) != nil {
		return err
	}
	req.Header = http.Header{
		"Authorization": {authHeader},
		"Content-Type":  {"application/x-www-form-urlencoded"},
	}

	httpclient := http.Client{}
	res, err := httpclient.Do(req)
	if util.Util.OnError(err) != nil {
		return err
	}
	defer res.Body.Close()
	responseBytes, err := io.ReadAll(res.Body)
	if util.Util.OnError(err) != nil {
		return err
	}
	response := string(responseBytes)
	rt := &responseToken{}
	util.Util.FromJson(response, rt)

	c.RefreshTokenExpiry = time.Now().Add(time.Hour * 168)
	c.BearerTokenExpiry = time.Now().Add(time.Duration(rt.ExpiresIn * int(time.Second)))
	c.RefreshToken = rt.RefreshToken
	c.BearerToken = rt.AccessToken

	util.Util.ToFile(os.Getenv("schwab_tokenpath"), c)

	return nil
}
