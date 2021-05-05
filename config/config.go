package config

import "os"

var token string = os.Getenv("CURRENCY_EXCHANGE_RATE_TOKEN")
var URL string = "https://v6.exchangerate-api.com/v6/" + token + "/latest/"
