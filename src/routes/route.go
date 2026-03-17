package routes

import (
	"github.com/gorilla/mux"
)

func RouteHandler(router *mux.Router) {
	router.HandleFunc("/metrics", controllers.GetMetrics).Methods("GET")
}
