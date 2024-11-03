package marketdata

import (
	"fmt"
	"schwab-client-go/internal"
	"time"
)

type Quote struct {
	AssetMainType string
	AssetSubType  string
	AuoteType     string
	RealTime      bool
	Ssid          float64
	Symbol        string

	Fundamental Field_Fundamental
	Quote       Field_Quote
	Reference   Field_Reference
	Regular     Field_Regular
}

type Field_Fundamental struct {
	Avg10DaysVolume    float64
	Avg1YearVolume     float64
	DivAmount          float64
	DivFreq            float64
	DivPayAmount       float64
	DivYield           float64
	Eps                float64
	FundLeverageFactor float64
	LastEarningsDate   *time.Time
	PeRatio            float64
}

type Field_Quote struct {
	FiftyTwoWeekHigh        float64 `json:"52WeekHigh"`
	FiftyTwoWeekLow         float64 `json:"52WeekLow"`
	AskMICId                string
	AskPrice                float64
	AskSize                 float64
	AskTime                 float64
	BidMICId                string
	BidPrice                float64
	BidSize                 float64
	BidTime                 float64
	ClosePrice              float64
	HighPrice               float64
	LastMICId               float64
	LastPrice               float64
	LastSize                float64
	LowPrice                float64
	Mark                    float64
	MarkChange              float64
	MarkPercentChange       float64
	NetChange               float64
	NetPercentChange        float64
	OpenPrice               float64
	PostMarketChange        float64
	PostMarketPercentChange float64
	QuoteTime               float64
	SecurityStatus          string
	TotalVolume             float64
	TradeTime               float64
}

type Field_Reference struct {
	Cusip          string
	Description    string
	Exchange       string
	ExchangeName   string
	IsHardToBorrow bool
	IsShortable    bool
	HtbRate        float64
}

type Field_Regular struct {
	RegularMarketLastPrice     float64
	RegularMarketLastSize      float64
	RegularMarketNetChange     float64
	RegularMarketPercentChange float64
	RegularMarketTradeTime     float64
}

type Quotes struct{}

func (c Quotes) GetQuote(symbol string) (rv map[string]Quote, err error) {
	url := fmt.Sprintf("%s/%s/quotes?fields=all", internal.Endpoints.MarketData, symbol)
	err = internal.API.Execute(url, &rv)
	return
}
