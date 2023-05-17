package main

import (
	"time"

	"github.com/ellofae/Financial-Market-Microservice/ClientServing/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/hashicorp/go-hclog"
)

func main() {
	// setting logger
	log := hclog.Default()

	// setting fiber.App structure
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		IdleTimeout:   120 * time.Second,
		ReadTimeout:   5 * time.Second,
		WriteTimeout:  10 * time.Second,
		AppName:       "ClientFinansialServing",
	})

	// setting struct's with handlers methods
	gr := handlers.NewGetRouter(log)

	// setting handlers for serving client's requests:
	// setting http.MethodGet handlers
	app.Get("/", gr.GetGreetingPage)

	app.Listen(":3000")
}
