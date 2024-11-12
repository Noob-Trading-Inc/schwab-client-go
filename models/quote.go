package models

type Quote struct {
	Symbol string `json:"symbol,omitempty"`

	Open  float64 `json:"open,omitempty"`
	Close float64 `json:"close,omitempty"`
	High  float64 `json:"high,omitempty"`
	Low   float64 `json:"low,omitempty"`

	AskPrice    float64 `json:"askPrice,omitempty"`
	BidPrice    float64 `json:"bidPrice,omitempty"`
	MarketPrice float64 `json:"marketPrice,omitempty"`

	FiftyTwoWeekHigh float64 `json:"52WeekHigh,omitempty"`
	FiftyTwoWeekLow  float64 `json:"52WeekLow,omitempty"`

	QuoteTimeInLong  int64   `json:"quoteTimeInLong,omitempty"`
	TotalVolume      float64 `json:"totalVolume,omitempty"`
	NetChange        float64 `json:"netChange,omitempty"`
	NetPercentChange float64 `json:"netPercentChange,omitempty"`
}
