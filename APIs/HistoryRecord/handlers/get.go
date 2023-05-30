package handlers

import (
	"net/http"

	"github.com/ellofae/Financial-Market-Microservice/APIs/HistoryRecord/data"
	"github.com/gorilla/mux"
)

// swagger:route GET /record singleRate
//
// # Returns the currency exchange rate record during some date interval
//
// Responses:
// 		200: recordCurrectResponse
// 		500: recordInternalErrorResponse

// GetCurrencyHistory returns the currency exchange rate record during some date interval
func (r *Record) GetCurrencyHistory(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	vars := mux.Vars(req)
	symbols := vars["symbols"]

	err := r.records.GetCurrencyHistory(symbols)
	if err != nil {
		r.log.Error("Unable to get data from the API")
		http.Error(rw, "Unable to get data from the API request", http.StatusInternalServerError)
		return
	}

	graphData := data.GraphFormatData{}

	for _, record := range r.records.RecordsObj {
		graphData.Dates = append(graphData.Dates, record.Date)
		graphData.Rates = append(graphData.Rates, record.Rates[symbols].(float64))
	}

	/*
		err = data.ToJSON(rw, r.records.RecordsObj)
		if err != nil {
			r.log.Error("Unable to marshall data to JSON format", "error", err)
			http.Error(rw, "Unable to marshall data", http.StatusInternalServerError)
		}
	*/
	err = graphData.ToJSON(rw)
	if err != nil {
		r.log.Error("Unable to marshall data to JSON format", "error", err)
		http.Error(rw, "Unable to marshall data", http.StatusInternalServerError)
	}
}
