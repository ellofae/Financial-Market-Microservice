package main

import (
	"net"
	"os"

	"github.com/ellofae/Financial-Market-Microservice/APIs/CurrencyRates/Rates/data"
	"github.com/ellofae/Financial-Market-Microservice/APIs/CurrencyRates/Rates/server"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"

	protos "github.com/ellofae/Financial-Market-Microservice/APIs/CurrencyRates/Rates/protos/rates"
	"google.golang.org/grpc/reflection"
)

func main() {
	// setting logger
	log := hclog.Default()

	// setting CurrencyDB
	db := data.NewCurrencyDB(log)

	// setting CurrencyRatesServer object
	cs := server.NewCurrency(log, db)

	// server configuration
	grpcServer := grpc.NewServer()

	protos.RegisterCurrencyRatesServer(grpcServer, cs)

	reflection.Register(grpcServer)

	// starting server on port 8000
	log.Info("Starting the server", "host", "localhost", "port", 8000)
	l, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Error("Unable to listen", "error", err)
		os.Exit(1)
	}

	grpcServer.Serve(l)
}
