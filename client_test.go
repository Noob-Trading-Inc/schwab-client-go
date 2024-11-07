package main

import (
	"fmt"
	"schwab-client-go/util"
	"testing"
)

func Test_Simple(t *testing.T) {
	q, err := Client.Quotes.GetQuote("TSLA")
	fmt.Println(util.SerializeReadable(q), err)

	a, err := Client.Acounts.GetAccountNumbers()
	fmt.Println(util.SerializeReadable(a), err)
}
