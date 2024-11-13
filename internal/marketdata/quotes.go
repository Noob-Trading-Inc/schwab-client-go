package marketdata

import (
	"fmt"
	"net/url"

	"github.com/Noob-Trading-Inc/schwab-client-go/internal"
	"github.com/Noob-Trading-Inc/schwab-client-go/internal/marketdata/model"
)

type Quotes struct{}

func (c Quotes) GetQuote(symbol string) (rv map[string]model.EquityResponse, err error) {
	url := fmt.Sprintf("%s/%s/quotes?fields=all", internal.Endpoints.MarketData, url.QueryEscape(symbol))
	err = internal.API.Get(url, &rv)
	return
}
