package schwab

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/Noob-Trading-Inc/schwab-client-go/internal"
	"github.com/Noob-Trading-Inc/schwab-client-go/internal/marketdata"
	"github.com/Noob-Trading-Inc/schwab-client-go/internal/stream"
	"github.com/Noob-Trading-Inc/schwab-client-go/internal/stream/model"
	"github.com/Noob-Trading-Inc/schwab-client-go/internal/trader"
	"github.com/Noob-Trading-Inc/schwab-client-go/models"
	"github.com/Noob-Trading-Inc/schwab-client-go/util"
)

type client struct {
	Acounts        trader.Accounts
	Orders         trader.Orders
	UserPreference trader.UserPreference
	Stream         stream.TDStream

	Quotes marketdata.Quotes
}

var Client = &client{}

func (c *client) InitWithRefreshToken(token string, expiresat time.Time) (err error) {
	internal.Token.SetRefreshToken(token, expiresat)
	err = c.Init()
	return
}

func (c *client) Init() (err error) {
	util.Log("Initializing schwab-client-go")
	token := internal.Token.GetToken()
	if token == "" {
		err = fmt.Errorf("Empty accesstoken")
		return
	}
	util.Logf("Token : %s....", token[0:5])

	c.Acounts = trader.Accounts{}
	c.Orders = trader.Orders{}
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
	return
}

var shutdownInprogress bool

func (c *client) Shutdown() {
	if shutdownInprogress {
		return
	}
	shutdownInprogress = true

	if isStreamInitiated {
		Client.Stream.Dispose()
	}
}

func (c *client) StreamOnReconnect(callback func()) {
	Client.Stream.OnConnect = callback
}

var isStreamInitiated bool
var isStreamInitiatedLock = sync.RWMutex{}

func (c *client) StreamQuotes(symbols []string, callback func(*models.Quote) error) error {
	if !isStreamInitiated {
		isStreamInitiatedLock.Lock()
		if !isStreamInitiated {
			up, err := Client.UserPreference.GetUserPreference()
			if err != nil {
				return util.OnError(err)
			}
			Client.Stream.Init(up)
		}
	}

	for _, symbol := range symbols {
		if symbol == "" {
			continue
		}

		if symbol[0:1] == "/" {
			Client.Stream.Subscribe_L1_Futures(symbol, func(err error, quote *model.TDWSResponse_L1_Content_Futures) {
				if err != nil {
					util.Log("ERROR Futures L1, %s", err.Error())
					return
				}
				util.Serialize(quote)
				callback(&models.Quote{
					Symbol: quote.Symbol,

					Open:  quote.OpenPrice,
					Close: quote.ClosePrice,
					High:  quote.HighPrice,
					Low:   quote.LowPrice,

					AskPrice:    quote.AskPrice,
					BidPrice:    quote.BidPrice,
					MarketPrice: quote.Mark,

					QuoteTimeInLong:  quote.QuoteTime,
					TotalVolume:      quote.TotalVolume,
					NetChange:        quote.NetChange,
					NetPercentChange: quote.FuturePercentChange,
				})
			})
			continue
		}

		Client.Stream.Subscribe_L1_Equity(symbol, func(err error, quote *model.TDWSResponse_L1_Content_Equity) {
			if err != nil {
				util.Log("ERROR Equity L1, %s", err.Error())
				return
			}
			util.Serialize(quote)
			callback(&models.Quote{
				Symbol: quote.Symbol,

				Open:  quote.OpenPrice,
				Close: quote.ClosePrice,
				High:  quote.HighPrice,
				Low:   quote.LowPrice,

				AskPrice:    quote.AskPrice,
				BidPrice:    quote.BidPrice,
				MarketPrice: quote.MarkPrice,

				FiftyTwoWeekHigh: quote.FiftyTwoWeekHigh,
				FiftyTwoWeekLow:  quote.FiftyTwoWeekLow,

				QuoteTimeInLong:  quote.QuoteTimeInLong,
				TotalVolume:      quote.TotalVolume,
				NetChange:        quote.NetChange,
				NetPercentChange: quote.NetPercentChange,
			})
		})
	}

	return nil
}
