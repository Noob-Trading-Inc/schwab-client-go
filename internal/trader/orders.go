package trader

import (
	"fmt"

	"github.com/Noob-Trading-Inc/schwab-client-go/internal"
)

type Order struct {
}

/*
[
  {
    "session": "NORMAL",
    "duration": "DAY",
    "orderType": "MARKET",
    "cancelTime": "2024-11-05T06:16:42.721Z",
    "complexOrderStrategyType": "NONE",
    "quantity": 0,
    "filledQuantity": 0,
    "remainingQuantity": 0,
    "requestedDestination": "INET",
    "destinationLinkName": "string",
    "releaseTime": "2024-11-05T06:16:42.721Z",
    "stopPrice": 0,
    "stopPriceLinkBasis": "MANUAL",
    "stopPriceLinkType": "VALUE",
    "stopPriceOffset": 0,
    "stopType": "STANDARD",
    "priceLinkBasis": "MANUAL",
    "priceLinkType": "VALUE",
    "price": 0,
    "taxLotMethod": "FIFO",
    "orderLegCollection": [
      {
        "orderLegType": "EQUITY",
        "legId": 0,
        "instrument": {
          "cusip": "string",
          "symbol": "string",
          "description": "string",
          "instrumentId": 0,
          "netChange": 0,
          "type": "SWEEP_VEHICLE"
        },
        "instruction": "BUY",
        "positionEffect": "OPENING",
        "quantity": 0,
        "quantityType": "ALL_SHARES",
        "divCapGains": "REINVEST",
        "toSymbol": "string"
      }
    ],
    "activationPrice": 0,
    "specialInstruction": "ALL_OR_NONE",
    "orderStrategyType": "SINGLE",
    "orderId": 0,
    "cancelable": false,
    "editable": false,
    "status": "AWAITING_PARENT_ORDER",
    "enteredTime": "2024-11-05T06:16:42.721Z",
    "closeTime": "2024-11-05T06:16:42.721Z",
    "tag": "string",
    "accountNumber": 0,
    "orderActivityCollection": [
      {
        "activityType": "EXECUTION",
        "executionType": "FILL",
        "quantity": 0,
        "orderRemainingQuantity": 0,
        "executionLegs": [
          {
            "legId": 0,
            "price": 0,
            "quantity": 0,
            "mismarkedQuantity": 0,
            "instrumentId": 0,
            "time": "2024-11-05T06:16:42.721Z"
          }
        ]
      }
    ],
    "replacingOrderCollection": [
      "string"
    ],
    "childOrderStrategies": [
      "string"
    ],
    "statusDescription": "string"
  }
]
*/

type Orders struct {
}

func (c Orders) GetWorkingOrders(number string) (rv []Order, err error) {
	url := fmt.Sprintf("%s/%s/%s?fields=positions?maxResults=100&fromEnteredTime=2024-01-01T00%%3A00%%3A00Z&toEnteredTime=2025-01-01T00%%3A00%%3A00Z&status=WORKING", internal.Endpoints.Trader, "accounts", number)
	err = internal.API.Execute(url, &rv)
	return
}
