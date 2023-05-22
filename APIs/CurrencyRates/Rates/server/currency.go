package server

import (
	"context"
	"io"
	"time"

	"github.com/ellofae/Financial-Market-Microservice/APIs/CurrencyRates/Rates/data"
	protos "github.com/ellofae/Financial-Market-Microservice/APIs/CurrencyRates/Rates/protos/rates"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Currency struct {
	log    hclog.Logger
	rateDB *data.CurrencyDB
	cache  map[protos.CurrencyRates_StreamingCurrencyRateServer][]*protos.RatesRequest
}

func NewCurrency(log hclog.Logger, rate_db *data.CurrencyDB) *Currency {
	c := &Currency{log: log, rateDB: rate_db}
	go c.handleUpdates()

	return c
}

// send client updated cached requested currency rates that have been previously requested
func (c *Currency) handleUpdates() {
	updateStatusChannel := c.rateDB.MonitorRates(30 * time.Second)

	for range updateStatusChannel {
		c.log.Info("Currency rates are updated")

		for k, v := range c.cache {
			for _, rr := range v {
				c.log.Info("Subscriber", "base:", rr.Base)

				resp, err := c.rateDB.GetCurrencyRate(rr.Base.String())
				if err != nil {
					c.log.Error("Unable to get rate for the requested currency", "error", err)
				}

				err = k.Send(&protos.StreamingRatesResponse{
					Message: &protos.StreamingRatesResponse_RatesResponse{
						RatesResponse: resp,
					},
				})

				if err != nil {
					c.log.Error("Unable to send data to the client", "error", err)
				}
			}
		}
	}
}

// every requests server's data gets updated by calling an api
func (c *Currency) GetCurrencyRate(ctx context.Context, rr *protos.RatesRequest) (*protos.RatesResponse, error) {
	c.log.Info("Requesting data from the server", "currency", rr.Base)

	requestedObject, err := c.rateDB.GetCurrencyRate(rr.GetBase().String())
	if err != nil {
		c.log.Error("Unable to get rate from the server", "error", err)

		grpcErr := status.Newf(
			codes.Internal,
			"Unable to get rate from the server, currency requested: %s",
			rr.Base.String(),
		)

		grpcErr, wde := grpcErr.WithDetails(rr)
		if wde != nil {
			return nil, wde
		}

		return nil, grpcErr.Err()
	}

	return requestedObject, nil
}

func (c *Currency) StreamingCurrencyRate(src protos.CurrencyRates_StreamingCurrencyRateServer) error {
	c.log.Info("Processing with bi-directional streaming")

	for {
		resp, err := src.Recv()

		if err == io.EOF {
			c.log.Info("Client has closed the connection")
			break
		}

		if err != nil {
			c.log.Error("Unable to get data from the client", "error", err)

			grpcErr := status.Newf(
				codes.Internal,
				"Unable to get data on the server side from the client",
			)

			return grpcErr.Err()
		}

		c.log.Info("Handle client request", "request_base", resp.GetBase())

		subs, ok := c.cache[src]
		if !ok {
			c.log.Info("Creating cache for constant bi-directional streaming subscription updates")
			subs = []*protos.RatesRequest{}
		}

		subs = append(subs, resp)
		c.cache[src] = subs
	}

	return nil
}
