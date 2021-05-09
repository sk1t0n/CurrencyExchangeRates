package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/sk1t0n/CurrencyExchangeRates/api"
	"github.com/urfave/cli/v2"
)

func printCurrencyNames() {
	currencyNames := api.CurrencyNames()
	i := 0
	for ; i < len(currencyNames)-1; i++ {
		if currencyNames[i][0] == currencyNames[i+1][0] {
			fmt.Printf("%s\t", currencyNames[i])
		} else {
			fmt.Println(currencyNames[i])
		}
	}
	fmt.Println(currencyNames[i])
}

func printCurrencyExchangeRates(currency string) {
	rates, err := api.CurrencyRates(currency)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	keys := make([]string, len(rates))
	i := 0
	for k := range rates {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	for _, v := range keys {
		fmt.Printf("%s\t%f\n", v, rates[v])
	}
}

func main() {
	app := &cli.App{
		Name:  "currency",
		Usage: "if no currency is set print currency names",
		Action: func(c *cli.Context) error {
			if len(c.Args().Get(0)) == 0 {
				printCurrencyNames()
			} else {
				printCurrencyExchangeRates(c.Args().Get(0))
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
