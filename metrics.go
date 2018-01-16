package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	metricMap = map[string]metric{
		"node": metric(
			constType,
			"aerospike_node_id",
			"Aerospike node id"),
	}
)

type metricType int

const (
	_ metricType = iota
	gaugeType
	counterType
	gaugeVectorType
	constType
)

type metric struct {
	metricType metricType
	name       string
	desc       string
}

func initMetrics() map[string]*metric {
	metrics := map[string]*prometheus.GaugeVec{}

	return metrics
}
