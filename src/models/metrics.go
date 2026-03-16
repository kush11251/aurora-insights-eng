package models

import (
	"fmt"
)

type Metric struct {
	Name  string
	Value float64
}

func GetMetrics() []Metric {
	return []Metric{{Name: "cpu_usage", Value: 0.5}, {Name: "memory_usage", Value: 0.2}}
}