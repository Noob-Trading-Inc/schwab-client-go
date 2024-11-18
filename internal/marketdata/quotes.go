package marketdata

import (
	"fmt"
	"net/url"

	"github.com/Noob-Trading-Inc/schwab-client-go/internal"
	"github.com/Noob-Trading-Inc/schwab-client-go/models"
)

type Quotes struct{}

func (c Quotes) GetQuote(symbol string) (rv map[string]models.EquityResponse, err error) {
	url := fmt.Sprintf("%s/%s/quotes?fields=all", internal.Endpoints.MarketData, url.QueryEscape(symbol))
	err = internal.API.Get(url, &rv)
	return
}

func (c Quotes) GetXMinuteCandles(symbol string, xminute int, from int64, to int64) (rv models.CandleList, err error) {
	url := fmt.Sprintf("%s/%s?symbol=%s&startDate=%d&endDate=%d&needExtendedHoursData=true&needPreviousClose=true&frequencyType=minute&frequency=%d",
		internal.Endpoints.MarketData, "pricehistory", url.QueryEscape(symbol), from, to, xminute)
	err = internal.API.Get(url, &rv)
	return

}
