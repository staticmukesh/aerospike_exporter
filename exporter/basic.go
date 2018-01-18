package exporter

import (
	"fmt"
	"strconv"
	"time"

	as "github.com/aerospike/aerospike-client-go"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	basicMetricInfo = map[string]MetricInfo{
		"node": MetricInfo{
			valType: String,
			name:    "node_name",
			help:    "Name of Aerospike node",
			labels:  []string{"addr", "alias", "id"},
		},
		"edition": MetricInfo{
			valType: String,
			name:    "edtion",
			help:    "Edition of Aerospike node",
			labels:  []string{"addr", "alias", "edition"},
		},
		"build": MetricInfo{
			valType: String,
			name:    "build",
			help:    "Build of Aerospike node",
			labels:  []string{"addr", "alias", "build"},
		},
		"partition-generation": MetricInfo{
			valType: Float,
			name:    "partition_generation",
			help:    "Partition generation of Aerospike node",
			labels:  []string{"addr", "alias"},
		},
		"cluster-generation": MetricInfo{
			valType: Float,
			name:    "cluster_generation",
			help:    "Cluster generation of Aerospike node",
			labels:  []string{"addr", "alias"},
		},
	}
)

// NewBasicExporter initializes BasicExporter
func NewBasicExporter(options Options) Exporter {
	return &BasicExporter{
		Addr:      options.Addr,
		Alias:     options.Alias,
		Namespace: options.Namespace,
		Metrics:   initBasicMetrics(options.Namespace),
	}
}

// BasicExporter provides
type BasicExporter struct {
	Addr      string
	Alias     string
	Namespace string
	Metrics   map[string]*prometheus.GaugeVec
}

// Process accepts scrapes and processes it
func (e *BasicExporter) Process(scraps <-chan Scrap) {
	for scrap := range scraps {
		for k, v := range scrap {
			if info, ok := basicMetricInfo[k]; ok {
				switch info.valType {
				case String:
					e.Metrics[k].WithLabelValues(e.Addr, e.Alias, v).Set(1)
				case Float:
					val, err := strconv.ParseFloat(v, 64)
					if err != nil {
						fmt.Println(err)
						continue
					}
					e.Metrics[k].WithLabelValues(e.Addr, e.Alias).Set(val)
				case Bool:
					val, err := strconv.ParseBool(v)
					if err != nil {
						fmt.Println(err)
						continue
					}
					if val {
						e.Metrics[k].WithLabelValues(e.Addr, e.Alias).Set(1)
					} else {
						e.Metrics[k].WithLabelValues(e.Addr, e.Alias).Set(0)
					}
				}
			}
		}
	}
}

// Extract accepts generates scrapes
func (e *BasicExporter) Extract(scraps chan<- Scrap) {
	defer close(scraps)

	conn, err := as.NewConnection(e.Addr, time.Minute)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	info, err := as.RequestInfo(conn)
	if err != nil {
		fmt.Println(err)
		return
	}

	scraps <- info
}

// Collect collects prometheus's Collector interface
func (e *BasicExporter) Collect(ch chan<- prometheus.Metric) {
	scrapes := make(chan Scrap)

	go e.Extract(scrapes)
	e.Process(scrapes)

	for _, metric := range e.Metrics {
		metric.Collect(ch)
	}
}

// Describe collects prometheus's Collector interface
func (e *BasicExporter) Describe(ch chan<- *prometheus.Desc) {
	for _, metric := range e.Metrics {
		metric.Describe(ch)
	}
}

func initBasicMetrics(namespace string) map[string]*prometheus.GaugeVec {
	metrics := map[string]*prometheus.GaugeVec{}

	for k, v := range basicMetricInfo {
		metrics[k] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name:      v.name,
			Help:      v.help,
			Namespace: namespace,
		}, v.labels)
	}

	return metrics
}
