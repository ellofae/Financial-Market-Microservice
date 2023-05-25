package main

import (
	"net"
	"os"

	"github.com/ellofae/Financial-Market-Microservice/CurrencyRates/data"
	protos "github.com/ellofae/Financial-Market-Microservice/CurrencyRates/protos/currency"
	"github.com/ellofae/Financial-Market-Microservice/CurrencyRates/server"
	hclog "github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// logger implementation
	log := hclog.Default()

	// data.CurrencyRate creation
	cr, err := data.NewCurrencyRates(log)
	if err != nil {
		log.Error("Unable to create data.CurrencyRates object", "error", err)
		os.Exit(1)
	}

	// CurrencyServer creation
	cs := server.NewCurrency(log, cr)

	// grpc server creation
	gs := grpc.NewServer()

	// registring currency server
	protos.RegisterCurrencyServer(gs, cs)

	reflection.Register(gs)

	// Specifing a port:
	l, err := net.Listen("tcp", ":9093")
	if err != nil {
		log.Error("Unable to listen", "error", err)
		os.Exit(1)
	}

	log.Info("Starting server on port :9093")
	gs.Serve(l)
}
