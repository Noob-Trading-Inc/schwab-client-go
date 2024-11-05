package main

import (
	"fmt"
	"schwab-client-go/internal"
	"schwab-client-go/internal/marketdata"
	"schwab-client-go/internal/trader"
	"schwab-client-go/util"
)

type client struct {
	Acounts        trader.Accounts
	UserPreference trader.UserPreference

	Quotes marketdata.Quotes
}

var Client = &client{}

func (c *client) Init() {
	util.Util.Log("Initializing")
	token := internal.Token.GetToken()
	util.Util.Log("Token : ", token)

	c.Acounts = trader.Accounts{}
	c.UserPreference = trader.UserPreference{}
	c.Quotes = marketdata.Quotes{}
}

func main() {
	q, err := Client.Quotes.GetQuote("TSLA")
	fmt.Println(util.Util.ToJsonReadable(q), err)

	an, err := Client.Acounts.GetAccountNumbers()
	fmt.Println(util.Util.ToJsonReadable(an), err)

	a, err := Client.Acounts.GetAccounts()
	fmt.Println(util.Util.ToJsonReadable(a), err)

	a1, err := Client.Acounts.GetAccount(an[0].HashValue)
	fmt.Println(util.Util.ToJsonReadable(a1), err)

	up, err := Client.UserPreference.GetUserPreference()
	fmt.Println(util.Util.ToJsonReadable(up), err)
}
