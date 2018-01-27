package metrics

var (
	// basicInfo stores aerospike basic field to metric meta mapping
	basicInfo = map[string]Info{
		"node": Info{
			Type:   String,
			Name:   "node_name",
			Help:   "Name of Aerospike node",
			Labels: []string{"addr", "alias", "id"},
		},
		"edition": Info{
			Type:   String,
			Name:   "edtion",
			Help:   "Edition of Aerospike node",
			Labels: []string{"addr", "alias", "edition"},
		},
		"build": Info{
			Type:   String,
			Name:   "build",
			Help:   "Build of Aerospike node",
			Labels: []string{"addr", "alias", "build"},
		},
		"partition-generation": Info{
			Type:   Float,
			Name:   "partition_generation",
			Help:   "Partition generation of Aerospike node",
			Labels: []string{"addr", "alias"},
		},
		"cluster-generation": Info{
			Type:   Float,
			Name:   "cluster_generation",
			Help:   "Cluster generation of Aerospike node",
			Labels: []string{"addr", "alias"},
		},
	}
)
