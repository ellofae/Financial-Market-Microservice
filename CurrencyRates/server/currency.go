package server

import (
	"context"
	"io"
	"time"

	"github.com/ellofae/Financial-Market-Microservice/CurrencyRates/data"
	protos "github.com/ellofae/Financial-Market-Microservice/CurrencyRates/protos/currency"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Currency struct {
	log           hclog.Logger
	rates         *data.CurrencyRates
	subscriprions map[protos.Currency_StreamingRatesServer][]*protos.RateRequest
}

func NewCurrency(log hclog.Logger, r *data.CurrencyRates) *Currency {
	c := &Currency{log: log, rates: r, subscriprions: map[protos.Currency_StreamingRatesServer][]*protos.RateRequest{}}
	go c.handleUpdates()

	return c
}

func (c *Currency) handleUpdates() {
	updatesChannel := c.rates.MonitorRates(15 * time.Second)

	for range updatesChannel {
		c.log.Info("Rates updated")

		for k, v := range c.subscriprions {
			for _, rr := range v {
				c.log.Info("Subscriber", "base:", rr.Base)

				r, err := c.rates.GetRates(rr.Base.String())
				if err != nil {
					c.log.Error("Unable to get update rate", "base", rr.GetBase().String())
				}

				err = k.Send(&protos.StreamingRateResponse{
					Message: &protos.StreamingRateResponse_RateResponse{
						RateResponse: &protos.RateResponse{Base: rr.Base, Rate: r},
					},
				})
				if err != nil {
					c.log.Error("Unable to send updated rate", "base", rr.GetBase().String())
				}
			}
		}
	}
}

func (c *Currency) GetRates(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Info("Sending a response to the client")

	if rr.Base.String() == "RUB" || rr.Base.String() == "" {
		c.log.Error("Invalid RateRequest.Base field data", "base", rr.Base.String())
		err := status.Newf(
			codes.InvalidArgument,
			"Base requested currency is %s, currency cannot be an empty string or RUB",
			rr.Base.String(),
		)

		err, wde := err.WithDetails(rr)
		if wde != nil {
			return nil, wde
		}

		return nil, err.Err()
	}

	rate, err := c.rates.GetRates(rr.GetBase().String())
	if err != nil {
		c.log.Error("Unable to get the rate", "currency", rr.Base.String())
		grpcError := status.Newf(
			codes.Internal,
			"Unable to get a rate for the requested currency %s, error message: %w",
			rr.Base.String(),
			err,
		)

		grpcError, wde := grpcError.WithDetails(rr)
		if wde != nil {
			return nil, wde
		}

		return nil, grpcError.Err()
	}

	return &protos.RateResponse{Base: protos.Currencies(protos.Currencies_value[rr.GetBase().String()]), Rate: rate}, nil
}

func (c *Currency) StreamingRates(src protos.Currency_StreamingRatesServer) error {
	c.log.Info("Connection is up by the streaming method")

	for {
		resp, err := src.Recv() // Recv is a blocking method which returns on client data

		// io.EOF signals that the client has closed the connection
		if err == io.EOF {
			c.log.Info("Client has closed the connection")
			break
		}

		// Unable to connect client and server
		if err != nil {
			c.log.Error("Unable to read from client", "error", err)
			return err
		}

		c.log.Info("Handle client request", "request_base", resp.GetBase())

		rrs, ok := c.subscriprions[src]
		if !ok {
			rrs = []*protos.RateRequest{}
		}

		rrs = append(rrs, resp)
		c.subscriprions[src] = rrs
	}

	return nil
}
