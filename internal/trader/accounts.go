package trader

import (
	"fmt"
	"schwab-client-go/internal"
)

type Account struct {
	AccountNumber string
	HashValue     string
}

type Accounts struct {
}

func (c Accounts) GetAccountNumbers() (rv []Account, err error) {
	url := fmt.Sprintf("%s/%s/%s", internal.Endpoints.Trader, "accounts", "accountNumbers")
	err = internal.API.Execute(url, &rv)
	return
}
