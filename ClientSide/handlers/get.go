package handlers

import (
	"net/http"
)

func (c *Currencies) GetRate(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
	c.log.Info("Requesting currency rates by client side")

	base := r.URL.Query().Get("currency")

	currency, err := c.db.GetRates(base)
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
