package exporter

import "github.com/prometheus/client_golang/prometheus"

// Exporter provides abstraction over aerospike metrices
type Exporter interface {
	Process(scrap <-chan Scrap)
	Extract(scrap chan<- Scrap)

	prometheus.Collector
}

// Options stores exporter related meta
type Options struct {
	Alias     string
	Addr      string
	Namespace string
}

// Scrap represent map of key-value pair from aeropsike info
type Scrap map[string]string

// ValueType stores types of value in give scrap
type ValueType int

const (
	_ ValueType = iota
	String
	Bool
	Float
	Array
	Histogram
)

// MetricInfo stores information about metric
type MetricInfo struct {
	valType   ValueType
	name      string
	help      string
	namespace string
	labels    []string
}
