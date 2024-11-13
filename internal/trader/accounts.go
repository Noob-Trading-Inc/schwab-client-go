package trader

import (
	"fmt"

	"github.com/Noob-Trading-Inc/schwab-client-go/internal"

	"github.com/Noob-Trading-Inc/schwab-client-go/internal/trader/model"
)

type Accounts struct {
}

var accountNumbers *[]model.AccountNumberHash

func (c Accounts) GetAccountNumbers() (rv []model.AccountNumberHash, err error) {
	if accountNumbers != nil {
		rv = *accountNumbers
		return
	}

	url := fmt.Sprintf("%s/%s/%s", internal.Endpoints.Trader, "accounts", "accountNumbers")
	err = internal.API.Get(url, &rv)

	if err == nil {
		accountNumbers = &rv
		internal.Token.DoOnTokenRefresh(func() {
			accountNumbers = nil
		})
	}
	return
}

func (c Accounts) GetAccounts() (rv []model.Account, err error) {
	url := fmt.Sprintf("%s/%s?fields=positions", internal.Endpoints.Trader, "accounts")
	err = internal.API.Get(url, &rv)
	return
}

func (c Accounts) GetAccount(number string) (rv model.Account, err error) {
	url := fmt.Sprintf("%s/%s/%s?fields=positions", internal.Endpoints.Trader, "accounts", number)
	err = internal.API.Get(url, &rv)
	return
}
