package collector

import "github.com/prometheus/client_golang/prometheus"

// MetricType stores types of metric in given scrap
type MetricType int

const (
	_ MetricType = iota
	// String type metric
	String
	// Bool type metric
	Bool
	// Float type metric
	Float
	// Array type metric
	Array
	// Histogram type metric
	Histogram
)

// MetricInfo stores information about metric
type MetricInfo struct {
	MetricType MetricType
	Name       string
	Help       string
	Labels     []string
}

// MetricsInfo stores metrics meta
type MetricsInfo map[string]MetricInfo

// Scraps stores data exported from Aerospike
type Scraps map[string]string

// Metrics stores prometheus metric for aerospike fields
type Metrics map[string]*prometheus.GaugeVec

// Options stores collecter related meta
type Options struct {
	Alias string
	Addr  string
}
