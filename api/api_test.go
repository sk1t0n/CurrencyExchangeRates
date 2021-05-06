package api

import (
	"testing"

	"github.com/sk1t0n/currency_exchange_rate/config"
)

func TestGetBytesByUrl(t *testing.T) {
	url := config.URL + config.CurrencyUSD
	data, err := getBytesByUrl(url)

	if err != nil {
		t.Fatalf("%v", err)
	}

	if len(data) == 0 {
		t.Fatalf("data is empty")
	}
}

func TestCurrencyRates(t *testing.T) {
	_, err := CurrencyRates("usd")

	if err != nil {
		t.Fatalf("%v", err)
	}
}
