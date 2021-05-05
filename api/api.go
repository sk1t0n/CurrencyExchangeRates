package api

import (
	"bytes"
	"encoding/json"
	"reflect"

	"github.com/sk1t0n/currency_exchange_rate/config"
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

func CurrencyRates(currency string) map[string]float32 {
	url := config.URL + currency
	data, err := getBytesByUrl(url)

	if err != nil {
		panic(err)
	}

	var response APIResponse
	err = json.Unmarshal(data, &response)

	if err != nil {
		panic(err)
	}

	v := reflect.ValueOf(response.ConversionRates)
	result := make(map[string]float32, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		result[v.Type().Field(i).Name] = v.Field(i).Interface().(float32)
	}
	return result
}
