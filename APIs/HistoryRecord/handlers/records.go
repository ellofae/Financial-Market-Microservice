package handlers

import (
	"github.com/ellofae/Financial-Market-Microservice/APIs/HistoryRecord/data"
	"github.com/hashicorp/go-hclog"
)

// RecordData defines the wrapper structure
// swagger:model
type Record struct {
	// server log object
	//
	// required: true
	log hclog.Logger

	// RecordData object for storing requested record
	//
	// required: true
	records data.RecordData
}

// NewRecord returns a poitner to Record object
func NewRecord(log hclog.Logger) *Record {
	return &Record{log: log}
}
