package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"reflect"
	"regexp"
	"sort"
	"strings"

	"github.com/sk1t0n/CurrencyExchangeRates/config"
	"github.com/valyala/fasthttp"
)

func getBytesByUrl(url string) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.SetRequestURI(url)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	err := fasthttp.Do(req, resp)
	if err != nil || resp.StatusCode() != fasthttp.StatusOK {
		return nil, err
	}

	contentType := resp.Header.Peek("Content-Type")
	if bytes.Index(contentType, []byte("application/json")) != 0 {
		return nil, err
	}

	body := resp.Body()
	return body, nil
}

type APIResponse struct {
	Result             string          `json:"result"`
	Documentation      string          `json:"documentation"`
	TermsOfUse         string          `json:"terms_of_use"`
	TimeLastUpdateUnix int             `json:"time_last_update_unix"`
	TimeLastUpdateUtc  string          `json:"time_last_update_utc"`
	TimeNextUpdateUnix int             `json:"time_next_update_unix"`
	TimeNextUpdateUtc  string          `json:"time_next_update_utc"`
	BaseCode           string          `json:"base_code"`
	ConversionRates    ConversionRates `json:"conversion_rates"`
}

type ConversionRates struct {
	USD float32
	AED float32
	AFN float32
	ALL float32
	AMD float32
	ANG float32
	AOA float32
	ARS float32
	AUD float32
	AWG float32
	AZN float32
	BAM float32
	BBD float32
	BDT float32
	BGN float32
	BHD float32
	BIF float32
	BMD float32
	BND float32
	BOB float32
	BRL float32
	BSD float32
	BTN float32
	BWP float32
	BYN float32
	BZD float32
	CAD float32
	CDF float32
	CHF float32
	CLP float32
	CNY float32
	COP float32
	CRC float32
	CUC float32
	CUP float32
	CVE float32
	CZK float32
	DJF float32
	DKK float32
	DOP float32
	DZD float32
	EGP float32
	ERN float32
	ETB float32
	EUR float32
	FJD float32
	FKP float32
	FOK float32
	GBP float32
	GEL float32
	GGP float32
	GHS float32
	GIP float32
	GMD float32
	GNF float32
	GTQ float32
	GYD float32
	HKD float32
	HNL float32
	HRK float32
	HTG float32
	HUF float32
	IDR float32
	ILS float32
	IMP float32
	INR float32
	IQD float32
	IRR float32
	ISK float32
	JMD float32
	JOD float32
	JPY float32
	KES float32
	KGS float32
	KHR float32
	KID float32
	KMF float32
	KRW float32
	KWD float32
	KYD float32
	KZT float32
	LAK float32
	LBP float32
	LKR float32
	LRD float32
	LSL float32
	LYD float32
	MAD float32
	MDL float32
	MGA float32
	MKD float32
	MMK float32
	MNT float32
	MOP float32
	MRU float32
	MUR float32
	MVR float32
	MWK float32
	MXN float32
	MYR float32
	MZN float32
	NAD float32
	NGN float32
	NIO float32
	NOK float32
	NPR float32
	NZD float32
	OMR float32
	PAB float32
	PEN float32
	PGK float32
	PHP float32
	PKR float32
	PLN float32
	PYG float32
	QAR float32
	RON float32
	RSD float32
	RUB float32
	RWF float32
	SAR float32
	SBD float32
	SCR float32
	SDG float32
	SEK float32
	SGD float32
	SHP float32
	SLL float32
	SOS float32
	SRD float32
	SSP float32
	STN float32
	SYP float32
	SZL float32
	THB float32
	TJS float32
	TMT float32
	TND float32
	TOP float32
	TRY float32
	TTD float32
	TVD float32
	TWD float32
	TZS float32
	UAH float32
	UGX float32
	UYU float32
	UZS float32
	VES float32
	VND float32
	VUV float32
	WST float32
	XAF float32
	XCD float32
	XDR float32
	XOF float32
	XPF float32
	YER float32
	ZAR float32
	ZMW float32
}

func CurrencyNames() []string {
	rates := ConversionRates{}
	v := reflect.ValueOf(rates)
	names := make([]string, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		names[i] = v.Type().Field(i).Name
	}
	sort.Strings(names)
	return names
}

func checkToken() bool {
	result := true
	if config.Token == "" {
		result = false
	}
	return result
}

func CurrencyRates(currency string) (map[string]float32, error) {
	if !checkToken() {
		return nil, errors.New("environment variable CURRENCY_EXCHANGE_RATES_TOKEN is not set")
	}

	re, _ := regexp.Compile("^[A-Za-z]{3}$")
	match := re.FindString(currency)
	if match == "" {
		return nil, errors.New("invalid currency")
	}

	url := config.URL + strings.ToUpper(currency)
	data, err := getBytesByUrl(url)
	if err != nil || len(data) == 0 {
		return nil, errors.New("failed to load data")
	}

	var response APIResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, errors.New("failed to parse json data")
	}

	v := reflect.ValueOf(response.ConversionRates)
	result := make(map[string]float32, v.NumField())
	zeros := 0
	for i := 0; i < v.NumField(); i++ {
		currency := v.Field(i).Interface().(float32)
		result[v.Type().Field(i).Name] = currency
		if currency == 0 {
			zeros++
		}
	}
	if zeros == v.NumField() {
		return nil, errors.New("invalid currency")
	}
	return result, nil
}
