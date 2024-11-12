package schwab

import (
	"fmt"
	"math"
	"os"
	"os/signal"
	"time"

	"github.com/Noob-Trading-Inc/schwab-client-go/internal"
	"github.com/Noob-Trading-Inc/schwab-client-go/internal/marketdata"
	"github.com/Noob-Trading-Inc/schwab-client-go/internal/stream"
	"github.com/Noob-Trading-Inc/schwab-client-go/internal/stream/model"
	"github.com/Noob-Trading-Inc/schwab-client-go/internal/trader"
	"github.com/Noob-Trading-Inc/schwab-client-go/util"
)

type client struct {
	Acounts        trader.Accounts
	UserPreference trader.UserPreference
	Stream         stream.TDStream

	Quotes marketdata.Quotes
}

var Client = &client{}

func (c *client) InitWithRefreshToken(token string, expiresat time.Time) {
	internal.Token.SetRefreshToken(token, expiresat)
	c.Init()
}

func (c *client) Init() {
	util.Log("Initializing")
	token := internal.Token.GetToken()
	util.Log("Token : ", token)

	c.Acounts = trader.Accounts{}
	c.UserPreference = trader.UserPreference{}
	c.Stream = stream.TDStream{}

	c.Quotes = marketdata.Quotes{}

	k := make(chan os.Signal, 1)
	signal.Notify(k, os.Interrupt)
	go func() {
		for sig := range k {
			// sig is a ^C, handle it
			if sig == os.Interrupt || sig == os.Kill {
				c.Shutdown()
			}
		}
	}()
}

func (c *client) Shutdown() {
	stop = true
}

var stop bool = false

func main() {

	an, err := Client.Acounts.GetAccountNumbers()
	fmt.Println(util.SerializeReadable(an), err)
	/*
		q, err := Client.Quotes.GetQuote("TSLA")
		fmt.Println(util.SerializeReadable(q), err)

		a, err := Client.Acounts.GetAccounts()
		fmt.Println(util.SerializeReadable(a), err)

		a1, err := Client.Acounts.GetAccount(an[0].HashValue)
		fmt.Println(util.SerializeReadable(a1), err)
	*/

	up, err := Client.UserPreference.GetUserPreference()
	fmt.Println(util.SerializeReadable(up), err)

	//Client.Stream.EnableLogging()
	Client.Stream.Init(up)

	var NQPrice float64
	go Client.Stream.GetFuturesSub("/NQ", func(err error, quote *model.TDWSResponse_L1_Content_Futures) {
		if err != nil {
			util.Log("ERROR Futures L1, %s", err.Error())
			return
		}
		util.Log(util.Serialize(quote))
		if NQPrice == 0 {
			NQPrice = quote.Mark
		}
	})

	for NQPrice == 0 {
		time.Sleep(10 * time.Millisecond)
	}
	NQPutPrice := fmt.Sprintf("%.0f", (math.Round(NQPrice/100)*100)-100)
	NQCallPrice := fmt.Sprintf("%.0f", (math.Round(NQPrice/100)*100)+100)
	go Client.Stream.GetFuturesOptionSub("./QN2X24C"+NQPutPrice, func(err error, quote *model.TDWSResponse_L1_Content_FuturesOption) {
		if err != nil {
			util.Log("ERROR Futures L1, %s", err.Error())
			return
		}
		util.Log(util.Serialize(quote))
	})
	go Client.Stream.GetFuturesOptionSub("./QN2X24C"+NQCallPrice, func(err error, quote *model.TDWSResponse_L1_Content_FuturesOption) {
		if err != nil {
			util.Log("ERROR Futures L1, %s", err.Error())
			return
		}
		util.Log(util.Serialize(quote))
	})

	go Client.Stream.GetFuturesSub("/ES", func(err error, quote *model.TDWSResponse_L1_Content_Futures) {
		if err != nil {
			util.Log("ERROR Futures L1, %s", err.Error())
			return
		}
		util.Log(util.Serialize(quote))
	})

	for !stop {
		time.Sleep(1000 * time.Millisecond)
	}
}
