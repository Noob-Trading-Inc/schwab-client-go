package schwab

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/Noob-Trading-Inc/schwab-client-go/models"
	"github.com/Noob-Trading-Inc/schwab-client-go/util"
)

func TestMain(m *testing.M) {
	Client.Init()
	code := m.Run()
	Client.Shutdown()
	os.Exit(code)
}

func Test_Accounts(t *testing.T) {
	a, err := Client.Acounts.GetAccountNumbers()
	fmt.Println(util.SerializeReadable(a), err)
}

func Test_Quotes(t *testing.T) {
	q, err := Client.Quotes.GetQuote("TSLA")
	fmt.Println(util.SerializeReadable(q), err)
}

func Test_Candles(t *testing.T) {
	q, err := Client.Quotes.GetXMinuteCandles("TSLA", 5, time.Now().UTC().AddDate(0, -1, 0).UnixMilli(), time.Now().UTC().UnixMilli())
	fmt.Println("TSLA Minute", len(q.Candles), err)

	q, err = Client.Quotes.GetXMinuteCandles("/NQ", 5, time.Now().UTC().AddDate(0, -1, 0).UnixMilli(), time.Now().UTC().UnixMilli())
	fmt.Println("/NQ Minute", len(q.Candles), err)

	q, err = Client.Quotes.GetXDaysCandles("/NQ", time.Now().UTC().AddDate(0, 0, -100).UnixMilli(), time.Now().UTC().UnixMilli())
	fmt.Println("/NQ Days", len(q.Candles), err)

	q, err = Client.Quotes.GetXWeeksCandles("/NQ", time.Now().UTC().AddDate(0, 0, -100).UnixMilli(), time.Now().UTC().UnixMilli())
	fmt.Println("/NQ Weeks", len(q.Candles), err)

	q, err = Client.Quotes.GetXMonthsCandles("/NQ", time.Now().UTC().AddDate(0, 0, -100).UnixMilli(), time.Now().UTC().UnixMilli())
	fmt.Println("/NQ Months", len(q.Candles), err)
}

func Test_Orders(t *testing.T) {
	a, _ := Client.Acounts.GetAccountNumbers()
	accountnumber := a[0].HashValue

	var err error
	o := models.Order{
		Quantity:          1,
		Price:             100,
		Session:           util.Ptr(models.SEAMLESS_Session),
		Duration:          util.Ptr(models.DAY_Duration),
		OrderType:         util.Ptr(models.LIMIT_OrderType),
		OrderStrategyType: util.Ptr(models.SINGLE_OrderStrategyType),
		OrderLegCollection: []models.OrderLegCollection{
			{
				Instrument: &models.AccountsInstrument{
					Symbol:    "AAPL",
					AssetType: "EQUITY",
				},
				Instruction: util.Ptr(models.BUY_Instruction),
				Quantity:    1,
			},
		},
	}
	o, err = Client.Orders.PlaceOrder(accountnumber, o)
	fmt.Println(util.SerializeReadable(o), err)

	o.Quantity = 2
	o.OrderLegCollection[0].Quantity = 2
	o = models.Order{
		OrderId:           o.OrderId,
		Quantity:          2,
		Price:             100,
		Session:           util.Ptr(models.SEAMLESS_Session),
		Duration:          util.Ptr(models.DAY_Duration),
		OrderType:         util.Ptr(models.LIMIT_OrderType),
		OrderStrategyType: util.Ptr(models.SINGLE_OrderStrategyType),
		OrderLegCollection: []models.OrderLegCollection{
			{
				Instrument: &models.AccountsInstrument{
					Symbol:    "AAPL",
					AssetType: "EQUITY",
				},
				Instruction: util.Ptr(models.BUY_Instruction),
				Quantity:    2,
			},
		},
	}
	o, err = Client.Orders.ReplaceOrder(accountnumber, o)
	fmt.Println(util.SerializeReadable(o), err)

	o, err = Client.Orders.GetOrder(accountnumber, o.OrderId)
	fmt.Println(util.SerializeReadable(o), err)

	q, err := Client.Orders.GetAllOrders(accountnumber)
	fmt.Println(util.SerializeReadable(q), err)

	err = Client.Orders.CancelOrder(accountnumber, o.OrderId)
	fmt.Println(err)

	o, err = Client.Orders.GetOrder(accountnumber, o.OrderId)
	fmt.Println(util.SerializeReadable(o), err)
}

func Test_Stream_L1_Equity(t *testing.T) {
	Client.StreamQuotes([]string{"AAPL"}, func(q *models.Quote) error {
		util.Log(q)
		return nil
	})

	for i := 0; i < 20; i++ {
		time.Sleep(100 * time.Millisecond)
	}
}
