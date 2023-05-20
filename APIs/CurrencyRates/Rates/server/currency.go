package server

import (
	"context"

	"github.com/ellofae/Financial-Market-Microservice/APIs/CurrencyRates/Rates/data"
	protos "github.com/ellofae/Financial-Market-Microservice/APIs/CurrencyRates/Rates/protos/rates"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Currency struct {
	log    hclog.Logger
	rateDB *data.CurrencyDB
}

func NewCurrency(log hclog.Logger, rate_db *data.CurrencyDB) *Currency {
	return &Currency{log: log, rateDB: rate_db}
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

/*
func (c *Currency) StreamingCurrencyRate(src protos.CurrencyRates_StreamingCurrencyRateServer) error {
	return nil
}
*/
