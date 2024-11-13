package trader

import (
	"fmt"

	"github.com/Noob-Trading-Inc/schwab-client-go/internal"
	"github.com/Noob-Trading-Inc/schwab-client-go/models"
)

type UserPreference struct {
}

func (c UserPreference) GetUserPreference() (rv models.UserPreference, err error) {
	url := fmt.Sprintf("%s/%s", internal.Endpoints.Trader, "userPreference")
	err = internal.API.Get(url, &rv)
	return
}
