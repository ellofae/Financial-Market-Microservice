package main

import (
	"time"

	"github.com/ellofae/Financial-Market-Microservice/ClientServing/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html"
	"github.com/hashicorp/go-hclog"
)

func main() {
	// setting logger
	log := hclog.Default()

	// engine setting
	engine := html.New("./pages", ".html")

	// setting fiber.App structure
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		Views:         engine,
		IdleTimeout:   120 * time.Second,
		ReadTimeout:   5 * time.Second,
		WriteTimeout:  10 * time.Second,
		AppName:       "ClientFinansialServing",
	})

	app.Use(cors.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins: "https:/localhost:9091",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// setting struct's with handlers methods
	gr := handlers.NewGetRouter(log)

	// setting handlers for serving client's requests:
	// setting http.MethodGet handlers
	app.Get("/", gr.GetGreetingPage)

	// starting server on port 3000
	log.Info("Starting the server", "host", "localhost", "port", 3000)
	app.Listen(":3000")
}
