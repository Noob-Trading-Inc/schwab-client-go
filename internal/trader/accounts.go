package trader

import (
	"fmt"

	"github.com/Noob-Trading-Inc/schwab-client-go/internal"

	"github.com/Noob-Trading-Inc/schwab-client-go/internal/trader/model"
)

type Accounts struct {
}

func (c Accounts) GetAccountNumbers() (rv []model.AccountNumberHash, err error) {
	url := fmt.Sprintf("%s/%s/%s", internal.Endpoints.Trader, "accounts", "accountNumbers")
	err = internal.API.Execute(url, &rv)
	return
}

func (c Accounts) GetAccounts() (rv []model.Account, err error) {
	url := fmt.Sprintf("%s/%s?fields=positions", internal.Endpoints.Trader, "accounts")
	err = internal.API.Execute(url, &rv)
	return
}

func (c Accounts) GetAccount(number string) (rv model.Account, err error) {
	url := fmt.Sprintf("%s/%s/%s?fields=positions", internal.Endpoints.Trader, "accounts", number)
	err = internal.API.Execute(url, &rv)
	return
}
