package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"

	"github.com/YAWAL/TodayInHistory/database"
	"github.com/YAWAL/TodayInHistory/logging"
	"github.com/YAWAL/TodayInHistory/model"
)

const (
	todaysHistoryURL    = "https://history.muffinlabs.com/date"
	historyURL          = "https://history.muffinlabs.com/date/%v/%v"
	monthKey            = "month"
	dayKey              = "day"
	gettingDataError    = "error during getting history data: %s"
	processingDataError = "error during processing history data: %s"
	renderingDataError  = "error during rendering history data: %s"
)

// ShowTodayHistory performs Save/Update all data for each requested date in PostgreSQL and return
// titles of each category as well as events_count, birth_count, deaths_count for today's date
func ShowTodayHistory(hr database.HistoryRepository) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := request(todaysHistoryURL)
		if err != nil {
			logging.Log.Errorf(gettingDataError, err.Error())
			return
		}

		if err = renderData(&data, w); err != nil {
			logging.Log.Errorf(renderingDataError, err.Error())
			return
		}

		if err = hr.ProcessHistoryData(&data); err != nil {
			logging.Log.Errorf(processingDataError, err.Error())
			return
		}
	}
}

// ShowHistory performs Save/Update all data for each requested date in PostgreSQL and return
// titles of each category as well as events_count, birth_count, deaths_count for requested date
func ShowHistory(hr database.HistoryRepository) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		month := vars[monthKey]
		day := vars[dayKey]

		if !validDate(month, day) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		data, err := request(fmt.Sprintf(historyURL, month, day))
		if err != nil {
			logging.Log.Errorf(gettingDataError, err.Error())
			return
		}

		if err = renderData(&data, w); err != nil {
			logging.Log.Errorf(renderingDataError, err.Error())
			return
		}

		if err = hr.ProcessHistoryData(&data); err != nil {
			logging.Log.Errorf(processingDataError, err.Error())
			return
		}
	}
}

func renderData(hd *model.HistoryData, w http.ResponseWriter) error {
	t, err := template.ParseFiles("static/info.html")
	if err != nil {
		return err
	}
	titles := []string{
		fmt.Sprintf("Events count: %d", len(hd.Data.Events)),
		fmt.Sprintf("Birth count: %d", len(hd.Data.Births)),
		fmt.Sprintf("Deaths count: %d", len(hd.Data.Deaths)),
	}
	if err = t.Execute(w, titles); err != nil {
		return err
	}
	return nil
}

func request(url string) (historyData model.HistoryData, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return model.HistoryData{}, err
	}

	if err = json.NewDecoder(resp.Body).Decode(&historyData); err != nil {
		return model.HistoryData{}, err
	}
	return historyData, nil
}

func validDate(month, day string) bool {
	intMonth, err := strconv.Atoi(month)
	if intMonth > 12 || intMonth < 1 || err != nil {
		return false
	}
	dayMonth, err := strconv.Atoi(day)
	if dayMonth > 31 || dayMonth < 1 || err != nil {
		return false
	}
	return true
}
