package data

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/hashicorp/go-hclog"
)

type CurrencyRates struct {
	log   hclog.Logger
	rates map[string]float64
}

func NewCurrencyRates(log hclog.Logger) (*CurrencyRates, error) {
	c := &CurrencyRates{log: log, rates: map[string]float64{}}
	err := c.getRates()
	if err != nil {
		log.Error("Didn't manage to get rates", "error", err)
		return nil, err
	}

	return c, nil
}

func (c *CurrencyRates) ReturnRates() map[string]float64 {
	return c.rates
}

func (c *CurrencyRates) MonitorRates(interval time.Duration) chan struct{} {
	ret := make(chan struct{})

	go func() {
		ticker := time.NewTicker(interval)

		for {
			select {
			case <-ticker.C:
				err := c.getRates()
				if err != nil {
					c.log.Error("Unable to update currency rates", "error", err)
				}

				ret <- struct{}{}
			}
		}
	}()

	return ret
}

func (c *CurrencyRates) GetRates(currency string) (float64, error) {
	c.log.Info("Receving currency rate", "currency", currency)

	if currency == "" {
		c.log.Error("Invalid requested currency format", "currency", currency)
		return -1, fmt.Errorf("invalid requested currency format")
	}

	rate, ok := c.rates[currency]
	if !ok {
		c.log.Error("Requested currency is not present in data the storage", "currency requested", currency)
		fmt.Printf("TESTING: %#v\n\n", len(c.rates))
		return -1, fmt.Errorf("no currency %s exists in the data storage", currency)
	}

	return rate, nil
}

func (c *CurrencyRates) getRates() error {
	resp, err := http.DefaultClient.Get("https://www.cbr-xml-daily.ru/latest.js")
	if err != nil {
		c.log.Error("Unable to get rates via API", "error", err)
		return err
	}

	if resp.StatusCode != http.StatusOK {
		c.log.Error("Expected status code 200", "recieved status code", resp.StatusCode)
		return fmt.Errorf("expected status code 200, recieved %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	rr := &Rates{}
	json.NewDecoder(resp.Body).Decode(&rr)

	if len(rr.Currs) == 0 {
		c.log.Error("Unable to decode JSON currency data")
		return fmt.Errorf("unable to decode JSON currency data, empty list")
	}

	for key, value := range rr.Currs {
		c.rates[key] = value
	}

	return nil
}

type Rates struct {
	Currs map[string]float64 `json:"rates"`
}
