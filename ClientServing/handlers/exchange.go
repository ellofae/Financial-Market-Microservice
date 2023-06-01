package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	protos "github.com/ellofae/Financial-Market-Microservice/APIs/CurrencyRates/Rates/protos/rates"
	"github.com/ellofae/Financial-Market-Microservice/ClientServing/data"
	"github.com/gofiber/fiber/v2"
)

func (g *GetRouter) ExchangePage(c *fiber.Ctx) error {
	c.GetRespHeader("Content-Type")

	g.log.Info("Sending page with currency exchange rates to the client's request", "request's URL", c.Path)

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
		continueQuery := c.Query("next")
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

	objs := []data.CurrencyObject{}

	waitGroup.Add(1)
	go g.getExchangeRates(&objs)
	waitGroup.Wait()

	c.Render("exchange", fiber.Map{
		"ExchangeRates": objs,
	})

	return c.SendStatus(fiber.StatusOK)
}

func (g *GetRouter) getExchangeRates(objs *[]data.CurrencyObject) error {
	obj := data.CurrencyObject{}
	for amount := elementsToBeShown; amount < elementsToBeShown+12; amount++ {
		base := protos.Currencies_name[int32(amount)]
		resp, err := http.DefaultClient.Get(fmt.Sprintf("http://localhost:9092/rate?currency=%s", base))
		if err != nil {
			return err
		}

		if resp.StatusCode != http.StatusOK {
			g.log.Error("Recieved response with status code not 200", "recieved status code", resp.StatusCode)
		}
		defer resp.Body.Close()

		obj = data.CurrencyObject{}
		err = obj.FromJSON(resp.Body)
		if err != nil {
			g.log.Error("Unable to unmarhsall requested data")
			return nil
		}
		*objs = append(*objs, obj)
	}

	waitGroup.Done()
	return nil
}
