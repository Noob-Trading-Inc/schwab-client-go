package internal

type endpoints struct {
	Auth       string
	Token      string
	Trader     string
	MarketData string
}

var Endpoints = &endpoints{
	Auth:       "https://api.schwabapi.com/v1/oauth/authorize",
	Token:      "https://api.schwabapi.com/v1/oauth/token",
	Trader:     "https://api.schwabapi.com/trader/v1",
	MarketData: "https://api.schwabapi.com/marketdata/v1",
}
