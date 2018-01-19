package collector

import "github.com/prometheus/client_golang/prometheus"

const (
	namespace = "aerospike"
)

// GetMetricsInfo returns metrics info for given field
func GetMetricsInfo(field string) map[string]MetricInfo {
	switch field {
	case "":
		return basicMetricInfo
	case "statistics":
		return statsMetricInfo
	}

	return map[string]MetricInfo{}
}

// GetMetrics initializes metrics from metrics meta
func GetMetrics(meta MetricsInfo) Metrics {
	metrics := Metrics{}
	for name, info := range meta {
		metrics[name] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name:      info.Name,
			Help:      info.Help,
			Namespace: namespace,
		}, info.Labels)
	}
	return metrics
}
