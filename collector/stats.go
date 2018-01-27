package collector

import (
	"strconv"
	"strings"
	"time"

	aerospike "github.com/aerospike/aerospike-client-go"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/staticmukesh/aerospike_exporter/metrics"
)

type statsCollector struct {
	c collector
}

func (s *statsCollector) Collect(ch chan<- prometheus.Metric) {
	s.c.Collect(ch)
}

func (s *statsCollector) Describe(ch chan<- *prometheus.Desc) {
	s.c.Describe(ch)
}

func (s *statsCollector) extract() (Scraps, error) {
	conn, err := aerospike.NewConnection(s.c.addr, 10*time.Second)
	if err != nil {
		logger.Println("Aerospike error:", err)
		return Scraps{}, err
	}
	defer conn.Close()

	scraps, err := aerospike.RequestInfo(conn, s.c.field)
	if err != nil {
		logger.Println("Aerospike error:", err)
		return Scraps{}, nil
	}

	if len(s.c.field) != 0 {
		if unparsed, ok := scraps[s.c.field]; ok {
			scraps = s.c.parse(unparsed)
			return scraps, nil
		}
	}

	return scraps, err
}

func (s *statsCollector) process(scraps Scraps) {
	s.c.Lock()
	defer s.c.Unlock()

	for k, v := range scraps {
		if info, ok := s.c.meta[k]; ok {
			switch info.Type {
			case metrics.String:
				s.c.metrics[k].WithLabelValues(s.c.addr, s.c.alias, v).Set(1)
			case metrics.Float:
				val, err := strconv.ParseFloat(v, 64)
				if err != nil {
					logger.Println(err)
					continue
				}
				s.c.metrics[k].WithLabelValues(s.c.addr, s.c.alias).Set(val)
			case metrics.Bool:
				val, err := strconv.ParseBool(v)
				if err != nil {
					logger.Println(err)
					continue
				}
				if val {
					s.c.metrics[k].WithLabelValues(s.c.addr, s.c.alias).Set(1)
				} else {
					s.c.metrics[k].WithLabelValues(s.c.addr, s.c.alias).Set(0)
				}
			}
		}
	}
}

func (s *statsCollector) parse(unparsed string) Scraps {
	scraps := Scraps{}

	pairs := strings.Split(unparsed, ";")
	for _, pair := range pairs {
		data := strings.Split(pair, "=")
		scraps[data[0]] = data[1]
	}

	return scraps
}
