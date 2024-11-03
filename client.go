package main

import (
	"fmt"
	"schwab-client-go/internal"
	"schwab-client-go/internal/marketdata"
	"schwab-client-go/internal/trader"
	"schwab-client-go/util"
)

type client struct {
	Acounts trader.Accounts
	Quotes  marketdata.Quotes
}

var Client = &client{}

func (c *client) Init() {
	util.Util.Log("Initializing")
	token := internal.Token.GetToken()
	util.Util.Log("Token : ", token)

	c.Acounts = trader.Accounts{}
	c.Quotes = marketdata.Quotes{}
}

func main() {
	q, err := Client.Quotes.GetQuote("TSLA")
	fmt.Println(util.Util.ToJsonReadable(q), err)

	a, err := Client.Acounts.GetAccountNumbers()
	fmt.Println(util.Util.ToJsonReadable(a), err)
}
