package marketdata

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/Noob-Trading-Inc/schwab-client-go/internal"
	"github.com/Noob-Trading-Inc/schwab-client-go/models"
)

type Quotes struct{}

///quotes?symbols=%2FMNQ&fields=quote%2Creference&indicative=false

func (c Quotes) GetEquityQuotes(symbols ...string) (rv map[string]models.EquityResponse, err error) {
	var eSymbols = make([]string, len(symbols))
	for i, symbol := range symbols {
		eSymbols[i] = url.QueryEscape(symbol)
	}
	url := fmt.Sprintf("%s/quotes?fields=all&symbols=%s", internal.Endpoints.MarketData, strings.Join(eSymbols, ","))
	err = internal.API.Get(url, &rv)
	return
}

func (c Quotes) GetFuturesQuotes(symbols ...string) (rv map[string]models.FutureResponse, err error) {
	var eSymbols = make([]string, len(symbols))
	for i, symbol := range symbols {
		eSymbols[i] = url.QueryEscape(symbol)
	}
	url := fmt.Sprintf("%s/quotes?fields=all&symbols=%s", internal.Endpoints.MarketData, strings.Join(eSymbols, ","))
	err = internal.API.Get(url, &rv)
	return
}

func (c Quotes) GetEquityQuote(symbol string) (rv map[string]models.EquityResponse, err error) {
	url := fmt.Sprintf("%s/%s/quotes?fields=all", internal.Endpoints.MarketData, url.QueryEscape(symbol))
	err = internal.API.Get(url, &rv)
	return
}

func (c Quotes) GetXMinuteCandles(symbol string, xminute int, from int64, to int64) (rv models.CandleList, err error) {
	return c.getCandles(symbol, 1, "day", xminute, "minute", from, to)
}

func (c Quotes) GetXDaysCandles(symbol string, from int64, to int64) (rv models.CandleList, err error) {
	return c.getCandles(symbol, 1, "month", 1, "daily", from, to)
}

func (c Quotes) GetXWeeksCandles(symbol string, from int64, to int64) (rv models.CandleList, err error) {
	return c.getCandles(symbol, 1, "month", 1, "weekly", from, to)
}

func (c Quotes) GetXMonthsCandles(symbol string, from int64, to int64) (rv models.CandleList, err error) {
	return c.getCandles(symbol, 1, "year", 1, "monthly", from, to)
}

func (c Quotes) getCandles(symbol string, period int, periodType string, frequency int, frequencyType string, from int64, to int64) (rv models.CandleList, err error) {
	url := fmt.Sprintf("%s/%s?symbol=%s&startDate=%d&endDate=%d&needExtendedHoursData=true&needPreviousClose=true&periodType=%s&period=%d&frequencyType=%s&frequency=%d",
		internal.Endpoints.MarketData, "pricehistory", url.QueryEscape(symbol), from, to, periodType, period, frequencyType, frequency)
	err = internal.API.Get(url, &rv)
	return
}

/*
periodType : day/month/year/ytd
string
The chart period being requested.


period
integer($int32)
The number of chart period types.

If the periodType is
• day - valid values are 1, 2, 3, 4, 5, 10
• month - valid values are 1, 2, 3, 6
• year - valid values are 1, 2, 3, 5, 10, 15, 20
• ytd - valid values are 1

If the period is not specified and the periodType is
• day - default period is 10.
• month - default period is 1.
• year - default period is 1.
• ytd - default period is 1.


frequencyType
string
The time frequencyType

If the periodType is
• day - valid value is minute
• month - valid values are daily, weekly
• year - valid values are daily, weekly, monthly
• ytd - valid values are daily, weekly

If frequencyType is not specified, default value depends on the periodType
• day - defaulted to minute.
• month - defaulted to weekly.
• year - defaulted to monthly.
• ytd - defaulted to weekly.


frequency
integer($int32)
The time frequency duration

If the frequencyType is
• minute - valid values are 1, 5, 10, 15, 30
• daily - valid value is 1
• weekly - valid value is 1
• monthly - valid value is 1

If frequency is not specified, default value is 1
*/
