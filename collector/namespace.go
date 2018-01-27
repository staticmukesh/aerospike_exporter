package collector

import (
	"errors"
	"strconv"
	"strings"
	"time"

	aerospike "github.com/aerospike/aerospike-client-go"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/staticmukesh/aerospike_exporter/metrics"
)

type namespaceCollector struct {
	c collector
}

func (n *namespaceCollector) Collect(ch chan<- prometheus.Metric) {
	logger.Println("Collecting metrics")
	namespaces, err := n.getNamespaces()
	if err != nil {
		return
	}

	for _, namespace := range namespaces {
		scraps, err := n.extract(namespace)
		if err != nil {
			return
		}

		n.process(namespace, scraps)

		for _, metric := range n.c.metrics {
			metric.Collect(ch)
		}
	}
}

func (n *namespaceCollector) Describe(ch chan<- *prometheus.Desc) {
	logger.Println("Describing metrics", len(n.c.metrics))
	n.c.Describe(ch)
}

func (n *namespaceCollector) getNamespaces() ([]string, error) {
	conn, err := aerospike.NewConnection(n.c.addr, 10*time.Second)
	if err != nil {
		logger.Println("Aerospike error:", err)
		return nil, err
	}
	defer conn.Close()

	data, err := aerospike.RequestInfo(conn, n.c.field)
	if err != nil {
		logger.Println("Aerospike error:", err)
		return nil, err
	}

	namespacesStr, ok := data[n.c.field]
	if !ok {
		return nil, errors.New("No namespace found")
	}

	return strings.Split(namespacesStr, ";"), nil
}

func (n *namespaceCollector) extract(namespace string) (Scraps, error) {
	conn, err := aerospike.NewConnection(n.c.addr, 10*time.Second)
	if err != nil {
		logger.Println("Aerospike error:", err)
		return nil, err
	}
	defer conn.Close()

	data, err := aerospike.RequestInfo(conn, "namespace/"+namespace)
	if err != nil {
		logger.Println("Aerospike error:", err)
		return nil, err
	}

	logger.Println(data)

	return data, nil
}

func (n *namespaceCollector) process(namespace string, scraps Scraps) {
	n.c.Lock()
	defer n.c.Unlock()

	for k, v := range scraps {
		if info, ok := n.c.meta[k]; ok {
			switch info.Type {
			case metrics.String:
				n.c.metrics[k].WithLabelValues(n.c.addr, n.c.alias, v, namespace).Set(1)
			case metrics.Float:
				val, err := strconv.ParseFloat(v, 64)
				if err != nil {
					logger.Println(err)
					continue
				}
				n.c.metrics[k].WithLabelValues(n.c.addr, n.c.alias, namespace).Set(val)
			case metrics.Bool:
				val, err := strconv.ParseBool(v)
				if err != nil {
					logger.Println(err)
					continue
				}
				if val {
					n.c.metrics[k].WithLabelValues(n.c.addr, n.c.alias, namespace).Set(1)
				} else {
					n.c.metrics[k].WithLabelValues(n.c.addr, n.c.alias, namespace).Set(0)
				}
			}
		}
	}
}
