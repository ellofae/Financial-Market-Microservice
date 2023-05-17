package handlers

import (
	"net/http"

	"github.com/ellofae/Financial-Market-Microservice/ClientSide/data"
	protos "github.com/ellofae/Financial-Market-Microservice/CurrencyRates/protos/currency"
)

// swagger:route GET /rate singleRate
//
// # Returns the rate for the requested currency recieved from the API call
//
// Responses:
// 		200: rateCorrectResponse
// 		500: rateInternalErrorResponse

// GetRate returns the rate for the requested currency passed to the function
func (c *Currencies) GetRate(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
	c.log.Info("Requesting currency rate by client side")

	base := r.URL.Query().Get("currency")

	currency, err := c.db.GetSingleRate(base)
	if err != nil {
		http.Error(rw, "Unable to get data from the server", http.StatusInternalServerError)
		return
	}

	c.log.Info("Requsted data recieved", "currency requested", base, "response object", currency)

	err = currency.ToJSON(rw)
	if err != nil {
		c.log.Error("Unable to serializing data to JSON format", "error", err)
		return
	}
}

// swagger:route GET /rates allRates
//
// # Returns rates of the all available currencies' rates that can be requested and recieved from the API
//
// Responses:
// 		200: rateCorrectResponse
// 		500: rateInternalErrorResponse

// GetAllRates rates of the all available currencies's rates, all the data is being taken from the server's map
func (c *Currencies) GetAllRates(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
	c.log.Info("Requesting currencies' rates by client side")

	var currencyObject *data.Currency
	var err error

	for currency_title, _ := range protos.Currencies_value {
		base := currency_title

		if r, ok := c.db.Rates[base]; !ok {
			currencyObject, err = c.db.GetSingleRate(base)
			if err != nil {
				http.Error(rw, "Unable to get data from the server", http.StatusInternalServerError)
				return
			}
		} else {
			currencyObject = &data.Currency{Base: base, Rate: r}
		}

		c.log.Info("Requsted data recieved", "currency requested", base, "response object", currencyObject)

		err = currencyObject.ToJSON(rw)
		if err != nil {
			c.log.Error("Unable to serializing data to JSON format", "error", err)
			return
		}
	}
}
