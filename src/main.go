package main

import (
	"aurora-insights-eng/src/controllers"
	"aurora-insights-eng/src/models"
	"aurora-insights-eng/src/utils"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Aurora Insights Engine started")
	http.HandleFunc("/metrics", controllers.MetricsHandler)
	http.HandleFunc("/insights", controllers.InsightsHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}