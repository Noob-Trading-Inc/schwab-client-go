package main

import (
	"fmt"
	"schwab-client-go/util"
	"testing"
)

func Test_Simple(t *testing.T) {
	q, err := Client.Quotes.GetQuote("TSLA")
	fmt.Println(util.Util.ToJsonReadable(q), err)

	a, err := Client.Acounts.GetAccountNumbers()
	fmt.Println(util.Util.ToJsonReadable(a), err)
}
