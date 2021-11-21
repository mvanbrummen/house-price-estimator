package main

import (
	"os"

	"github.com/go-resty/resty/v2"
)

func main() {
	client := resty.New()
	gateway := NewPropertyGateway("https://api.corelogic.asia", os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"), client)

	gateway.GetAccessToken()
}
