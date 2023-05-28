package handlers

import (
	"net/http"
	"regexp"

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
	interval := req.URL.Query().Get("interval")

	re := regexp.MustCompile("([0-9]{4}-[0-9]{2}-[0-9]{2})")
	found := re.FindAllString(interval, -1)
	if len(found) != 2 {
		r.log.Error("incorrect interval was requested", "recieved data", found)
	}

	start_date := found[0]
	end_date := found[1]

	r.log.Info("Requesting historical record", "symbols", symbols, "start_date", start_date, "end_date", end_date)

	err := r.records.GetCurrencyHistory(symbols, start_date, end_date)
	if err != nil {
		r.log.Error("Unable to get data from the API")
		http.Error(rw, "Unable to get data from the API request", http.StatusInternalServerError)
		return
	}

	err = data.ToJSON(rw, r.records.RecordsObj)
	if err != nil {
		r.log.Error("Unable to marshall data to JSON format", "error", err)
		http.Error(rw, "Unable to marshall data", http.StatusInternalServerError)
	}
}
