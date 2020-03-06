package router

import (
	"github.com/gorilla/mux"
	"../handlers"
)

func Router() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/", handlers.index)
	router.HandleFunc("/api/topic", handlers.getTopic).Methods("POST", "OPTIONS")
	return router
}
