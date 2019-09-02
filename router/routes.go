package router

import (
	"net/http"

	"github.com/YAWAL/TodayInHistory/database"
	"github.com/YAWAL/TodayInHistory/handlers"

	"github.com/gorilla/mux"
)

func InitRouter(er database.HistoryRepository) (r *mux.Router) {
	r = mux.NewRouter()
	api := r.PathPrefix("/history").Subrouter()
	api.HandleFunc("/date", handlers.ShowTodayHistory(er)).Methods(http.MethodGet)
	api.HandleFunc("/date/{month}/{day}", handlers.ShowHistory(er)).Methods(http.MethodGet)
	return r
}
