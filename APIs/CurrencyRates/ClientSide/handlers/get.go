package handlers

import (
	"net/http"

	"github.com/ellofae/Financial-Market-Microservice/APIs/CurrencyRates/ClientSide/data"
	"github.com/hashicorp/go-hclog"
)

type Currency struct {
	log  hclog.Logger
	data *data.CurrencyData
}

func NewCurrency(log hclog.Logger, dataObj *data.CurrencyData) *Currency {
	return &Currency{log: log, data: dataObj}
}

func (c *Currency) GetCurrencyRateByBase(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	base := r.URL.Query().Get("currency")
	c.log.Info("Base", "base", base)
	if base == "" {
		c.log.Error("Bad request: requested rate for the empty string")
		http.Error(rw, "Bad request: empty string was passed to the call", http.StatusBadRequest)
		return
	}

	c.log.Info("Requesting currency rate", "currency", base)
	currency, err := c.data.GetCurrencyRate(base)
	if err != nil {
		c.log.Error("Unable to get data from the server", "error", err)
		http.Error(rw, "Unable to get data from server", http.StatusInternalServerError)
		return
	}

	c.log.Info("Requsted data recieved", "currency requested", base, "response object", currency)

	err = currency.ToJSON(rw)
	if err != nil {
		c.log.Error("Unable to marhsall data to JSON format", "error", err)
		return
	}
}
