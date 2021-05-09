# CurrencyExchangeRates

CLI application for viewing currency exchange rates.

## Setup

Register on the website <https://www.exchangerate-api.com/>  
Set the environment variable CURRENCY_EXCHANGE_RATES_TOKEN

```sh
export CURRENCY_EXCHANGE_RATES_TOKEN=<website token exchangerate-api.com>
```

## Install

```sh
go get
go build -o CurrencyExchangeRates main.go
```

## Usage

```sh
# Help
./CurrencyExchangeRates help

# Currency names
./CurrencyExchangeRates

# Exchange rates, for example USD
./CurrencyExchangeRates USD
```
