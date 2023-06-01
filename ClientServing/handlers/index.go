package handlers

import (
	"fmt"
	"math"
	"net/http"
	"sync"
	"time"

	protos "github.com/ellofae/Financial-Market-Microservice/APIs/CurrencyRates/Rates/protos/rates"
	"github.com/ellofae/Financial-Market-Microservice/ClientServing/data"
	"github.com/gofiber/fiber/v2"
	"github.com/hashicorp/go-hclog"
)

var (
	CurrenciesAmountAvailable int = len(protos.Currencies_name)
	continueQuerySupport      int = 0
	elementsToBeShown         int = 0
)

var waitGroup sync.WaitGroup

// GetRouter is a structure that registers all the http.MethodGet handlers
type GetRouter struct {
	log        hclog.Logger
	RecordList []data.Records
}

func NewGetRouter(log hclog.Logger) *GetRouter {
	gr := &GetRouter{log: log, RecordList: []data.Records{}}
	return gr
}

// GetGreetingPage is a handler that provides the greeting page to the client
func (g *GetRouter) GetGreetingPage(c *fiber.Ctx) error {
	c.GetRespHeader("Content-Type")

	g.log.Info("Sending greeting page to the client's request", "request's URL", c.Path)

	objs := []data.CurrencyObject{}
	rates := []data.CurrencyRatesWithPercentage{}

	waitGroup.Add(1)
	go g.requestCurrencyExchangeRates(&objs)
	waitGroup.Add(1)
	go g.requestCurrencyRates(&rates)

	waitGroup.Wait()
	c.Render("index", fiber.Map{
		"Currencies": objs[:5],
		"Rates":      rates,
		"Year":       time.Now().Year(),
	})

	return c.SendStatus(fiber.StatusOK)
}

func (g GetRouter) requestCurrencyExchangeRates(objs *[]data.CurrencyObject) error {
	g.log.Info("Requesting currency exchange rates")

	obj := data.CurrencyObject{}

	for amount := 0; amount < CurrenciesAmountAvailable; amount++ {
		base := protos.Currencies_name[int32(amount)]
		resp, err := http.DefaultClient.Get(fmt.Sprintf("http://localhost:9092/rate?currency=%s", base))
		if err != nil {
			return err
		}

		if resp.StatusCode != http.StatusOK {
			g.log.Error("Recieved response with status code not 200", "recieved status code", resp.StatusCode)
		}
		defer resp.Body.Close()
		//g.log.Info("Recieved response", "response", resp)

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

func (g *GetRouter) requestCurrencyRates(rates *[]data.CurrencyRatesWithPercentage) error {
	g.log.Info("Requesting currency rates")

	ratesObject := data.CurrencyRates{}
	tmp := data.CurrencyRatesWithPercentage{}
	for amount := 0; amount < CurrenciesAmountAvailable; amount++ {
		base := protos.Currencies_name[int32(amount)]
		resp, err := http.DefaultClient.Get(fmt.Sprintf("http://localhost:9094/rate?currency=%s", base))
		if err != nil {
			return err
		}

		if resp.StatusCode != http.StatusOK {
			g.log.Error("Recieved response with status code not 200", "recieved status code", resp.StatusCode)
		}
		defer resp.Body.Close()
		//g.log.Info("Recieved response", "response", resp)

		ratesObject = data.CurrencyRates{}
		err = ratesObject.FromJSON(resp.Body)
		if err != nil {
			g.log.Error("Unable to unmarhsall requested data")
			return nil
		}
		tmp = data.CurrencyRatesWithPercentage{Base: ratesObject.Base, Rate: ratesObject.Rate, Curr: math.Round((1/ratesObject.Rate)*100) / 100}
		ratesObject.Rate = math.Round(ratesObject.Rate*100000) / 100000
		*rates = append(*rates, tmp)
	}

	waitGroup.Done()

	return nil
}
