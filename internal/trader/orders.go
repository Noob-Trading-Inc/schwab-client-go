package trader

import (
	"fmt"
	"net/url"
	"time"

	"github.com/Noob-Trading-Inc/schwab-client-go/internal"
	"github.com/Noob-Trading-Inc/schwab-client-go/models"
)

type Order struct {
	Session   string
	Duration  string
	OrderType string
}

type Orders struct {
}

func (c Orders) PlaceOrder(number string, order models.Order) (rv models.Order, err error) {
	url := fmt.Sprintf("%s/%s/%s/%s", internal.Endpoints.Trader, "accounts", number, "orders")
	location := ""
	location, err = internal.API.Post(url, order, nil)
	internal.API.Get(location, &rv)
	return
}

func (c Orders) ReplaceOrder(number string, order models.Order) (rv models.Order, err error) {
	url := fmt.Sprintf("%s/%s/%s/%s/%d", internal.Endpoints.Trader, "accounts", number, "orders", order.OrderId)
	location := ""
	location, err = internal.API.Put(url, order, nil)
	internal.API.Get(location, &rv)
	return
}

func (c Orders) CancelOrder(number string, orderId int64) (err error) {
	url := fmt.Sprintf("%s/%s/%s/%s/%d", internal.Endpoints.Trader, "accounts", number, "orders", orderId)
	err = internal.API.Delete(url, nil)
	return
}

func (c Orders) GetOrder(number string, orderId int64) (rv models.Order, err error) {
	url := fmt.Sprintf("%s/%s/%s/%s/%d", internal.Endpoints.Trader, "accounts", number, "orders", orderId)
	err = internal.API.Get(url, &rv)
	return
}

func (c Orders) GetAllOrders(number string) (rv []models.Order, err error) {
	var u *url.URL
	u, err = url.Parse(fmt.Sprintf("%s/%s/%s/%s", internal.Endpoints.Trader, "accounts", number, "orders"))
	if err != nil {
		return
	}
	query := u.Query()
	//query.Set("status", "WORKING")
	//query.Set("maxResults", "3000")
	query.Set("fromEnteredTime", time.Now().UTC().AddDate(0, -6, 0).Format(time.RFC3339))
	query.Set("toEnteredTime", time.Now().UTC().AddDate(0, 1, 0).Format(time.RFC3339))
	u.RawQuery = query.Encode()

	err = internal.API.Get(u.String(), &rv)
	return
}
