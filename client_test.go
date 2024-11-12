package schwab

import (
	"fmt"
	"testing"
	"time"

	"github.com/Noob-Trading-Inc/schwab-client-go/internal/stream/model"
	"github.com/Noob-Trading-Inc/schwab-client-go/util"
)

func Test_Simple(t *testing.T) {
	q, err := Client.Quotes.GetQuote("TSLA")
	fmt.Println(util.SerializeReadable(q), err)

	a, err := Client.Acounts.GetAccountNumbers()
	fmt.Println(util.SerializeReadable(a), err)

	up, err := Client.UserPreference.GetUserPreference()
	if err != nil {
		util.OnError(err)
	}
	Client.Stream.Init(up)
	Client.Stream.Subscribe_L1_Equity("AAPL", func(err error, quote *model.TDWSResponse_L1_Content_Equity) {
		if err != nil {
			util.Log("ERROR Equity L1, %s", err.Error())
			return
		}
		util.Serialize(quote)
		util.Log(quote)
	})

	for i := 0; i < 20; i++ {
		time.Sleep(100 * time.Millisecond)
	}
}
