package data

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/hashicorp/go-hclog"

	protos "github.com/ellofae/Financial-Market-Microservice/APIs/CurrencyRates/Rates/protos/rates"
)

type CurrencyDB struct {
	log   hclog.Logger
	rates map[string]float64
}

func NewCurrencyDB(log hclog.Logger) *CurrencyDB {
	c := &CurrencyDB{log: log, rates: make(map[string]float64)}

	return c
}

func (c *CurrencyDB) GetCurrencyRate(base string) (*protos.RatesResponse, error) {
	if base == "" {
		c.log.Error("Incorrect currency's base data", "currency base", base)
		return nil, fmt.Errorf("currency's base is not supposed to be an empty string")
	}

	rateObj, err := c.getRate(base)
	if err != nil {
		c.log.Error("Unable to get rate for the requested currency", "currency", base)
		return nil, fmt.Errorf("unable to get rate for the requested currency")
	}

	return rateObj, nil
}

func (c *CurrencyDB) getRate(base string) (*protos.RatesResponse, error) {
	resp, err := RecieveAPIsData()
	if err != nil {
		c.log.Error("Unable to get data from the API", "error", err)
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		c.log.Error("Recieved response's status code is not 200, unable to get rate from recieved body's response")
		return nil, fmt.Errorf("recieved response's status code is not 200")
	}
	defer resp.Body.Close()

	rr := &Rates{}
	err = xml.NewDecoder(resp.Body).Decode(&rr)
	if err != nil {
		c.log.Error("Unable to decode XML data", "error", err)
		return nil, err
	}

	//fmt.Printf("data: %#v\n", rr.Rates)

	var rateObj *protos.RatesResponse
	var rate float64

	for _, obj := range rr.Rates {
		obj.Rate = strings.Replace(obj.Rate, ",", ".", 1)
		rate, err = strconv.ParseFloat(obj.Rate, 64)
		if err != nil {
			fmt.Println("ERRORRRRR ", err)
			return nil, err
		}

		c.rates[obj.Base] = rate

		if obj.Base == base {
			rateObj = &protos.RatesResponse{Base: protos.Currencies(protos.Currencies_value[obj.Base]), Title: obj.Name, NumCode: obj.NumCode, Rate: rate}
		}
	}

	return rateObj, nil
}

func RecieveAPIsData() (*http.Response, error) {
	resp, err := http.DefaultClient.Get("https://www.cbr-xml-daily.ru/daily_utf8.xml")
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type Rates struct {
	Rates []Rate `xml:"Valute"`
}

type Rate struct {
	Base    string `xml:"CharCode"`
	Name    string `xml:"Name"`
	NumCode string `xml:"NumCode"`
	Rate    string `xml:"Value"`
}
