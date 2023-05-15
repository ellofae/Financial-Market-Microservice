package server

import (
	"context"
	"time"

	"github.com/ellofae/Financial-Market-Microservice/CurrencyRates/data"
	protos "github.com/ellofae/Financial-Market-Microservice/CurrencyRates/protos/currency"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Currency struct {
	log   hclog.Logger
	rates *data.CurrencyRates
}

func NewCurrency(log hclog.Logger, r *data.CurrencyRates) *Currency {
	c := &Currency{log: log, rates: r}
	go c.handleUpdates()

	return c
}

func (c *Currency) handleUpdates() {
	updatesChannel := c.rates.MonitorRates(5 * time.Second)

	for range updatesChannel {
		c.log.Info("Rates updated")

		/*
			var src protos.Currency_SubscribeRatesServer

			for k, v := range c.rates.ReturnRates() {
				err := src.Send(
					&protos.StreamingRateResponse{
						Message: &protos.StreamingRateResponse_RateResponse{
							RateResponse: &protos.RateResponse{Base: protos.Currencies(protos.Currencies_value[k]), Rate: v},
						},
					},
				)

				if err != nil {
					c.log.Error("Unable to send updated rates", "base", k)
				}
			}
		*/
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

/*

func (c *Currency) SubscribeRates(src protos.Currency_SubscribeRatesServer) error {
	// handle client messages
	for {
		rr, err := src.Recv()

		// io.EOF signals that the client has closed the connection
		if err == io.EOF {
			c.log.Info("Client has closed connection")
			break
		}

		// transport between client and server is unavailable
		if err != nil {
			c.log.Error("Unable to read from client", "error", err)
			return err
		}

		c.log.Info("Handle client request", "request_base", rr.GetBase())

		rrs, ok := c.subscriptions[src]
		if !ok {
			rrs = []*protos.RateRequest{}
		}

		var validationError *status.Status
		for _, request := range rrs {
			if request.Base == rr.Base {
				validationError = status.Newf(
					codes.AlreadyExists,
					"Unable to subscribe for currency as subscription already exists",
				)

				validationError, err = validationError.WithDetails(rr)
				if err != nil {
					c.log.Error("Unable to add metadata to error", "error", err)
					break
				}

				break
			}
		}

		if validationError != nil {
			src.Send(&protos.StreamingRateResponse{
				Message: &protos.StreamingRateResponse_Error{
					Error: validationError.Proto(),
				},
			},
			)
			continue
		}

		rrs = append(rrs, rr)
		c.subscriptions[src] = rrs
	}

	return nil
}
*/
