package model

import "time"

type Quote struct {
	Open         float64 `json:"open,omitempty"`
	High         float64 `json:"high,omitempty"`
	Low          float64 `json:"low,omitempty"`
	Close        float64 `json:"close,omitempty"`
	DateTimeEpoc int64   `json:"datetime,omitempty"`

	Bid     float64 `json:"bid,omitempty"`
	Ask     float64 `json:"ask,omitempty"`
	BidSize float64 `json:"bidsize,omitempty"`
	AskSize float64 `json:"asksize,omitempty"`

	Volume float64 `json:"volume,omitempty"`

	VWapU float64
	VWapL float64
	VWap  float64

	RSIU float64
	RSID float64
	RSI  float64

	ATR          float64
	ATRBuyPrice  float64
	ATRSellPrice float64

	EMA12          float64
	EMA26          float64
	MACDSignalLine float64
	MACDLine       float64
	MACDHistogram  float64

	DateTime time.Time
	Candle   string
	Pattern  string
	Trend    string

	VWapAbs string
	RSIAbs  string
	ATRAbs  string
	MACDAbs string

	AVGC13 float64
	AVGH21 float64
	AVGL21 float64
	MINL3  float64
	MAXH10 float64
}

type TDWSRequests struct {
	Requests []TDWSRequest `json:"requests,omitempty"`
}

type TDWSRequest struct {
	Requestid              string            `json:"requestid,omitempty"`
	Service                string            `json:"service,omitempty"`
	Command                string            `json:"command,omitempty"`
	SchwabClientCustomerId string            `json:"SchwabClientCustomerId,omitempty"`
	SchwabClientCorrelId   string            `json:"SchwabClientCorrelId,omitempty"`
	Parameters             map[string]string `json:"parameters,omitempty"`
}

type TDWSHeartBeat struct {
	Notify []TDWSHeartBeatTime `json:"notify,omitempty"`
}
type TDWSHeartBeatTime struct {
	Heartbeat string `json:"heartbeat,omitempty"`
}

type TDWSHistoryResponse struct {
	Snapshots []TDWSSnapshot `json:"snapshot,omitempty"`
}

type TDWSResponse[T any] struct {
	Response []T `json:"response,omitempty"`
}

type TDWSResponse_General struct {
	Requestid            string `json:"requestid,omitempty"`
	Service              string `json:"service,omitempty"`
	Timestamp            int64  `json:"timestamp,omitempty"`
	Command              string `json:"command,omitempty"`
	SchwabClientCorrelId string `json:"SchwabClientCorrelId,omitempty"`
	Content              any    `json:"content,omitempty"`
}

type TDWSSnapshot struct {
	Requestid            string                `json:"requestid,omitempty"`
	Service              string                `json:"service,omitempty"`
	Timestamp            int64                 `json:"timestamp,omitempty"`
	Command              string                `json:"command,omitempty"`
	SchwabClientCorrelId string                `json:"SchwabClientCorrelId,omitempty"`
	Content              []TDWSSnapshotContent `json:"content,omitempty"`
}

type TDWSSnapshotContent struct {
	Requestid string      `json:"0,omitempty"`
	One       int         `json:"1,omitempty"`
	Two       int         `json:"2,omitempty"`
	Quotes    []TDWSQuote `json:"3,omitempty"`
	Symbol    string      `json:"key,omitempty"`
}

type TDWSQuote struct {
	DateTime int64   `json:"0,omitempty"`
	Open     float64 `json:"1,omitempty"`
	High     float64 `json:"2,omitempty"`
	Low      float64 `json:"3,omitempty"`
	Close    float64 `json:"4,omitempty"`
	Volume   float64 `json:"5,omitempty"`
}

type TDWSResponse_L1_Root struct {
	Data []TDWSResponse_L1_Item `json:"data,omitempty"`
}

type TDWSResponse_L1_Item struct {
	Service   string           `json:"service,omitempty"`
	Timestamp int64            `json:"timestamp,omitempty"`
	Command   string           `json:"command,omitempty"`
	Content   []map[string]any `json:"content,omitempty"`
}

type TDWSResponse_L1_Content_Common struct {
	Symbol  string `json:"key,omitempty"`     //Ticker symbol in upper case.
	Delayed bool   `json:"delayed,omitempty"` //Is Data delayed.
}

type TDWSResponse_L1_Content_Futures struct {
	TDWSResponse_L1_Content_Common
	AssetMainType string `json:"assetMainType,omitempty"` //Underlying asset type.

	BidPrice              float64 `json:"1,omitempty"`  //Current Best Bid Price
	AskPrice              float64 `json:"2,omitempty"`  //Current Best Ask Price
	LastPrice             float64 `json:"3,omitempty"`  //Price at which the last trade was matched
	BidSize               int64   `json:"4,omitempty"`  //Number of shares for bid
	AskSize               int64   `json:"5,omitempty"`  //Number of shares for ask
	BidID                 string  `json:"6,omitempty"`  //Exchange with the best bid
	AskID                 string  `json:"7,omitempty"`  //Exchange with the best ask
	TotalVolume           float64 `json:"8,omitempty"`  //Aggregated shares traded throughout the day, including pre/post market hours.
	LastSize              int64   `json:"9,omitempty"`  //Number of shares traded with last trade
	QuoteTime             int64   `json:"10,omitempty"` //Trade time of the last quote in milliseconds since epoch
	TradeTime             int64   `json:"11,omitempty"` //Trade time of the last trade in milliseconds since epoch
	HighPrice             float64 `json:"12,omitempty"` //Day’s high trade price
	LowPrice              float64 `json:"13,omitempty"` //Day’s low trade price
	ClosePrice            float64 `json:"14,omitempty"` //Previous day’s closing price
	ExchangeID            string  `json:"15,omitempty"` //Primary "listing" Exchange
	Description           string  `json:"16,omitempty"` //Description of the product
	LastID                string  `json:"17,omitempty"` //Exchange where last trade was executed
	OpenPrice             float64 `json:"18,omitempty"` //Day's Open Price
	NetChange             float64 `json:"19,omitempty"` //Current Last-Prev Close
	FuturePercentChange   float64 `json:"20,omitempty"` //Current percent change
	ExhangeName           string  `json:"21,omitempty"` //Name of exchange
	SecurityStatus        string  `json:"22,omitempty"` //Trading status of the symbol
	OpenInterest          int     `json:"23,omitempty"` //The total number of futures ontracts that are not closed or delivered on a particular day
	Mark                  float64 `json:"24,omitempty"` //Mark-to-Market value is calculated daily using current prices to determine profit/loss
	Tick                  float64 `json:"25,omitempty"` //Minimum price movement
	TickAmount            float64 `json:"26,omitempty"` //Minimum amount that the price of the market can change
	Product               string  `json:"27,omitempty"` //Futures product
	FuturePriceFormat     string  `json:"28,omitempty"` //Display in fraction or decimal format.
	FutureTradingHours    string  `json:"29,omitempty"` //Trading hours
	FutureIsTradable      bool    `json:"30,omitempty"` //Flag to indicate if this future contract is tradable
	FutureMultiplier      float64 `json:"31,omitempty"` //Point value
	FutureIsActive        bool    `json:"32,omitempty"` //Indicates if this contract is active
	FutureSettlementPrice float64 `json:"33,omitempty"` //Closing price
	FutureActiveSymbol    string  `json:"34,omitempty"` //Symbol of the active contract
	FutureExpirationDate  int64   `json:"35,omitempty"` //Expiration date of this contract
	ExpirationStyle       string  `json:"36,omitempty"`
	AskTime               int64   `json:"37,omitempty"`
	BidTime               int64   `json:"38,omitempty"`
	QuotedInSession       bool    `json:"39,omitempty"`
	SettlementDate        int64   `json:"40,omitempty"`
}

type TDWSResponse_L1_Content_FuturesOption struct {
	TDWSResponse_L1_Content_Common

	BidPrice              float64 `json:"1,omitempty"`  //Current Best Bid Price
	AskPrice              float64 `json:"2,omitempty"`  //Current Best Ask Price
	LastPrice             float64 `json:"3,omitempty"`  //Price at which the last trade was matched
	BidSize               int64   `json:"4,omitempty"`  //Number of shares for bid
	AskSize               int64   `json:"5,omitempty"`  //Number of shares for ask
	BidID                 float64 `json:"6,omitempty"`  //Exchange with the best bid
	AskID                 float64 `json:"7,omitempty"`  //Exchange with the best ask
	TotalVolume           int64   `json:"8,omitempty"`  //Aggregated shares traded throughout the day, including pre/post market hours.
	LastSize              int64   `json:"9,omitempty"`  //Number of shares traded with last trade
	QuoteTime             int64   `json:"10,omitempty"` //Trade time of the last quote in milliseconds since epoch
	TradeTime             int64   `json:"11,omitempty"` //Trade time of the last trade in milliseconds since epoch
	HighPrice             float64 `json:"12,omitempty"` //Day’s high trade price
	LowPrice              float64 `json:"13,omitempty"` //Day’s low trade price
	ClosePrice            float64 `json:"14,omitempty"` //Previous day’s closing price
	LastID                float64 `json:"15,omitempty"` //Primary "listing" Exchange
	Description           string  `json:"16,omitempty"` //Description of the product
	OpenPrice             float64 `json:"17,omitempty"` //Exchange where last trade was executed
	OpenInterest          float64 `json:"18,omitempty"` //Day's Open Price
	Mark                  float64 `json:"19,omitempty"` //Market Price
	Tick                  float64 `json:"20,omitempty"` //Current percent change
	TickAmount            float64 `json:"21,omitempty"` //Name of exchange
	FutureMultiplier      float64 `json:"22,omitempty"` //Trading status of the symbol
	FutureSettlementPrice float64 `json:"23,omitempty"` //The total number of futures ontracts that are not closed or delivered on a particular day
	UnderlyingSymbol      string  `json:"24,omitempty"` //Mark-to-Market value is calculated daily using current prices to determine profit/loss
	StrikePrice           float64 `json:"25,omitempty"` //Minimum price movement
	FutureExpirationDate  int64   `json:"26,omitempty"` //Minimum amount that the price of the market can change
	ExpirationStyle       string  `json:"27,omitempty"` //Futures product
	ContractType          string  `json:"28,omitempty"` //Display in fraction or decimal format.
	SecurityStatus        string  `json:"29,omitempty"` //Trading hours
	Exchange              string  `json:"30,omitempty"` //Flag to indicate if this future contract is tradable
	ExchangeName          string  `json:"31,omitempty"` //Point value\
}

type TDWSResponse_L1_Content_Any struct {
	TDWSResponse_L1_Content_Common
	Field_1  any `json:"1,omitempty"`
	Field_2  any `json:"2,omitempty"`
	Field_3  any `json:"3,omitempty"`
	Field_4  any `json:"4,omitempty"`
	Field_5  any `json:"5,omitempty"`
	Field_6  any `json:"6,omitempty"`
	Field_7  any `json:"7,omitempty"`
	Field_8  any `json:"8,omitempty"`
	Field_9  any `json:"9,omitempty"`
	Field_10 any `json:"10,omitempty"`
	Field_11 any `json:"11,omitempty"`
	Field_12 any `json:"12,omitempty"`
	Field_13 any `json:"13,omitempty"`
	Field_14 any `json:"14,omitempty"`
	Field_15 any `json:"15,omitempty"`
	Field_16 any `json:"16,omitempty"`
	Field_17 any `json:"17,omitempty"`
	Field_18 any `json:"18,omitempty"`
	Field_19 any `json:"19,omitempty"`
	Field_20 any `json:"20,omitempty"`
	Field_21 any `json:"21,omitempty"`
	Field_22 any `json:"22,omitempty"`
	Field_23 any `json:"23,omitempty"`
	Field_24 any `json:"24,omitempty"`
	Field_25 any `json:"25,omitempty"`
	Field_26 any `json:"26,omitempty"`
	Field_27 any `json:"27,omitempty"`
	Field_28 any `json:"28,omitempty"`
	Field_29 any `json:"29,omitempty"`
	Field_30 any `json:"30,omitempty"`
	Field_31 any `json:"31,omitempty"`
	Field_32 any `json:"32,omitempty"`
	Field_33 any `json:"33,omitempty"`
	Field_34 any `json:"34,omitempty"`
	Field_35 any `json:"35,omitempty"`
	Field_36 any `json:"36,omitempty"`
	Field_37 any `json:"37,omitempty"`
	Field_38 any `json:"38,omitempty"`
	Field_39 any `json:"39,omitempty"`
	Field_40 any `json:"40,omitempty"`
	Field_41 any `json:"41,omitempty"`
	Field_42 any `json:"42,omitempty"`
	Field_43 any `json:"43,omitempty"`
	Field_44 any `json:"44,omitempty"`
	Field_45 any `json:"45,omitempty"`
	Field_46 any `json:"46,omitempty"`
	Field_47 any `json:"47,omitempty"`
	Field_48 any `json:"48,omitempty"`
	Field_49 any `json:"49,omitempty"`
	Field_50 any `json:"50,omitempty"`
	Field_51 any `json:"51,omitempty"`
}
