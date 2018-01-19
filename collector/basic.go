package collector

var (
	basicMetricInfo = map[string]MetricInfo{
		"node": MetricInfo{
			MetricType: String,
			Name:       "node_name",
			Help:       "Name of Aerospike node",
			Labels:     []string{"addr", "alias", "id"},
		},
		"edition": MetricInfo{
			MetricType: String,
			Name:       "edtion",
			Help:       "Edition of Aerospike node",
			Labels:     []string{"addr", "alias", "edition"},
		},
		"build": MetricInfo{
			MetricType: String,
			Name:       "build",
			Help:       "Build of Aerospike node",
			Labels:     []string{"addr", "alias", "build"},
		},
		"partition-generation": MetricInfo{
			MetricType: Float,
			Name:       "partition_generation",
			Help:       "Partition generation of Aerospike node",
			Labels:     []string{"addr", "alias"},
		},
		"cluster-generation": MetricInfo{
			MetricType: Float,
			Name:       "cluster_generation",
			Help:       "Cluster generation of Aerospike node",
			Labels:     []string{"addr", "alias"},
		},
	}
)
