package internal

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/Noob-Trading-Inc/schwab-client-go/util"
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

func (c *token) Reset() {
	os.Remove(os.Getenv("schwab_tokenpath"))
	c.BearerToken = ""
	c.RefreshToken = ""
}

func (c *token) GetTokenForHeader() string {
	return "Bearer " + c.GetToken()
}

func (c *token) GetToken() string {
	clientID = os.Getenv("schwab_appkey")
	clientSecret = os.Getenv("schwab_secret")
	redirectURL = os.Getenv("schwab_redirecturl")

	err := util.FromFile(os.Getenv("schwab_tokenpath"), c)
	if err != nil || c.BearerToken == "" || time.Now().After(c.RefreshTokenExpiry) {
		err = c.LoadNewToken()

		if err != nil {
			util.OnError(err)
			return ""
		}
	}

	if time.Now().After(c.BearerTokenExpiry) {
		payload := fmt.Sprintf("grant_type=refresh_token&client_id=%s&client_secret=%s&refresh_token=%s", clientID, clientSecret, c.RefreshToken)
		err = c.fetchTokens(payload)
		if err != nil {
			util.OnError(err)
			return ""
		}
	}

	if time.Now().After(c.RefreshTokenExpiry) {
		log.Fatal("Refresh Token Expired, get new token")
	}

	return c.BearerToken
}

func (c *token) LoadNewToken() error {
	util.Log("Getting New Token...")

	// 1 : Start local call back Server
	authCodeEncoded := ""
	http.HandleFunc("/schwab/callback", func(w http.ResponseWriter, r *http.Request) {
		// Use the code to get a token.
		authCodeEncoded = r.URL.Query().Get("code")
	})
	go http.ListenAndServeTLS("127.0.0.1:9999", "ssl/localhost.crt", "ssl/localhost.key", nil)

	// 2 : OAuth - Authorization Code
	util.OpenBrowser(fmt.Sprintf("%s?client_id=%s&redirect_uri=%s&scope=readonly&response_type=code", Endpoints.Auth, clientID, redirectURL))
	util.Log("Waiting fot Auth Code...")
	for authCodeEncoded == "" {
		time.Sleep(1 * time.Second)
	}
	authCode, err := url.QueryUnescape(authCodeEncoded)
	if util.OnError(err) != nil {
		return err
	}

	// 3 : Get Refresh, Bearer Tokens
	payload := fmt.Sprintf("grant_type=authorization_code&code=%s&client_id=%s&redirect_uri=%s", authCode, clientID, redirectURL)
	return c.fetchTokens(payload)
}

func (c *token) fetchTokens(payload string) error {
	util.Log("Refreshing Token...")
	authHeader := fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", clientID, clientSecret))))
	req, err := http.NewRequest("POST", Endpoints.Token, bytes.NewBuffer([]byte(payload)))
	if util.OnError(err) != nil {
		return err
	}
	req.Header = http.Header{
		"Authorization": {authHeader},
		"Content-Type":  {"application/x-www-form-urlencoded"},
	}

	httpclient := http.Client{}
	res, err := httpclient.Do(req)
	if util.OnError(err) != nil {
		return err
	}
	defer res.Body.Close()
	responseBytes, err := io.ReadAll(res.Body)
	if util.OnError(err) != nil {
		return err
	}
	response := string(responseBytes)
	rt := &responseToken{}
	util.Deserialize(response, rt)

	if c.RefreshToken != rt.RefreshToken {
		c.RefreshTokenExpiry = time.Now().Add(time.Hour * 7 * 24)
		c.RefreshToken = rt.RefreshToken
	}
	c.BearerTokenExpiry = time.Now().Add(time.Duration(rt.ExpiresIn * int(time.Second)))
	c.BearerToken = rt.AccessToken

	util.ToFile(os.Getenv("schwab_tokenpath"), c)

	return nil
}
