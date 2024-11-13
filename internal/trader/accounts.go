package trader

import (
	"fmt"

	"github.com/Noob-Trading-Inc/schwab-client-go/internal"
	"github.com/Noob-Trading-Inc/schwab-client-go/models"
)

type Accounts struct {
}

var accountNumbers *[]models.AccountNumberHash

func (c Accounts) GetAccountNumbers() (rv []models.AccountNumberHash, err error) {
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

func (c Accounts) GetAccounts() (rv []models.Account, err error) {
	url := fmt.Sprintf("%s/%s?fields=positions", internal.Endpoints.Trader, "accounts")
	err = internal.API.Get(url, &rv)
	return
}

func (c Accounts) GetAccount(number string) (rv models.Account, err error) {
	url := fmt.Sprintf("%s/%s/%s?fields=positions", internal.Endpoints.Trader, "accounts", number)
	err = internal.API.Get(url, &rv)
	return
}
