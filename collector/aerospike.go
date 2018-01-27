package collector

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Aerospike is aerospike exporter
type Aerospike struct {
	collectors []Collector
}

// Collect collects metrics from aerospike collectors
func (a *Aerospike) Collect(ch chan<- prometheus.Metric) {
	for _, collector := range a.collectors {
		collector.Collect(ch)
	}
}

// Describe describe metrics for aerospike collectors
func (a *Aerospike) Describe(ch chan<- *prometheus.Desc) {
	for _, collector := range a.collectors {
		collector.Describe(ch)
	}
}

// NewAerospike initializes collectors and return it
func NewAerospike(options Options) (*Aerospike, error) {
	// basicCollector, err := NewCollector(options, "")
	// if err != nil {
	// 	return nil, err
	// }

	// statsCollector, err := NewCollector(options, "statistics")
	// if err != nil {
	// 	return nil, err
	// }

	nsCollector, err := NewCollector(options, "namespaces")
	if err != nil {
		return nil, err
	}

	return &Aerospike{
		collectors: []Collector{
			// basicCollector,
			// statsCollector,
			nsCollector,
		},
	}, nil
}
