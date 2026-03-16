package controllers

import (
	"aurora-insights-eng/src/models"
	"encoding/json"
	"net/http"
)

func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	metrics := models.GetMetrics()
	json.NewEncoder(w).Encode(metrics)
}