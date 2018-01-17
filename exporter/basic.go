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
			name:    "aerospike_node_name",
			help:    "Name of Aerospike node",
			labels:  []string{"addr", "alias", "id"},
		},
	}
)

// NewBasicExporter initializes BasicExporter
func NewBasicExporter(options Options) Exporter {
	return &BasicExporter{
		Addr:    options.Addr,
		Alias:   options.Alias,
		Metrics: initBasicMetrics(),
	}
}

// BasicExporter provides
type BasicExporter struct {
	Addr    string
	Alias   string
	Metrics map[string]*prometheus.GaugeVec
}

// Process accepts scrapes and processes it
func (e *BasicExporter) Process(scraps <-chan Scrap) {
	fmt.Println("Processing scraps....")
	for scrap := range scraps {
		fmt.Println("Scrap received:", scrap)
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
	fmt.Println("Processed scraps.")
}

// Extract accepts generates scrapes
func (e *BasicExporter) Extract(scraps chan<- Scrap) {
	defer close(scraps)

	fmt.Println("Extracting scraps....")
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
	fmt.Println("Extracted scraps.")
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

func initBasicMetrics() map[string]*prometheus.GaugeVec {
	metrics := map[string]*prometheus.GaugeVec{}

	for k, v := range basicMetricInfo {
		metrics[k] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: v.name,
			Help: v.help,
		}, v.labels)
	}

	return metrics
}
