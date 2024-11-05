package trader

import (
	"fmt"
	"schwab-client-go/internal"
)

type AccountNumber struct {
	AccountNumber string
	HashValue     string
}

type Account struct {
	SecuritiesAccount Field_Account
	AggregatedBalance Field_AggregatedBalance
}

type Field_Account struct {
	AccountNumber           string
	HashValue               string
	RoundTrips              float64
	IsDayTrader             bool
	IsClosingOnlyRestricted bool
	PfcbFlag                bool

	Positions         []Field_Position
	InitialBalances   Field_Balance
	CurrentBalances   Field_Balance
	ProjectedBalances Field_Balance
}

type Field_Position struct {
	ShortQuantity                  float64
	AveragePrice                   float64
	CurrentDayProfitLoss           float64
	CurrentDayProfitLossPercentage float64
	LongQuantity                   float64
	SettledLongQuantity            float64
	SettledShortQuantity           float64
	AgedQuantity                   float64
	Instrument                     Field_Instrument
	MarketValue                    float64
	MaintenanceRequirement         float64
	AverageLongPrice               float64
	AverageShortPrice              float64
	TaxLotAverageLongPrice         float64
	TaxLotAverageShortPrice        float64
	LongOpenProfitLoss             float64
	ShortOpenProfitLoss            float64
	PreviousSessionLongQuantity    float64
	PreviousSessionShortQuantity   float64
	CurrentDayCost                 float64
}

type Field_Instrument struct {
	Cusip        string
	Symbol       string
	Description  string
	InstrumentId float64
	NetChange    float64
	Type         string
}

type Field_AggregatedBalance struct {
	CurrentLiquidationValue float64
	LiquidationValue        float64
}

type Field_Balance struct {
	LiquidationValue                 float64
	AccruedInterest                  float64
	AvailableFunds                   float64
	AvailableFundsNonMarginableTrade float64
	BondValue                        float64
	BuyingPower                      float64
	BuyingPowerNonMarginableTrade    float64
	StockBuyingPower                 float64
	OptionBuyingPower                float64
	CashBalance                      float64
	CashAvailableForTrading          float64
	CashReceipts                     float64
	DayTradingBuyingPower            float64
	DayTradingBuyingPowerCall        float64
	DayTradingEquityCall             float64
	Equity                           float64
	EquityPercentage                 float64
	LongMarginValue                  float64
	LongOptionMarketValue            float64
	LongStockValue                   float64
	MaintenanceCall                  float64
	MaintenanceRequirement           float64
	Margin                           float64
	MarginEquity                     float64
	MoneyMarketFund                  float64
	MutualFundValue                  float64
	RegTCall                         float64
	ShortMarginValue                 float64
	ShortOptionMarketValue           float64
	ShortStockValue                  float64
	TotalCash                        float64
	Sma                              float64
	IsInCall                         bool
	UnsettledCash                    float64
	PendingDeposits                  float64
	MarginBalance                    float64
	ShortBalance                     float64
	AccountValue                     float64
}

type Accounts struct {
}

func (c Accounts) GetAccountNumbers() (rv []AccountNumber, err error) {
	url := fmt.Sprintf("%s/%s/%s", internal.Endpoints.Trader, "accounts", "accountNumbers")
	err = internal.API.Execute(url, &rv)
	return
}

func (c Accounts) GetAccounts() (rv []Account, err error) {
	url := fmt.Sprintf("%s/%s?fields=positions", internal.Endpoints.Trader, "accounts")
	err = internal.API.Execute(url, &rv)
	return
}

func (c Accounts) GetAccount(number string) (rv Account, err error) {
	url := fmt.Sprintf("%s/%s/%s?fields=positions", internal.Endpoints.Trader, "accounts", number)
	err = internal.API.Execute(url, &rv)
	return
}
