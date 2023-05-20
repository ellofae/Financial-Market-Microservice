package data

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/hashicorp/go-hclog"

	protos "github.com/ellofae/Financial-Market-Microservice/APIs/CurrencyRates/Rates/protos/rates"
)

type CurrencyDB struct {
	log   hclog.Logger
	rates map[string]*Rate
}

func NewCurrencyDB(log hclog.Logger) *CurrencyDB {
	c := &CurrencyDB{log: log, rates: make(map[string]*Rate)}
	go func() {
		err := c.getRates()
		if err != nil {
			log.Error("Unable to data for the server")
			os.Exit(1)
		}
	}()

	return c
}

func (c *CurrencyDB) GetCurrencyRate(base string) (*protos.RatesResponse, error) {
	if base == "" {
		c.log.Error("Incorrect currency's base data", "currency base", base)
		return nil, fmt.Errorf("currency's base is not supposed to be an empty string")
	}

	rateObj, ok := c.rates[base]
	if !ok {
		c.log.Error("There is no such currency rate in the database", "currency", base)
		return nil, fmt.Errorf("no such currency rate in the database")
	}

	rateObj.Rate = strings.Replace(rateObj.Rate, ",", ".", 1)
	rate, err := strconv.ParseFloat(rateObj.Rate, 64)
	if err != nil {
		c.log.Error("Unable to convert string data type to float64 data type")
		return nil, fmt.Errorf("unable to convert from string to float64")
	}

	respObj := &protos.RatesResponse{Base: protos.Currencies(protos.Currencies_value[rateObj.Base]), Rate: rate, NumCode: rateObj.NumCode, Title: rateObj.Name}

	return respObj, nil
}

func (c *CurrencyDB) getRates() error {
	resp, err := RecieveAPIsData()
	if err != nil {
		c.log.Error("Unable to get data from the API", "error", err)
		return err
	}

	if resp.StatusCode != http.StatusOK {
		c.log.Error("Recieved response's status code is not 200, unable to get rate from recieved body's response")
		return fmt.Errorf("recieved response's status code is not 200")
	}
	defer resp.Body.Close()

	rr := &Rates{}
	err = xml.NewDecoder(resp.Body).Decode(&rr)
	if err != nil {
		c.log.Error("Unable to decode XML data", "error", err)
		return err
	}

	for _, obj := range rr.Rates {
		nobj := obj
		c.rates[obj.Base] = &nobj
	}

	return nil
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
