package main

import (
	"fmt"
	"sync"
	"time"

	as "github.com/aerospike/aerospike-client-go"
	"github.com/prometheus/client_golang/prometheus"
)

// Scrap is aerospike info object
type Scrap map[string]string

// AerospikeExporter stores properties of aeropsike object
type AerospikeExporter struct {
	addr    string
	scraps  chan Scrap
	conn    *as.Connection
	metrics map[string]*metric
	sync.RWMutex
}

// NewAerospikeExporter initializes and return AerospikeExporter
func NewAerospikeExporter(addr string) (*AerospikeExporter, error) {
	conn, err := as.NewConnection(addr, time.Minute)
	if err != nil {
		return nil, err
	}

	return &AerospikeExporter{
		addr:    addr,
		scraps:  make(chan Scrap),
		conn:    conn,
		metrics: initMetrics(),
	}, nil
}

// Collect implements prometheus's Collector interface
func (e *AerospikeExporter) Collect(ch chan<- prometheus.Metric) {
	e.Lock()
	defer e.Unlock()

	scrapes := make(chan Scrap)

	go e.scrape(scrapes)
	e.process(scrapes)

	for _, m := range e.metrics {
		m.Collect(ch)
	}

	close(scrapes)
}

// Describe implements prometheus's Collector interface
func (e *AerospikeExporter) Describe(ch chan<- *prometheus.Desc) {
	for _, m := range e.metrics {
		m.Describe(ch)
	}
}

// scrape connects to aerospike and fetches info
func (e *AerospikeExporter) scrape(scrap chan<- Scrap) {
	if !e.conn.IsConnected() {
		conn, err := as.NewConnection(e.addr, time.Minute)
		if err != nil {
			return
		}
		e.conn = conn
	}

	info, err := as.RequestInfo(e.conn)
	if err != nil {
		fmt.Println(err)
	}

	scrap <- info
}

// process converts scraps to metrics
func (e *AerospikeExporter) process(scrap <-chan Scrap) {
	for scrap := range e.scraps {
		for k, v := range scrap {
			if m, ok := e.metrics[k]; ok {
				if m.metricType == constType {
					//
				}
			}
		}
	}
}
