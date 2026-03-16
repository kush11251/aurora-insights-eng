package controllers

import (
	"aurora-insights-eng/src/models"
	"encoding/json"
	"net/http"
)

func InsightsHandler(w http.ResponseWriter, r *http.Request) {
	insights := models.GetInsights()
	json.NewEncoder(w).Encode(insights)
}