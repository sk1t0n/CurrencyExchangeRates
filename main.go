package main

import (
	"fmt"

	"github.com/sk1t0n/currency_exchange_rate/api"
)

func main() {
	currency := "USD"
	rates := api.CurrencyRates(currency)

	for k, v := range rates {
		fmt.Printf("k: %s\tv: %f\n", k, v)
	}
}
