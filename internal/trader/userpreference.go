package trader

import (
	"fmt"
	"schwab-client-go/internal"
	"schwab-client-go/internal/trader/model"
)

type UserPreference struct {
}

func (c UserPreference) GetUserPreference() (rv model.UserPreference, err error) {
	url := fmt.Sprintf("%s/%s", internal.Endpoints.Trader, "userPreference")
	err = internal.API.Execute(url, &rv)
	return
}
