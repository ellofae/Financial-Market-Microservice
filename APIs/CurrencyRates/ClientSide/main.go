package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ellofae/Financial-Market-Microservice/APIs/CurrencyRates/ClientSide/data"
	"github.com/ellofae/Financial-Market-Microservice/APIs/CurrencyRates/ClientSide/handlers"
	protos "github.com/ellofae/Financial-Market-Microservice/APIs/CurrencyRates/Rates/protos/rates"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
)

func main() {
	// setting logger
	log := hclog.Default()

	conn, err := grpc.Dial(":9091", grpc.WithInsecure())
	if err != nil {
		log.Error("Unable to establish connection with API on port :9091", "error", err)
		os.Exit(1)
	}
	defer conn.Close()

	// currency rate client creation
	cc := protos.NewCurrencyRatesClient(conn)

	// handler interface implementation objects
	cdb := data.NewCurrencyData(log, cc)
	gh := handlers.NewCurrency(log, cdb)

	// create a new serve mux and register the handlers
	sm := mux.NewRouter()
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/rate", gh.GetCurrencyRateByBase).Queries("currency", "{[A-Z]{3}}")
	getRouter.HandleFunc("/rate", gh.GetCurrencyRateByBase)

	// swagger documentation handling
	/*
		ops := middleware.RedocOpts{SpecURL: "swagger.yaml"}
		sh := middleware.Redoc(ops, nil)
		getRouter.Handle("/docs", sh)
		getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))
	*/

	// CORS set up
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"})) // open api

	// server establishing
	srv := http.Server{
		Addr:         ":9092",
		Handler:      ch(sm),
		ErrorLog:     log.StandardLogger(&hclog.StandardLoggerOptions{}),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		log.Info("Starting client on port 9092")

		err := srv.ListenAndServe()
		if err != nil {
			log.Error("Unable to connect to the server", "error", err)
			os.Exit(1)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	log.Info("Got signal", "signal", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	srv.Shutdown(ctx)

}
