package collector

import (
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/aerospike/aerospike-client-go"
	"github.com/staticmukesh/aerospike_exporter/metrics"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	logger *log.Logger
)

func init() {
	logger = log.New(os.Stdout, "[aerospike_exporter] ", log.LstdFlags)
}

// Options store Collector initializing parameters
type Options struct {
	Addr  string
	Alias string
}

// Scraps stores data exported from Aerospike
type Scraps map[string]string

// Collector abstracts prometheus collector
type Collector interface {
	prometheus.Collector
}

// NewCollector initializes collector
func NewCollector(options Options, field string) (Collector, error) {
	meta := metrics.GetMetricsInfo(field)
	metrics := metrics.GetMetrics(meta)

	switch field {
	case "":
		return &collector{
			addr:    options.Addr,
			alias:   options.Alias,
			field:   field,
			meta:    meta,
			metrics: metrics,
		}, nil
	case "statistics":
		return &statsCollector{
			c: collector{
				addr:    options.Addr,
				alias:   options.Alias,
				field:   field,
				meta:    meta,
				metrics: metrics,
			},
		}, nil
	}

	return &collector{}, errors.New("Collector doesn't exist")
}

type collector struct {
	addr    string
	alias   string
	field   string
	meta    map[string]metrics.Info
	metrics metrics.Metrics
	sync.RWMutex
}

func (c *collector) Collect(ch chan<- prometheus.Metric) {
	scraps, err := c.extract()
	if err != nil {
		return
	}

	c.process(scraps)

	for _, metric := range c.metrics {
		metric.Collect(ch)
	}
}

func (c *collector) Describe(ch chan<- *prometheus.Desc) {
	for _, metric := range c.metrics {
		metric.Describe(ch)
	}
}

func (c *collector) extract() (Scraps, error) {
	conn, err := aerospike.NewConnection(c.addr, 10*time.Second)
	if err != nil {
		logger.Println("Aerospike error:", err)
		return Scraps{}, err
	}
	defer conn.Close()

	scraps, err := aerospike.RequestInfo(conn, c.field)
	if err != nil {
		logger.Println("Aerospike error:", err)
		return Scraps{}, nil
	}

	if len(c.field) != 0 {
		if unparsed, ok := scraps[c.field]; ok {
			scraps = c.parse(unparsed)
			return scraps, nil
		}
	}

	return scraps, err
}

func (c *collector) process(scraps Scraps) {
	c.Lock()
	defer c.Unlock()

	for k, v := range scraps {
		if info, ok := c.meta[k]; ok {
			switch info.Type {
			case metrics.String:
				c.metrics[k].WithLabelValues(c.addr, c.alias, v).Set(1)
			case metrics.Float:
				val, err := strconv.ParseFloat(v, 64)
				if err != nil {
					logger.Println(err)
					continue
				}
				c.metrics[k].WithLabelValues(c.addr, c.alias).Set(val)
			case metrics.Bool:
				val, err := strconv.ParseBool(v)
				if err != nil {
					logger.Println(err)
					continue
				}
				if val {
					c.metrics[k].WithLabelValues(c.addr, c.alias).Set(1)
				} else {
					c.metrics[k].WithLabelValues(c.addr, c.alias).Set(0)
				}
			}
		}
	}
}

func (c *collector) parse(unparsed string) Scraps {
	scraps := Scraps{}

	pairs := strings.Split(unparsed, ";")
	for _, pair := range pairs {
		data := strings.Split(pair, "=")
		scraps[data[0]] = data[1]
	}

	return scraps
}
