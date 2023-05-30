package data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/hashicorp/go-hclog"
)

// swagger:model
type GraphFormatData struct {
	Dates []string  `json:"Dates"`
	Rates []float64 `json:"Rates"`
}

// Records defines the structure for an API response object
// swagger:model
type Records struct {
	Date  string                 `json:"date"`
	Rates map[string]interface{} `json:"rates"`
}

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
	RecordsObj []Records
}

// Creates a pointer to RecordData object
func NewRecordData(log hclog.Logger) *RecordData {
	//	return &RecordData{log: log, RecordsObj: make(map[string]interface{})}
	return &RecordData{log: log, RecordsObj: []Records{}}
}

// GetCurrencyHistory is the function that requests record data for the requested currency and interval
func (r *RecordData) GetCurrencyHistory(symbols string) error {
	err := r.getRecords(symbols)
	if err != nil {
		r.log.Error("Unable to get records", "error", err)
		return err
	}

	return nil
}

// getRecords is a helper function that makes a call to the API to recieve JSON format record data
func (r *RecordData) getRecords(symbols string) error {
	recordData := []Records{}

	currentTime := time.Now()
	for i := 1; i < 25; i++ {
		record := Records{}

		resp, err := RecieveAPIsHistoryRecord(symbols, currentTime.Format("2006-01-02"))
		if err != nil {
			r.log.Error("Unable to get data from API", "error", err)
		}

		if resp.StatusCode != http.StatusOK {
			r.log.Error("Recieved response's status code is not 200, unable to get rate from recieved body's response")
			return fmt.Errorf("recieved response's status code is not 200")
		}
		defer resp.Body.Close()

		err = json.NewDecoder(resp.Body).Decode(&record)
		if err != nil {
			r.log.Error("Unable to unmarshall", "error", err)
			return err
		}

		record.Rates[symbols] = 1.0 / record.Rates[symbols].(float64)

		recordData = append(recordData, record)
		currentTime = currentTime.Add(-24 * time.Hour)
	}

	for i, j := 0, len(recordData)-1; i < j; i, j = i+1, j-1 {
		recordData[i], recordData[j] = recordData[j], recordData[i]
	}

	r.RecordsObj = recordData

	return nil
}

// RecieveAPIsHistoryRecord returns the response made by the call to the API
func RecieveAPIsHistoryRecord(symbols string, date string) (*http.Response, error) {
	resp, err := http.DefaultClient.Get(fmt.Sprintf("https://api.exchangerate.host/%s?symbols=%s&base=RUB", date, symbols))
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// ToJSON converts data to JSON format and sends it to the client
func ToJSON(rw io.Writer, data []Records) error {
	encoder := json.NewEncoder(rw)
	return encoder.Encode(data)
}

// ToJSON converts data to JSON format and sends it to the client
func (g *GraphFormatData) ToJSON(rw io.Writer) error {
	encoder := json.NewEncoder(rw)
	return encoder.Encode(g)
}
