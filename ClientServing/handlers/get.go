package handlers

import (
	"fmt"
	"net/http"

	protos "github.com/ellofae/Financial-Market-Microservice/APIs/CurrencyRates/Rates/protos/rates"
	"github.com/ellofae/Financial-Market-Microservice/ClientServing/data"
	"github.com/gofiber/fiber/v2"
	"github.com/hashicorp/go-hclog"
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

	objs := []data.CurrencyObject{}
	obj := data.CurrencyObject{}

	for amount := 0; amount < 5; amount++ {
		base := protos.Currencies_name[int32(amount)]
		resp, err := http.DefaultClient.Get(fmt.Sprintf("http://localhost:9092/rate?currency=%s", base))
		if err != nil {
			return err
		}

		if resp.StatusCode != http.StatusOK {
			g.log.Error("Recieved response with status code not 200", "recieved status code", resp.StatusCode)
		}
		defer resp.Body.Close()
		g.log.Info("Recieved response", "response", resp)

		obj = data.CurrencyObject{}
		err = obj.FromJSON(resp.Body)
		if err != nil {
			g.log.Error("Unable to unmarhsall requested data")
			return nil
		}
		objs = append(objs, obj)
	}

	g.log.Info("Recieved data", "data", objs)

	return c.Render("index", fiber.Map{
		"Currencies": objs,
	})

	return nil
}
