package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	namespace = "as"
)

// Type stores types of metric in given scrap
type Type int

const (
	_ Type = iota
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

// Info stores information about metric
type Info struct {
	Type      Type
	Name      string
	Help      string
	Subsystem string
	Labels    []string
}

// Metrics stores prometheus metric for aerospike fields
type Metrics map[string]*prometheus.GaugeVec

// GetMetricsInfo returns metrics info for given field
func GetMetricsInfo(field string) map[string]Info {
	switch field {
	case "":
		return basicInfo
	case "statistics":
		return statsInfo
	case "namespaces":
		return namespaceInfo
	}

	return map[string]Info{}
}

// GetMetrics initializes metrics from metrics meta
func GetMetrics(meta map[string]Info) Metrics {
	metrics := Metrics{}
	for name, info := range meta {
		metrics[name] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name:      info.Name,
			Help:      info.Help,
			Namespace: namespace,
			Subsystem: info.Subsystem,
		}, info.Labels)
	}
	return metrics
}
