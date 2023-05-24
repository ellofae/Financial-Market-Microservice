package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	protos "github.com/ellofae/Financial-Market-Microservice/APIs/CurrencyRates/Rates/protos/rates"
	"github.com/ellofae/Financial-Market-Microservice/ClientServing/data"
	"github.com/gofiber/fiber/v2"
	"github.com/hashicorp/go-hclog"
)

var CurrenciesAmountAvailable int = len(protos.Currencies_name)
var continueQuerySupport int = 0
var elementsToBeShown int = 0

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

func (g *GetRouter) ExchangePage(c *fiber.Ctx) error {
	g.log.Info("Sending greeting page to the client's request", "request's URL", c.Path)

	pageIndexQuery := c.Query("page")
	if pageIndexQuery != "" && pageIndexQuery != "" {
		index, err := strconv.Atoi(pageIndexQuery)
		if err != nil {
			g.log.Error("Unable to convert index query from strong to integer")
			return err
		}

		elementsToBeShown = 0
		elementsToBeShown = elementsToBeShown + 4*(index-1)
		continueQuerySupport = elementsToBeShown
	} else {
		continueQuery := c.Query("continue")
		if continueQuery != "" {
			elementsToBeShown = continueQuerySupport + 4
			continueQuerySupport += 4

			if elementsToBeShown+4 > CurrenciesAmountAvailable {
				elementsToBeShown = 0
				continueQuerySupport = 12
				g.log.Info("Redirecting to the exchange initial page")
				return c.Redirect("/exchange")
			}
		} else {
			previousQuery := c.Query("previous")
			if previousQuery != "" {

				if elementsToBeShown <= 0 {
					elementsToBeShown = 0
					continueQuerySupport = elementsToBeShown
					g.log.Info("Redirecting to the exchange initial page")
					return c.Redirect("/exchange")
				}

				elementsToBeShown = elementsToBeShown - 4
				continueQuerySupport -= 4
			}
		}
	}

	g.log.Info("Test", "elemtnsToBeShown", elementsToBeShown, "continuesQuerySupport", continueQuerySupport)

	objs := []data.CurrencyObject{}
	obj := data.CurrencyObject{}

	for amount := elementsToBeShown; amount < elementsToBeShown+4; amount++ {
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

	return c.Render("exchange", fiber.Map{
		"ExchangeRates": objs,
	})

	return nil
}
