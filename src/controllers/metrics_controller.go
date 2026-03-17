package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"velocity-tracker-api/src/models"
)

func GetMetrics(w http.ResponseWriter, r *http.Request) {
	metrics := []models.Metric{{Name: "metric1", Value: 10}, {Name: "metric2", Value: 20}}
	json.NewEncoder(w).Encode(metrics)
}
