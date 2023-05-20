// Package classification of ClientSide API
//
// # Documentation for ClientSide API
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
// swagger:meta
package handlers

import (
	protos "github.com/ellofae/Financial-Market-Microservice/CurrencyRates/protos/currency"
)

// RateCorrectResponse is a Correct response from the server to the client with status code 200
// swagger:response rateCorrectResponse
type rateCorrectResponseWrapper struct {
	// RateResponse object which is the response of the server
	// in: body
	Body protos.RateResponse
}

// RateInternalErrorResponse is an Internal Error with status code 500 returned when the server didn't manage to get requested currency's rate
// swagger:response rateInternalErrorResponse
type rateInternalErrorResponseWrapper struct {
	// RateResponse object which is the response of the server which was supposed to be made if no internal error occured
	// in: body
	Body protos.RateResponse
}

// CurrencyQueryParam is a currency URI parameter
// swagger:parameters singleRate
type currencyQueryParam struct {
	// Currency used when requesting the rate of the currency
	// when currency is not specified an error is returned

	Currency string
}
