package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
)

// AerospikeExporter implements prometheus collectors
type AerospikeExporter struct {
	Adrr      string
	Alias     string
	Namespace string
	Exporters []Exporter
}

// New initializes AerospikeExporter and returns it
func New(options Options) *AerospikeExporter {
	return &AerospikeExporter{
		Adrr:      options.Addr,
		Alias:     options.Alias,
		Namespace: options.Namespace,
		Exporters: []Exporter{
			NewBasicExporter(options),
		},
	}
}

// Collect collects prometheus's Collector interface
func (a *AerospikeExporter) Collect(ch chan<- prometheus.Metric) {
	for _, exporter := range a.Exporters {
		exporter.Collect(ch)
	}
}

// Describe collects prometheus's Collector interface
func (a *AerospikeExporter) Describe(ch chan<- *prometheus.Desc) {
	for _, exporter := range a.Exporters {
		exporter.Describe(ch)
	}
}
