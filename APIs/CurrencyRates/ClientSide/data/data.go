package data

import (
	"context"
	"encoding/json"
	"io"

	protos "github.com/ellofae/Financial-Market-Microservice/APIs/CurrencyRates/Rates/protos/rates"
	"github.com/hashicorp/go-hclog"
)

// CurrencyData defines the structure for the client/server communication and servers handlers with data
// swagger:model
type CurrencyData struct {
	// server log object
	//
	// required: false
	log hclog.Logger

	// recieved currency rates data storage
	//
	// required: true
	Rates map[string]*CurrencyObject

	// CurrencyRatesClient object required for making calls to the server
	//
	// required: true
	currency protos.CurrencyRatesClient

	// CurrencyRates_StreamingCurrencyRateClient object required for client/server stream communication
	//
	// required: true
	client protos.CurrencyRates_StreamingCurrencyRateClient
}

// CurrencyObject defines the structure for an API currency object
// swagger:model
type CurrencyObject struct {
	Base    string
	Title   string
	NumCode string
	Rate    float64
}

func NewCurrencyData(log hclog.Logger, cc protos.CurrencyRatesClient) *CurrencyData {
	cdb := &CurrencyData{log: log, Rates: make(map[string]*CurrencyObject), currency: cc}

	go cdb.processUpdates()

	return cdb
}

func (cdb *CurrencyData) processUpdates() {
	sub, err := cdb.currency.StreamingCurrencyRate(context.Background())
	if err != nil {
		cdb.log.Error("Unable to subscribe for rates", "error", err)
		return
	}

	cdb.client = sub

	for {
		rr, err := sub.Recv()

		if grpcErr := rr.GetError(); grpcErr != nil {
			cdb.log.Error("Error subscribing for rates", "error", err)
			continue
		}

		cdb.log.Info("Recieved updated data from the server", "response", rr)

		if resp := rr.GetRatesResponse(); resp != nil {
			cdb.log.Info("Recieved updated rate from server", "base", resp.GetBase().String())
			if err != nil {
				cdb.log.Error("Error receiving message", "error", err)
				return
			}

			obj := &CurrencyObject{Base: protos.Currencies_name[int32(resp.Base)], Title: resp.Title, NumCode: resp.NumCode, Rate: resp.Rate}
			nobj := *obj
			cdb.Rates[resp.GetBase().String()] = &nobj
		}

	}
}

func (c *CurrencyData) GetCurrencyRate(base string) (*CurrencyObject, error) {
	currObj, err := c.getData(base)
	if err != nil {
		c.log.Error("Unable to get data for the requested currency", "currency", base, "error", err)
		return nil, err
	}

	return currObj, nil

}

func (c *CurrencyData) getData(base string) (*CurrencyObject, error) {
	c.log.Info("Requesting rate for currency", "currency", base)

	request := &protos.RatesRequest{Base: protos.Currencies(protos.Currencies_value[base])}

	rr, err := c.currency.GetCurrencyRate(context.Background(), request)
	if err != nil {
		c.log.Error("Unable to get data from the server for the requested currency", "error", err)
		return nil, err
	}

	c.client.Send(request)

	return &CurrencyObject{Base: base, Title: rr.Title, NumCode: rr.NumCode, Rate: rr.Rate}, nil
}

func (cobj *CurrencyObject) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(cobj)
}
