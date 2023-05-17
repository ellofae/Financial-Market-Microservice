package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hashicorp/go-hclog"
	//"html/template"
)

// GetRouter is a structure that registers all the http.MethodGet handlers
type GetRouter struct {
	log hclog.Logger
}

func NewGetRouter(log hclog.Logger) *GetRouter {
	gr := &GetRouter{log: log}
	return gr
}

// GetGreetingPage is a handler that provides the greeting page to the client
func (g *GetRouter) GetGreetingPage(c *fiber.Ctx) error {
	g.log.Info("Sending greeting page to the client's request", "request's URL", c.Path)

	// c.SendFile("./test.html")
	c.Render("./test.html", fiber.Map{})

	return nil
}
