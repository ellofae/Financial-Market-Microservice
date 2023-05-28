package data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hashicorp/go-hclog"
)

// Records defines the structure for an API response object
// swagger:model
type Records = map[string]interface{}

// RecordData defines the wrapper structure
// swagger:model
type RecordData struct {
	// server log object
	//
	// required: true
	log hclog.Logger

	// recieved record data
	//
	// required: true
	RecordsObj Records
}

// Creates a pointer to RecordData object
func NewRecordData(log hclog.Logger) *RecordData {
	return &RecordData{log: log, RecordsObj: make(map[string]interface{})}
}

// GetCurrencyHistory is the function that requests record data for the requested currency and interval
func (r *RecordData) GetCurrencyHistory(symbols string, start_date string, end_date string) error {
	err := r.getRecords(symbols, start_date, end_date)
	if err != nil {
		r.log.Error("Unable to get records", "error", err)
		return err
	}

	return nil
}

// getRecords is a helper function that makes a call to the API to recieve JSON format record data
func (r *RecordData) getRecords(symbols string, start_date string, end_date string) error {
	resp, err := RecieveAPIsHistoryRecord(symbols, start_date, end_date)
	if err != nil {
		r.log.Error("Unable to get data from API", "error", err)
	}

	if resp.StatusCode != http.StatusOK {
		r.log.Error("Recieved response's status code is not 200, unable to get rate from recieved body's response")
		return fmt.Errorf("recieved response's status code is not 200")
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&r.RecordsObj)
	if err != nil {
		r.log.Error("Unable to decode JSON data", "error", err)
		return err
	}

	return nil
}

// RecieveAPIsHistoryRecord returns the response made by the call to the API
func RecieveAPIsHistoryRecord(symbols string, start_date string, end_date string) (*http.Response, error) {
	resp, err := http.DefaultClient.Get(fmt.Sprintf("https://api.exchangerate.host/timeseries?base=RUB&symbols=%s&start_date=%s&end_date=%s", symbols, start_date, end_date))
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// ToJSON converts data to JSON format and sends it to the client
func ToJSON(rw io.Writer, data Records) error {
	encoder := json.NewEncoder(rw)
	return encoder.Encode(data)
}
