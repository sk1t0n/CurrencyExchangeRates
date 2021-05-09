package config

import "os"

var Token string = os.Getenv("CURRENCY_EXCHANGE_RATES_TOKEN")
var URL string = "https://v6.exchangerate-api.com/v6/" + Token + "/latest/"
var CurrencyUSD = "USD"
