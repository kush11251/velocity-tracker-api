package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/metrics", getMetrics).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getMetrics(w http.ResponseWriter, r *http.Request) {
	metrics := []string{"metric1", "metric2", "metric3"}
	json.NewEncoder(w).Encode(metrics)
}
