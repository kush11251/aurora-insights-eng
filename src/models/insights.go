package models

import (
	"fmt"
)

type Insight struct {
	Name  string
	Value string
}

func GetInsights() []Insight {
	return []Insight{{Name: "sales_trend", Value: "increasing"}, {Name: "customer_satisfaction", Value: "high"}}
}