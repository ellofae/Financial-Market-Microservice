package data

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	protos "github.com/ellofae/Financial-Market-Microservice/CurrencyRates/protos/currency"
	"github.com/hashicorp/go-hclog"
)

// Currency defines the structure for an API currency object
// swagger:model
type Currency struct {
	// the currency title
	//
	// required: true
	Base string `json:"base"`

	// the currency rate
	//
	// required: true
	Rate float64 `json:"rate"`
}

// CurrencyDB defines the structure for the client/server communication and servers handlers with data
// swagger:model
type CurrencyDB struct {
	// server log object
	//
	// required: false
	log hclog.Logger

	// recieved currency rates data storage
	//
	// required: true
	rates map[string]float64

	// CurrencyClient object required for making calls to the server
	//
	// required: true
	currency protos.CurrencyClient

	// Currency_StreamingRatesClient object required for client/server stream communication
	//
	// required: true
	client protos.Currency_StreamingRatesClient
}

// NewCurrencyDB created Currency object with specified logger and Client
func NewCurrencyDB(l hclog.Logger, cc protos.CurrencyClient) *CurrencyDB {
	c := &CurrencyDB{log: l, rates: make(map[string]float64), currency: cc}

	go c.processUpdates()

	return c
}

// processUpdates is a function that works in the background
// which updates the local data storage of currency rates when
// the updated rates are recieved from the server
func (c *CurrencyDB) processUpdates() {
	sub, err := c.currency.StreamingRates(context.Background())
	if err != nil {
		c.log.Error("Unable to subscribe for rates", "error", err)
		return
	}
	c.client = sub

	for {
		rr, err := sub.Recv()
		c.log.Info("Recieved updated data from the server", "response", rr)

		if grpcError := rr.GetError(); grpcError != nil {
			c.log.Error("Error subscribing for rates", "error", err)
			continue
		}

		if resp := rr.GetRateResponse(); resp != nil {
			c.log.Info("Recieved updated rate from server", "base", resp.GetBase().String())
			if err != nil {
				c.log.Error("Error receiving message", "error", err)
				return
			}

			c.rates[resp.GetBase().String()] = resp.Rate
		}
	}
}

// GetSingleRate returns reference to the Currency object with currency rate
// recieved from the server for the requested currency
func (c *CurrencyDB) GetSingleRate(currency string) (*Currency, error) {
	if currency == "" {
		c.log.Error("Unable to process the request because currency is not in the correct form")
		return nil, fmt.Errorf("not correct currency format, requested: %s", currency)
	}

	rate, err := c.getRate(currency)
	if err != nil {
		c.log.Error("Unable to get rate for the requested currency", "error", err)
		return nil, err
	}

	cur := &Currency{Base: currency, Rate: rate}

	return cur, nil
}

func (c *CurrencyDB) getRate(currency string) (float64, error) {
	c.log.Info("Requesting rate for currency", "currency", currency)

	request := &protos.RateRequest{Base: protos.Currencies(protos.Currencies_value[currency])}

	rr, err := c.currency.GetRates(context.Background(), request)
	if err != nil {
		c.log.Error("Unable to get rates by sending the request to the server", "error", err)
		return -1, err
	}

	c.client.Send(request)

	return rr.Rate, nil
}

// ToJSON converts Currency data to the JSON format and sends
// the converted data as the response
func (p *Currency) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(p)
}

// FromJSON converts Currency data from the JSON format and sends
// the converted data as the response
func (p *Currency) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(p)
}
