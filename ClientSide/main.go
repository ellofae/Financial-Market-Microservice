package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ellofae/Financial-Market-Microservice/ClientSide/data"
	"github.com/ellofae/Financial-Market-Microservice/ClientSide/handlers"
	protos "github.com/ellofae/Financial-Market-Microservice/CurrencyRates/protos/currency"
	"github.com/go-openapi/runtime/middleware"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
)

func main() {
	log := hclog.Default()

	// Connection setting
	conn, err := grpc.Dial(":9090", grpc.WithInsecure())
	if err != nil {
		log.Error("Unable to extablish connection", "error", err)
		os.Exit(1)
	}
	defer conn.Close()

	// Currency rate client creation
	cc := protos.NewCurrencyClient(conn)

	// CurrenciesDB
	db := data.NewCurrencyDB(log, cc)

	// make connection between client and server
	pc := handlers.NewCurrencies(log, db)

	// create a new serve mux and register the handlers
	sm := mux.NewRouter()
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/rate", pc.GetRate).Queries("currency", "{[A-Z]{3}}")
	getRouter.HandleFunc("/rate", pc.GetRate)
	getRouter.HandleFunc("/rates", pc.GetAllRates)

	// swagger documentation handling
	ops := middleware.RedocOpts{SpecURL: "swagger.yaml"}
	sh := middleware.Redoc(ops, nil)
	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	// CORS
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"})) // as an open-api

	// server establishing
	srv := http.Server{
		Addr:         ":9091",
		Handler:      ch(sm),
		ErrorLog:     log.StandardLogger(&hclog.StandardLoggerOptions{}),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		log.Info("Starting server on port :9091")

		err := srv.ListenAndServe()
		if err != nil {
			log.Error("Unable to connect to the server", "error", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	log.Info("Got signal", "signal", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	srv.Shutdown(ctx)
}
