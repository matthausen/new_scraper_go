package router

import (
	"../handlers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/topic", handlers.GetTopic).Methods("POST", "OPTIONS")
	return router
}
