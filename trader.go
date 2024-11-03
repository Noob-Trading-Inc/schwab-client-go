package main

import (
	"schwab-client-go/internal"
	"schwab-client-go/util"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	util.Util.Log("Starting")

	token := internal.Token.GetToken()
	util.Util.Log("Token : ", token)
}
