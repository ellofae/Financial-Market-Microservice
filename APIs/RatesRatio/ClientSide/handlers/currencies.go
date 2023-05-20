package handlers

import (
	"github.com/ellofae/Financial-Market-Microservice/ClientSide/data"
	"github.com/hashicorp/go-hclog"
)

type Currencies struct {
	log hclog.Logger
	db  *data.CurrencyDB // for client and server communication on transporting data
}

func NewCurrencies(log hclog.Logger, db *data.CurrencyDB) *Currencies {
	return &Currencies{log: log, db: db}
}
