package exporter

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	as "github.com/aerospike/aerospike-client-go"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	statsMetricInfo = map[string]MetricInfo{
		"cluster_size": MetricInfo{
			valType: Float,
			name:    "cluster_size",
			help:    "Size of Aerospike cluster",
			labels:  []string{"addr", "alias"},
		},
		"cluster_key": MetricInfo{
			valType: String,
			name:    "cluster_key",
			help:    "Key of Aerospike cluster",
			labels:  []string{"addr", "alias", "key"},
		},
		"cluster_integrity": MetricInfo{
			valType: Bool,
			name:    "cluster_integrity",
			help:    "Is Integrity of Aerospike cluster maintained ?",
			labels:  []string{"addr", "alias"},
		},
		"cluster_is_member": MetricInfo{
			valType: Bool,
			name:    "cluster_is_member",
			help:    "Is node member of Aerospike cluster",
			labels:  []string{"addr", "alias"},
		},
		"uptime": MetricInfo{
			valType: Float,
			name:    "uptime",
			help:    "Uptime of Aerospike node",
			labels:  []string{"addr", "alias"},
		},
		"system_free_mem_pct": MetricInfo{
			valType: Float,
			name:    "system_free_mem_pct",
			help:    "Percent of free memory on system",
			labels:  []string{"addr", "alias"},
		},
		"system_swapping": MetricInfo{
			valType: Bool,
			name:    "system_swapping",
			help:    "Is swapping allowed on system?",
			labels:  []string{"addr", "alias"},
		},
		"heap_allocated_kbytes": MetricInfo{
			valType: Float,
			name:    "heap_allocated_kbytes",
			help:    "Kilobytes of memory allocated to heap",
			labels:  []string{"addr", "alias"},
		},
		"heap_active_kbytes": MetricInfo{
			valType: Float,
			name:    "heap_active_kbytes",
			help:    "Kilobytes of memory active in heap",
			labels:  []string{"addr", "alias"},
		},
		"heap_mapped_kbytes": MetricInfo{
			valType: Float,
			name:    "heap_mapped_kbytes",
			help:    "Kilobytes of memeory mapped to heap",
			labels:  []string{"addr", "alias"},
		},
		"heap_efficiency_pct": MetricInfo{
			valType: Float,
			name:    "heap_efficiency_pct",
			help:    "Percent of Heap efficiency",
			labels:  []string{"addr", "alias"},
		},
		"heap_site_count": MetricInfo{
			valType: Float,
			name:    "heap_site_count",
			help:    "Number of heap sites",
			labels:  []string{"addr", "alias"},
		},
		"objects": MetricInfo{
			valType: Float,
			name:    "objects",
			help:    "Number of objects in Aerospike node",
			labels:  []string{"addr", "alias"},
		},
		"tombstones": MetricInfo{
			valType: Float,
			name:    "tombstones",
			help:    "Number of tombstones in Aerospike node",
			labels:  []string{"addr", "alias"},
		},
		"tsvc_queue": MetricInfo{
			valType: Float,
			name:    "tsvc_queue",
			help:    "Size of tsvc queue",
			labels:  []string{"addr", "alias"},
		},
		"info_queue": MetricInfo{
			valType: Float,
			name:    "info_queue",
			help:    "Size of info queue",
			labels:  []string{"addr", "alias"},
		},
		"delete_queue": MetricInfo{
			valType: Float,
			name:    "delete_queue",
			help:    "Size of delete queue",
			labels:  []string{"addr", "alias"},
		},
		"rw_in_progress": MetricInfo{
			valType: Float,
			name:    "rw_in_progress",
			help:    "Number of rw in progress",
			labels:  []string{"addr", "alias"},
		},
		"proxy_in_progress": MetricInfo{
			valType: Float,
			name:    "proxy_in_progress",
			help:    "Number of proxy in progress",
			labels:  []string{"addr", "alias"},
		},
		"tree_gc_queue": MetricInfo{
			valType: Float,
			name:    "tree_gc_queue",
			help:    "Size of tree_gc queue",
			labels:  []string{"addr", "alias"},
		},
		"client_connections": MetricInfo{
			valType: Float,
			name:    "cient_connections",
			help:    "Nunber of client connections",
			labels:  []string{"addr", "alias"},
		},
		"heartbeat_connections": MetricInfo{
			valType: Float,
			name:    "heartbeat_connections",
			help:    "Number of heartbeat connections",
			labels:  []string{"addr", "alias"},
		},
		"fabric_connections": MetricInfo{
			valType: Float,
			name:    "fabric_connections",
			help:    "Nunber of fabric connections",
			labels:  []string{"addr", "alias"},
		},
		"heartbeat_received_self": MetricInfo{
			valType: Float,
			name:    "heartbeat_received_self",
			help:    "Number of heartbeat received from itself",
			labels:  []string{"addr", "alias"},
		},
		"heartbeat_received_foreign": MetricInfo{
			valType: Float,
			name:    "heartbeat_received_foreign",
			help:    "Number of heartbeat received from foreign",
			labels:  []string{"addr", "alias"},
		},
		"reaped_fds": MetricInfo{
			valType: Float,
			name:    "reaped_fds",
			help:    "Number of reaped fds",
			labels:  []string{"addr", "alias"},
		},
		"info_complete": MetricInfo{
			valType: Float,
			name:    "info_complete",
			help:    "Number of completed info",
			labels:  []string{"addr", "alias"},
		},
		"proxy_retry": MetricInfo{
			valType: Float,
			name:    "proxy_retry",
			help:    "Number of retried proxy",
			labels:  []string{"addr", "alias"},
		},
		"demarshal_error": MetricInfo{
			valType: Float,
			name:    "demarshal_error",
			help:    "Number of demarshal error",
			labels:  []string{"addr", "alias"},
		},
		"early_tsvc_client_error": MetricInfo{
			valType: Float,
			name:    "early_tsvc_client_error",
			help:    "Number of early_tsvc_client error",
			labels:  []string{"addr", "alias"},
		},
		"early_tsvc_batch_sub_error": MetricInfo{
			valType: Float,
			name:    "early_tsvc_batch_sub_error",
			help:    "Number of early_tsvc_batch_sub error",
			labels:  []string{"addr", "alias"},
		},
		"early_tsvc_udf_sub_error": MetricInfo{
			valType: Float,
			name:    "early_tsvc_udf_sub_error",
			help:    "Number of early_tsvc_udf_sub error",
			labels:  []string{"addr", "alias"},
		},
		"batch_index_initiate": MetricInfo{
			valType: Float,
			name:    "batch_index_initiate",
			help:    "Count of batch_index_initiate",
			labels:  []string{"addr", "alias"},
		},
		"batch_index_complete": MetricInfo{
			valType: Float,
			name:    "batch_index_complete",
			help:    "NUmber of completed batch_index",
			labels:  []string{"addr", "alias"},
		},
		"batch_index_error": MetricInfo{
			valType: Float,
			name:    "batch_index_error",
			help:    "NUmber of batch_index error",
			labels:  []string{"addr", "alias"},
		},
		"batch_index_timeout": MetricInfo{
			valType: Float,
			name:    "batch_index_timeout",
			help:    "Timeout of batch_index",
			labels:  []string{"addr", "alias"},
		},
		"batch_index_unused_buffers": MetricInfo{
			valType: Float,
			name:    "batch_index_unused_buffers",
			help:    "NUmber of batch_index unused buffers",
			labels:  []string{"addr", "alias"},
		},
		"batch_index_huge_buffers": MetricInfo{
			valType: Float,
			name:    "batch_index_huge_buffers",
			help:    "NUmber of batch_index huge buffers",
			labels:  []string{"addr", "alias"},
		},
		"batch_index_created_buffers": MetricInfo{
			valType: Float,
			name:    "batch_index_created_buffers",
			help:    "NUmber of batch_index created buffers",
			labels:  []string{"addr", "alias"},
		},
		"batch_index_destroyed_buffers": MetricInfo{
			valType: Float,
			name:    "batch_index_destroyed_buffers",
			help:    "NUmber of batch_index destroyed buffers",
			labels:  []string{"addr", "alias"},
		},
		"batch_initiate": MetricInfo{
			valType: Float,
			name:    "batch_initiate",
			help:    "batch initiate",
			labels:  []string{"addr", "alias"},
		},
		"batch_queue": MetricInfo{
			valType: Float,
			name:    "batch_queue",
			help:    "Size of batch queue",
			labels:  []string{"addr", "alias"},
		},
		"batch_error": MetricInfo{
			valType: Float,
			name:    "batch_error",
			help:    "Number of batch error",
			labels:  []string{"addr", "alias"},
		},
		"batch_timeout": MetricInfo{
			valType: Float,
			name:    "batch_timeout",
			help:    "Timeout of batch",
			labels:  []string{"addr", "alias"},
		},
		"scans_active": MetricInfo{
			valType: Float,
			name:    "scans_active",
			help:    "Number of active scans",
			labels:  []string{"addr", "alias"},
		},
		"query_short_running": MetricInfo{
			valType: Float,
			name:    "query_short_running",
			help:    "Number of running short query",
			labels:  []string{"addr", "alias"},
		},
		"query_long_running": MetricInfo{
			valType: Float,
			name:    "query_long_running",
			help:    "Number of running long query",
			labels:  []string{"addr", "alias"},
		},
		"sindex_ucgarbage_found": MetricInfo{
			valType: Float,
			name:    "sindex_ucgarbage_found",
			help:    "Number of sindex_ucgarbage found",
			labels:  []string{"addr", "alias"},
		},
		"sindex_gc_locktimedout": MetricInfo{
			valType: Float,
			name:    "sindex_gc_locktimedout",
			help:    "Locktimedout of sindex_gc",
			labels:  []string{"addr", "alias"},
		},
		"sindex_gc_list_creation_time": MetricInfo{
			valType: Float,
			name:    "sindex_gc_list_creation_time",
			help:    "Creation time of sindex_gc list",
			labels:  []string{"addr", "alias"},
		},
		"sindex_gc_list_deletion_time": MetricInfo{
			valType: Float,
			name:    "sindex_gc_list_deletion_time",
			help:    "Deletion time of  sindex_gc list",
			labels:  []string{"addr", "alias"},
		},
		"sindex_gc_objects_validated": MetricInfo{
			valType: Float,
			name:    "sindex_gc_objects_validated",
			help:    "Number of validated sindex_gc objects",
			labels:  []string{"addr", "alias"},
		},
		"sindex_gc_garbage_found": MetricInfo{
			valType: Float,
			name:    "sindex_gc_garbage_found",
			help:    "Number of sindex_gc_garbage found",
			labels:  []string{"addr", "alias"},
		},
		"sindex_gc_garbage_cleaned": MetricInfo{
			valType: Float,
			name:    "sindex_gc_garbage_cleaned",
			help:    "Number of sindex_gc_garbage cleaned",
			labels:  []string{"addr", "alias"},
		},
		"paxos_principal": MetricInfo{
			valType: String,
			name:    "paxos_principal",
			help:    "Aerospike paxos principal",
			labels:  []string{"addr", "alias", "paxos_principal"},
		},
		"migrate_allowed": MetricInfo{
			valType: Bool,
			name:    "migrate_allowed",
			help:    "Is migrate allowed ?",
			labels:  []string{"addr", "alias"},
		},
		"migrate_partitions_remaining": MetricInfo{
			valType: Float,
			name:    "migrate_partitions_remaining",
			help:    "Number of migrate_partitions remaining",
			labels:  []string{"addr", "alias"},
		},
		"fabric_bulk_send_rate": MetricInfo{
			valType: Float,
			name:    "fabric_bulk_send_rate",
			help:    "Rate of sent fabric_bulk",
			labels:  []string{"addr", "alias"},
		},
		"fabric_bulk_recv_rate": MetricInfo{
			valType: Float,
			name:    "fabric_bulk_recv_rate",
			help:    "Rate of received fabric_bulk",
			labels:  []string{"addr", "alias"},
		},
		"fabric_ctrl_send_rate": MetricInfo{
			valType: Float,
			name:    "fabric_ctrl_send_rate",
			help:    "Rate of sent fabric_ctrl",
			labels:  []string{"addr", "alias"},
		},
		"fabric_ctrl_recv_rate": MetricInfo{
			valType: Float,
			name:    "fabric_ctrl_recv_rate",
			help:    "Rate of received fabric_ctrl",
			labels:  []string{"addr", "alias"},
		},
		"fabric_meta_send_rate": MetricInfo{
			valType: Float,
			name:    "fabric_meta_send_rate",
			help:    "Rate of sent fabric_meta",
			labels:  []string{"addr", "alias"},
		},
		"fabric_meta_recv_rate": MetricInfo{
			valType: Float,
			name:    "fabric_meta_recv_rate",
			help:    "Rate of received fabric_meta",
			labels:  []string{"addr", "alias"},
		},
		"fabric_rw_send_rate": MetricInfo{
			valType: Float,
			name:    "fabric_rw_send_rate",
			help:    "Rate of sent fabric_rw",
			labels:  []string{"addr", "alias"},
		},
		"fabric_rw_recv_rate": MetricInfo{
			valType: Float,
			name:    "fabric_rw_recv_rate",
			help:    "Rate of receivied fabric_rw",
			labels:  []string{"addr", "alias"},
		},
	}
)

// NewStatsExporter initializes StatsExporter
func NewStatsExporter(options Options) Exporter {
	return &StatsExporter{
		Addr:      options.Addr,
		Alias:     options.Alias,
		Namespace: options.Namespace,
		Metrics:   initStatsMetrics(options.Namespace),
	}
}

// StatsExporter provides
type StatsExporter struct {
	Addr      string
	Alias     string
	Namespace string
	Metrics   map[string]*prometheus.GaugeVec
}

// Process accepts scrapes and processes it
func (e *StatsExporter) Process(scraps <-chan Scrap) {
	for scrap := range scraps {
		for k, v := range scrap {
			if strings.Compare(k, "statistics") != 0 {
				continue
			}

			pairs := strings.Split(v, ";")
			for _, pair := range pairs {
				m := strings.Split(pair, "=")
				name := m[0]
				value := m[1]
				if info, ok := statsMetricInfo[name]; ok {
					switch info.valType {
					case String:
						e.Metrics[name].WithLabelValues(e.Addr, e.Alias, value).Set(1)
					case Float:
						val, err := strconv.ParseFloat(value, 64)
						if err != nil {
							fmt.Println(err)
							continue
						}
						e.Metrics[name].WithLabelValues(e.Addr, e.Alias).Set(val)
					case Bool:
						val, err := strconv.ParseBool(value)
						if err != nil {
							fmt.Println(err)
							continue
						}
						if val {
							e.Metrics[name].WithLabelValues(e.Addr, e.Alias).Set(1)
						} else {
							e.Metrics[name].WithLabelValues(e.Addr, e.Alias).Set(0)
						}
					}
				}
			}
		}
	}
}

// Extract accepts generates scrapes
func (e *StatsExporter) Extract(scraps chan<- Scrap) {
	defer close(scraps)

	conn, err := as.NewConnection(e.Addr, time.Minute)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	info, err := as.RequestInfo(conn, "statistics")
	if err != nil {
		fmt.Println(err)
		return
	}

	scraps <- info
}

// Collect collects prometheus's Collector interface
func (e *StatsExporter) Collect(ch chan<- prometheus.Metric) {
	scrapes := make(chan Scrap)

	go e.Extract(scrapes)
	e.Process(scrapes)

	for _, metric := range e.Metrics {
		metric.Collect(ch)
	}
}

// Describe collects prometheus's Collector interface
func (e *StatsExporter) Describe(ch chan<- *prometheus.Desc) {
	for _, metric := range e.Metrics {
		metric.Describe(ch)
	}
}

func initStatsMetrics(namespace string) map[string]*prometheus.GaugeVec {
	metrics := map[string]*prometheus.GaugeVec{}

	for k, v := range statsMetricInfo {
		metrics[k] = prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name:      v.name,
			Help:      v.help,
			Namespace: namespace,
		}, v.labels)
	}

	return metrics
}
