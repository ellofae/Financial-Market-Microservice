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

import "github.com/ellofae/Financial-Market-Microservice/APIs/HistoryRecord/data"

// RecordCurrectResponse is a correct response from the server to the client with status code 200
// swagger:response recordCurrectResponse
type recordCurrectResponseWrapper struct {
	// Record object is the response of the server in JSON format
	// in: body
	Body data.Records
}

// RecordInternalErrorResponse is an Internal Error with status code 500 returned when the server didn't manage to get requested currency's rate
// swagger:response recordInternalErrorResponse
type recordInternalErrorResponseWrapper struct {
	// Record object is the response of the server in JSON format which was supposed to be made if no internal error occured
	// in: body
	Body data.Records
}
