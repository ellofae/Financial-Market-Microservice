package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ellofae/Financial-Market-Microservice/APIs/HistoryRecord/handlers"
	"github.com/go-openapi/runtime/middleware"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
)

func main() {
	log := hclog.Default()
	recordObject := handlers.NewRecord(log)

	sm := mux.NewRouter()
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/record/{symbols:[A-Z]{3}}", recordObject.GetCurrencyHistory)

	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	// CORS
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"})) // as an open-api

	srv := &http.Server{
		Addr:         ":9095",
		Handler:      ch(sm),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 3 * time.Second,
	}

	go func() {
		log.Info("Starting server on port 9095")
		err := srv.ListenAndServe()
		if err != nil {
			os.Exit(1)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	log.Info("Recived terminate, gracefil shutdown", "signal", sig)

	// Graceful shutdown
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	srv.Shutdown(tc)
}
