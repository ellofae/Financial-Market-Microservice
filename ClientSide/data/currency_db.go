package data

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	protos "github.com/ellofae/Financial-Market-Microservice/CurrencyRates/protos/currency"
	"github.com/hashicorp/go-hclog"
)

type Currency struct {
	Base string  `json:"base"`
	Rate float64 `json:"rate"`
}

type CurrencyDB struct {
	log      hclog.Logger
	rates    map[string]float64
	currency protos.CurrencyClient
	client   protos.Currency_StreamingRatesClient
}

func NewCurrencyDB(l hclog.Logger, cc protos.CurrencyClient) *CurrencyDB {
	c := &CurrencyDB{log: l, rates: make(map[string]float64), currency: cc}

	go c.processUpdates()

	return c
}

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

func (c *CurrencyDB) GetRates(currency string) (*Currency, error) {
	if currency == "" {
		c.log.Error("Unable to process the request because currency is not in the correct form")
		return nil, fmt.Errorf("not correct currency format, requested: %s", currency)
	}

	rate, err := c.getRates(currency)
	if err != nil {
		c.log.Error("Unable to get rate for the requested currency", "error", err)
		return nil, err
	}

	cur := &Currency{Base: currency, Rate: rate}

	return cur, nil
}

func (c *CurrencyDB) getRates(currency string) (float64, error) {
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

func (p *Currency) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(p)
}

func (p *Currency) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(p)
}
