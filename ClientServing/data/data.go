package data

import (
	"encoding/json"
	"io"

	"github.com/hashicorp/go-hclog"
)

type CurrencyData struct {
	log hclog.Logger
}

func NewCurrencyData(log hclog.Logger) *CurrencyData {
	c := &CurrencyData{log: log}
	return c
}

type CurrencyObject struct {
	Base    string  `json:"Base"`
	Title   string  `json:"Title"`
	NumCode string  `json:"NumCode"`
	Rate    float64 `json:"Rate"`
}

func (c *CurrencyObject) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(c)
}
