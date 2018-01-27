package metrics

var (
	// namespaceInfo stores aerospike namespace field to metric meta mapping
	namespaceInfo = map[string]Info{
		"ns_cluster_size": Info{
			Type:      Float,
			Name:      "ns_cluster_size",
			Help:      "Namespace cluster size",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"effective_replication_factor": Info{
			Type:      Float,
			Name:      "effective_replication_factor",
			Help:      "Effective replication factor",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"objects": Info{
			Type:      Float,
			Name:      "objects",
			Help:      "Number of objects",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"tombstones": Info{
			Type:      Float,
			Name:      "tombstones",
			Help:      "Number of tombstones",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"master_objects": Info{
			Type:      Float,
			Name:      "master_objects",
			Help:      "Number of master objects",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"master_tombstones": Info{
			Type:      Float,
			Name:      "master_tombstones",
			Help:      "Number of master tombstones",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"prole_objects": Info{
			Type:      Float,
			Name:      "prole_objects",
			Help:      "Number of prole objects",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"prole_tombstones": Info{
			Type:      Float,
			Name:      "prole_tombstones",
			Help:      "Number of prole tombstones",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"non_replica_objects": Info{
			Type:      Float,
			Name:      "non_replica_objects",
			Help:      "Number of non replica objects",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"non_replica_tombstones": Info{
			Type:      Float,
			Name:      "non_replica_tombstones",
			Help:      "Number of non replica tombstones",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"stop_writes": Info{
			Type:      Float,
			Name:      "stop_writes",
			Help:      "Number of stop writes",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"hwm_breached": Info{
			Type:      Float,
			Name:      "hwm_breached",
			Help:      "Number of hwm breached",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"current_time": Info{
			Type:      Float,
			Name:      "current_time",
			Help:      "Current time",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"non_expirable_objects": Info{
			Type:      Float,
			Name:      "non_expirable_objects",
			Help:      "Number of non expirable objects",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"expired_objects": Info{
			Type:      Float,
			Name:      "expired_objects",
			Help:      "Number of expired objects",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"evicted_objects": Info{
			Type:      Float,
			Name:      "evicted_objects",
			Help:      "Number of evicted objects",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"evict_ttl": Info{
			Type:      Float,
			Name:      "evict_ttl",
			Help:      "Evict TTL",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"nsup_cycle_duration": Info{
			Type:      Float,
			Name:      "nsup_cycle_duration",
			Help:      "Namespace Up Cycle duration",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"nsup_cycle_sleep_pct": Info{
			Type:      Float,
			Name:      "nsup_cycle_sleep_pct",
			Help:      "Namespace Up Cycle sleep percent",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"truncate_lut": Info{
			Type:      Float,
			Name:      "truncate_lut",
			Help:      "Truncate lut",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"truncated_records": Info{
			Type:      Float,
			Name:      "truncated_records",
			Help:      "Number of truncated records",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"memory_used_bytes": Info{
			Type:      Float,
			Name:      "memory_used_bytes",
			Help:      "Memory used in bytes",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"memory_used_data_bytes": Info{
			Type:      Float,
			Name:      "memory_used_data_bytes",
			Help:      "Memory used for data in bytes",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"memory_used_index_bytes": Info{
			Type:      Float,
			Name:      "memory_used_index_bytes",
			Help:      "Memory used for index in bytes",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"memory_used_sindex_bytes": Info{
			Type:      Float,
			Name:      "memory_used_sindex_bytes",
			Help:      "Memory used for sindex in bytes",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"memory_free_pct": Info{
			Type:      Float,
			Name:      "memory_free_pct",
			Help:      "Memory free in percent",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"xmem_id": Info{
			Type:      Float,
			Name:      "xmem_id",
			Help:      "X-Memory id",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"available_bin_names": Info{
			Type:      Float,
			Name:      "available_bin_names",
			Help:      "Number available bin names",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"device_total_bytes": Info{
			Type:      Float,
			Name:      "device_total_bytes",
			Help:      "Device total size in bytes",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"device_used_bytes": Info{
			Type:      Float,
			Name:      "device_used_bytes",
			Help:      "Device size used in bytes",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"device_free_pct": Info{
			Type:      Float,
			Name:      "device_free_pct",
			Help:      "Device size free in pct",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"device_available_pct": Info{
			Type:      Float,
			Name:      "device_available_pct",
			Help:      "Device available in percent",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"cache_read_pct": Info{
			Type:      Float,
			Name:      "cache_read_pct",
			Help:      "Cache read in percent",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"migrate_tx_partitions_imbalance": Info{
			Type:      Float,
			Name:      "migrate_tx_partitions_imbalance",
			Help:      "Number of Migrate tx_partitions imbalance",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"migrate_tx_instances": Info{
			Type:      Float,
			Name:      "migrate_tx_instances",
			Help:      "Number of migrate tx instances",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"migrate_rx_instances": Info{
			Type:      Float,
			Name:      "migrate_rx_instances",
			Help:      "Number of migrate rx instances",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"migrate_tx_partitions_active": Info{
			Type:      Float,
			Name:      "migrate_tx_partitions_active",
			Help:      "Number of migrate active tx_partitions",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"migrate_rx_partitions_active": Info{
			Type:      Float,
			Name:      "migrate_rx_partitions_active",
			Help:      "Number of migrate rx_partitions",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"migrate_tx_partitions_initial": Info{
			Type:      Float,
			Name:      "migrate_tx_partitions_initial",
			Help:      "Number of intial migrate tx_partitions",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"migrate_tx_partitions_remaining": Info{
			Type:      Float,
			Name:      "migrate_tx_partitions_remaining",
			Help:      "Number of migrate tx_partitions remaining",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"migrate_rx_partitions_initial": Info{
			Type:      Float,
			Name:      "migrate_rx_partitions_initial",
			Help:      "Number of migrate_rx_partitions initial",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"migrate_rx_partitions_remaining": Info{
			Type:      Float,
			Name:      "migrate_rx_partitions_remaining",
			Help:      "Number of migrate rx_partitions remaining",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"migrate_records_skipped": Info{
			Type:      Float,
			Name:      "migrate_records_skipped",
			Help:      "Number of migrate_records_skipped",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"migrate_records_transmitted": Info{
			Type:      Float,
			Name:      "migrate_records_transmitted",
			Help:      "Number of migrate_records_transmitted",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"migrate_record_retransmits": Info{
			Type:      Float,
			Name:      "migrate_record_retransmits",
			Help:      "Number of migrate_record_retransmits",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"migrate_record_receives": Info{
			Type:      Float,
			Name:      "migrate_record_receives",
			Help:      "Number of migrate_record_receives",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"migrate_signals_active": Info{
			Type:      Float,
			Name:      "migrate_signals_active",
			Help:      "Number of migrate_signals_active",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"migrate_signals_remaining": Info{
			Type:      Float,
			Name:      "migrate_signals_remaining",
			Help:      "Number of migrate_signals_remaining",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"client_tsvc_error": Info{
			Type:      Float,
			Name:      "client_tsvc_error",
			Help:      "Number of client_tsvc_error",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"client_tsvc_timeout": Info{
			Type:      Float,
			Name:      "client_tsvc_timeout",
			Help:      "Number of client_tsvc_timeout",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"client_proxy_complete": Info{
			Type:      Float,
			Name:      "client_proxy_complete",
			Help:      "Number of client_proxy_complete",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"client_proxy_error": Info{
			Type:      Float,
			Name:      "client_proxy_error",
			Help:      "Number of client_proxy_error",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"client_proxy_timeout": Info{
			Type:      Float,
			Name:      "client_proxy_timeout",
			Help:      "Number of client_proxy_timeout",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"client_read_success": Info{
			Type:      Float,
			Name:      "client_read_success",
			Help:      "Number of client_read_success",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"client_read_error": Info{
			Type:      Float,
			Name:      "client_read_error",
			Help:      "Number of client_read_error",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"client_read_timeout": Info{
			Type:      Float,
			Name:      "client_read_timeout",
			Help:      "Number of client_read_timeout",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"client_read_not_found": Info{
			Type:      Float,
			Name:      "client_read_not_found",
			Help:      "Number of client_read_not_found",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"client_write_success": Info{
			Type:      Float,
			Name:      "client_write_success",
			Help:      "Number of client_write_success",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"client_write_error": Info{
			Type:      Float,
			Name:      "client_write_error",
			Help:      "Number of client_write_error",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"client_write_timeout": Info{
			Type:      Float,
			Name:      "client_write_timeout",
			Help:      "Number of client_write_timeout",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"xdr_write_success": Info{
			Type:      Float,
			Name:      "xdr_write_success",
			Help:      "Number of xdr_write_success",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"xdr_write_error": Info{
			Type:      Float,
			Name:      "xdr_write_error",
			Help:      "Number of xdr_write_error",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"xdr_write_timeout": Info{
			Type:      Float,
			Name:      "xdr_write_timeout",
			Help:      "Number of xdr_write_timeout",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"client_delete_success": Info{
			Type:      Float,
			Name:      "client_delete_success",
			Help:      "Number of client_delete_success",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"client_delete_error": Info{
			Type:      Float,
			Name:      "client_delete_error",
			Help:      "Number of client_delete_error",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"client_delete_timeout": Info{
			Type:      Float,
			Name:      "client_delete_timeout",
			Help:      "Number of client_delete_timeout",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"client_delete_not_found": Info{
			Type:      Float,
			Name:      "client_delete_not_found",
			Help:      "Number of client_delete_not_found",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"client_udf_complete": Info{
			Type:      Float,
			Name:      "client_udf_complete",
			Help:      "Number of client_udf_complete",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"client_udf_error": Info{
			Type:      Float,
			Name:      "client_udf_error",
			Help:      "Number of client_udf_error",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"client_udf_timeout": Info{
			Type:      Float,
			Name:      "client_udf_timeout",
			Help:      "client_udf_timeout",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"client_lang_read_success": Info{
			Type:      Float,
			Name:      "client_lang_read_success",
			Help:      "Number of client_lang_read_success",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"client_lang_write_success": Info{
			Type:      Float,
			Name:      "client_lang_write_success",
			Help:      "Number of client_lang_write_success",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"client_lang_delete_success": Info{
			Type:      Float,
			Name:      "client_lang_delete_success",
			Help:      "Number of client_lang_delete_success",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"client_lang_error": Info{
			Type:      Float,
			Name:      "client_lang_error",
			Help:      "Number of client_lang_error",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"batch_sub_tsvc_error": Info{
			Type:      Float,
			Name:      "batch_sub_tsvc_error",
			Help:      "Number of batch_sub_tsvc_error",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"batch_sub_tsvc_timeout": Info{
			Type:      Float,
			Name:      "batch_sub_tsvc_timeout",
			Help:      "Number of batch_sub_tsvc_timeout",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"batch_sub_proxy_complete": Info{
			Type:      Float,
			Name:      "batch_sub_proxy_complete",
			Help:      "Number of batch_sub_proxy_complete",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"batch_sub_proxy_error": Info{
			Type:      Float,
			Name:      "batch_sub_proxy_error",
			Help:      "Number of batch_sub_proxy_error",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"batch_sub_proxy_timeout": Info{
			Type:      Float,
			Name:      "batch_sub_proxy_timeout",
			Help:      "Number of batch_sub_proxy_timeout",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"batch_sub_read_success": Info{
			Type:      Float,
			Name:      "batch_sub_read_success",
			Help:      "Number of batch_sub_read_success",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"batch_sub_read_error": Info{
			Type:      Float,
			Name:      "batch_sub_read_error",
			Help:      "Number of batch_sub_read_error",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"batch_sub_read_timeout": Info{
			Type:      Float,
			Name:      "batch_sub_read_timeout",
			Help:      "Number of batch_sub_read_timeout",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"batch_sub_read_not_found": Info{
			Type:      Float,
			Name:      "batch_sub_read_not_found",
			Help:      "Number of batch_sub_read_not_found",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"udf_sub_tsvc_error": Info{
			Type:      Float,
			Name:      "udf_sub_tsvc_error",
			Help:      "Number of udf_sub_tsvc_error",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"udf_sub_tsvc_timeout": Info{
			Type:      Float,
			Name:      "udf_sub_tsvc_timeout",
			Help:      "Number of udf_sub_tsvc_timeout",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"udf_sub_udf_complete": Info{
			Type:      Float,
			Name:      "udf_sub_udf_complete",
			Help:      "Number of udf_sub_udf_complete",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"udf_sub_udf_error": Info{
			Type:      Float,
			Name:      "udf_sub_udf_error",
			Help:      "Number of udf_sub_udf_error",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"udf_sub_udf_timeout": Info{
			Type:      Float,
			Name:      "udf_sub_udf_timeout",
			Help:      "Number of udf_sub_udf_timeout",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"udf_sub_lang_read_success": Info{
			Type:      Float,
			Name:      "udf_sub_lang_read_success",
			Help:      "Number of udf_sub_lang_read_success",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"udf_sub_lang_write_success": Info{
			Type:      Float,
			Name:      "udf_sub_lang_write_success",
			Help:      "Number of udf_sub_lang_write_success",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"udf_sub_lang_delete_success": Info{
			Type:      Float,
			Name:      "udf_sub_lang_delete_success",
			Help:      "Number of udf_sub_lang_delete_success",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"udf_sub_lang_error": Info{
			Type:      Float,
			Name:      "udf_sub_lang_error",
			Help:      "Number of udf_sub_lang_error",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"retransmit_client_read_dup_res": Info{
			Type:      Float,
			Name:      "retransmit_client_read_dup_res",
			Help:      "Number of retransmit_client_read_dup_res",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"retransmit_client_write_dup_res": Info{
			Type:      Float,
			Name:      "retransmit_client_write_dup_res",
			Help:      "Number of retransmit_client_write_dup_res",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"retransmit_client_write_repl_write": Info{
			Type:      Float,
			Name:      "retransmit_client_write_repl_write",
			Help:      "Number of retransmit_client_write_repl_write",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"retransmit_client_delete_dup_res": Info{
			Type:      Float,
			Name:      "retransmit_client_delete_dup_res",
			Help:      "Number of retransmit_client_delete_dup_res",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"retransmit_client_delete_repl_write": Info{
			Type:      Float,
			Name:      "retransmit_client_delete_repl_write",
			Help:      "Number of retransmit_client_delete_repl_write",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"retransmit_client_udf_dup_res": Info{
			Type:      Float,
			Name:      "retransmit_client_udf_dup_res",
			Help:      "Number of retransmit_client_udf_dup_res",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"retransmit_client_udf_repl_write": Info{
			Type:      Float,
			Name:      "retransmit_client_udf_repl_write",
			Help:      "Number of retransmit_client_udf_repl_write",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"retransmit_batch_sub_dup_res": Info{
			Type:      Float,
			Name:      "retransmit_batch_sub_dup_res",
			Help:      "Number of retransmit_batch_sub_dup_res",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"retransmit_udf_sub_dup_res": Info{
			Type:      Float,
			Name:      "retransmit_udf_sub_dup_res",
			Help:      "Number of retransmit_udf_sub_dup_res",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"retransmit_udf_sub_repl_write": Info{
			Type:      Float,
			Name:      "retransmit_udf_sub_repl_write",
			Help:      "Number of retransmit_udf_sub_repl_write",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"retransmit_nsup_repl_write": Info{
			Type:      Float,
			Name:      "retransmit_nsup_repl_write",
			Help:      "Number of retransmit_nsup_repl_write",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"scan_basic_complete": Info{
			Type:      Float,
			Name:      "scan_basic_complete",
			Help:      "Number of scan_basic_complete",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"scan_basic_error": Info{
			Type:      Float,
			Name:      "scan_basic_error",
			Help:      "Number of scan_basic_error",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"scan_basic_abort": Info{
			Type:      Float,
			Name:      "scan_basic_abort",
			Help:      "Number of scan_basic_abort",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"scan_aggr_complete": Info{
			Type:      Float,
			Name:      "scan_aggr_complete",
			Help:      "Number of scan_aggr_complete",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"scan_aggr_error": Info{
			Type:      Float,
			Name:      "scan_aggr_error",
			Help:      "Number of scan_aggr_error",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"scan_aggr_abort": Info{
			Type:      Float,
			Name:      "scan_aggr_abort",
			Help:      "Number of scan_aggr_abort",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"scan_udf_bg_complete": Info{
			Type:      Float,
			Name:      "scan_udf_bg_complete",
			Help:      "Number of scan_udf_bg_complete",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"scan_udf_bg_error": Info{
			Type:      Float,
			Name:      "scan_udf_bg_error",
			Help:      "Number of scan_udf_bg_error",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"scan_udf_bg_abort": Info{
			Type:      Float,
			Name:      "scan_udf_bg_abort",
			Help:      "Number of scan_udf_bg_abort",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"query_reqs": Info{
			Type:      Float,
			Name:      "query_reqs",
			Help:      "Number of query_reqs",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"query_fail": Info{
			Type:      Float,
			Name:      "query_fail",
			Help:      "Number of query_fail",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"query_short_queue_full": Info{
			Type:      Float,
			Name:      "query_short_queue_full",
			Help:      "Number of query_short_queue_full",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"query_long_queue_full": Info{
			Type:      Float,
			Name:      "query_long_queue_full",
			Help:      "Number of query_long_queue_full",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"query_short_reqs": Info{
			Type:      Float,
			Name:      "query_short_reqs",
			Help:      "Number of query_short_reqs",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"query_long_reqs": Info{
			Type:      Float,
			Name:      "query_long_reqs",
			Help:      "Number of query_long_reqs",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"query_agg": Info{
			Type:      Float,
			Name:      "query_agg",
			Help:      "Number of query_agg",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"query_agg_success": Info{
			Type:      Float,
			Name:      "query_agg_success",
			Help:      "Number of query_agg_success",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"query_agg_error": Info{
			Type:      Float,
			Name:      "query_agg_error",
			Help:      "Number of query_agg_error",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"query_agg_abort": Info{
			Type:      Float,
			Name:      "query_agg_abort",
			Help:      "Number of query_agg_abort",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"query_agg_avg_rec_count": Info{
			Type:      Float,
			Name:      "query_agg_avg_rec_count",
			Help:      "Number of query_agg_avg_rec_count",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"query_lookups": Info{
			Type:      Float,
			Name:      "query_lookups",
			Help:      "Number of query_lookups",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"query_lookup_success": Info{
			Type:      Float,
			Name:      "query_lookup_success",
			Help:      "Number of query_lookup_success",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"query_lookup_error": Info{
			Type:      Float,
			Name:      "query_lookup_error",
			Help:      "Number of query_lookup_error",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"query_lookup_abort": Info{
			Type:      Float,
			Name:      "query_lookup_abort",
			Help:      "Number of query_lookup_abort",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"query_lookup_avg_rec_count": Info{
			Type:      Float,
			Name:      "query_lookup_avg_rec_count",
			Help:      "Number of query_lookup_avg_rec_count",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"query_udf_bg_success": Info{
			Type:      Float,
			Name:      "query_udf_bg_success",
			Help:      "Number of query_udf_bg_success",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"query_udf_bg_failure": Info{
			Type:      Float,
			Name:      "query_udf_bg_failure",
			Help:      "Number of query_udf_bg_failure",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"geo_region_query_reqs": Info{
			Type:      Float,
			Name:      "geo_region_query_reqs",
			Help:      "Number of geo_region_query_reqs",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"geo_region_query_cells": Info{
			Type:      Float,
			Name:      "geo_region_query_cells",
			Help:      "Number of geo_region_query_cells",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"geo_region_query_points": Info{
			Type:      Float,
			Name:      "geo_region_query_points",
			Help:      "Number of geo_region_query_points",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"geo_region_query_falsepos": Info{
			Type:      Float,
			Name:      "geo_region_query_falsepos",
			Help:      "Number of geo_region_query_falsepos",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"fail_xdr_forbidden": Info{
			Type:      Float,
			Name:      "fail_xdr_forbidden",
			Help:      "Number of fail_xdr_forbidden",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"fail_key_busy": Info{
			Type:      Float,
			Name:      "fail_key_busy",
			Help:      "Number of fail_key_busy",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"fail_generation": Info{
			Type:      Float,
			Name:      "fail_generation",
			Help:      "Number of fail_generation",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"fail_record_too_big": Info{
			Type:      Float,
			Name:      "fail_record_too_big",
			Help:      "Number of fail_record_too_big",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"deleted_last_bin": Info{
			Type:      Float,
			Name:      "deleted_last_bin",
			Help:      "Number of deleted_last_bin",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"replication-factor": Info{
			Type:      Float,
			Name:      "replication-factor",
			Help:      "Number of replication-factor",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"memory-size": Info{
			Type:      Float,
			Name:      "memory-size",
			Help:      "memory-size",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"default-ttl": Info{
			Type:      Float,
			Name:      "default-ttl",
			Help:      "default-ttl",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"enable-xdr": Info{
			Type:      Bool,
			Name:      "enable-xdr",
			Help:      "Number of enable-xdr",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"sets-enable-xdr": Info{
			Type:      Bool,
			Name:      "sets-enable-xdr",
			Help:      "sets-enable-xdr",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"ns-forward-xdr-writes": Info{
			Type:      Bool,
			Name:      "ns-forward-xdr-writes",
			Help:      "ns-forward-xdr-writes",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"allow-nonxdr-writes": Info{
			Type:      Bool,
			Name:      "allow-nonxdr-writes",
			Help:      "allow-nonxdr-writes",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"allow-xdr-writes": Info{
			Type:      Bool,
			Name:      "allow-xdr-writes",
			Help:      "allow-xdr-writes",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"cold-start-evict-ttl": Info{
			Type:      Float,
			Name:      "cold-start-evict-ttl",
			Help:      "cold-start-evict-ttl",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"conflict-resolution-policy": Info{
			Type:      String,
			Name:      "conflict-resolution-policy",
			Help:      "conflict-resolution-policy",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"data-in-index": Info{
			Type:      Bool,
			Name:      "data-in-index",
			Help:      "data-in-index",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"disable-write-dup-res": Info{
			Type:      Bool,
			Name:      "disable-write-dup-res",
			Help:      "disable-write-dup-res",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"disallow-null-setname": Info{
			Type:      Bool,
			Name:      "disallow-null-setname",
			Help:      "disallow-null-setname",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"enable-benchmarks-batch-sub": Info{
			Type:      Bool,
			Name:      "enable-benchmarks-batch-sub",
			Help:      "enable-benchmarks-batch-sub",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"enable-benchmarks-read": Info{
			Type:      Bool,
			Name:      "enable-benchmarks-read",
			Help:      "enable-benchmarks-read",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"enable-benchmarks-udf": Info{
			Type:      Bool,
			Name:      "enable-benchmarks-udf",
			Help:      "enable-benchmarks-udf",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"enable-benchmarks-udf-sub": Info{
			Type:      Bool,
			Name:      "enable-benchmarks-udf-sub",
			Help:      "enable-benchmarks-udf-sub",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"enable-benchmarks-write": Info{
			Type:      Bool,
			Name:      "enable-benchmarks-write",
			Help:      "enable-benchmarks-write",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"enable-hist-proxy": Info{
			Type:      Bool,
			Name:      "enable-hist-proxy",
			Help:      "enable-hist-proxy",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"evict-hist-buckets": Info{
			Type:      Float,
			Name:      "evict-hist-buckets",
			Help:      "evict-hist-buckets",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"evict-tenths-pct": Info{
			Type:      Float,
			Name:      "evict-tenths-pct",
			Help:      "evict-tenths-pct",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"high-water-disk-pct": Info{
			Type:      Float,
			Name:      "high-water-disk-pct",
			Help:      "high-water-disk-pct",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"high-water-memory-pct": Info{
			Type:      Float,
			Name:      "high-water-memory-pct",
			Help:      "high-water-memory-pct",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"max-ttl": Info{
			Type:      Float,
			Name:      "max-ttl",
			Help:      "max-ttl",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"migrate-order": Info{
			Type:      Float,
			Name:      "migrate-order",
			Help:      "migrate-order",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"migrate-retransmit-ms": Info{
			Type:      Float,
			Name:      "migrate-retransmit-ms",
			Help:      "migrate-retransmit-ms",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"migrate-sleep": Info{
			Type:      Float,
			Name:      "migrate-sleep",
			Help:      "migrate-sleep",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"obj-size-hist-max": Info{
			Type:      Float,
			Name:      "obj-size-hist-max",
			Help:      "obj-size-hist-max",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"partition-tree-locks": Info{
			Type:      Float,
			Name:      "partition-tree-locks",
			Help:      "partition-tree-locks",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"partition-tree-sprigs": Info{
			Type:      Float,
			Name:      "partition-tree-sprigs",
			Help:      "partition-tree-sprigs",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"rack-id": Info{
			Type:      Float,
			Name:      "rack-id",
			Help:      "rack-id",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"read-consistency-level-override": Info{
			Type:      String,
			Name:      "read-consistency-level-override",
			Help:      "read-consistency-level-override",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"single-bin": Info{
			Type:      Float,
			Name:      "single-bin",
			Help:      "single-bin",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"stop-writes-pct": Info{
			Type:      Float,
			Name:      "stop-writes-pct",
			Help:      "stop-writes-pct",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"tomb-raider-eligible-age": Info{
			Type:      Float,
			Name:      "tomb-raider-eligible-age",
			Help:      "tomb-raider-eligible-age",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"tomb-raider-period": Info{
			Type:      Float,
			Name:      "tomb-raider-period",
			Help:      "tomb-raider-period",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"write-commit-level-override": Info{
			Type:      Float,
			Name:      "write-commit-level-override",
			Help:      "write-commit-level-override",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"storage-engine": Info{
			Type:      String,
			Name:      "storage-engine",
			Help:      "storage-engine",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"storage-engine.file": Info{
			Type:      String,
			Name:      "storage-engine.file",
			Help:      "storage-engine.file",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"storage-engine.filesize": Info{
			Type:      Float,
			Name:      "storage-engine.filesize",
			Help:      "storage-engine.filesize",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"storage-engine.write-block-size": Info{
			Type:      Float,
			Name:      "storage-engine.write-block-size",
			Help:      "storage-engine.write-block-size",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"storage-engine.data-in-memory": Info{
			Type:      Bool,
			Name:      "storage-engine.data-in-memory",
			Help:      "storage-engine.data-in-memory",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"storage-engine.cold-start-empty": Info{
			Type:      Bool,
			Name:      "storage-engine.cold-start-empty",
			Help:      "storage-engine.cold-start-empty",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"storage-engine.defrag-lwm-pct": Info{
			Type:      Float,
			Name:      "storage-engine.defrag-lwm-pct",
			Help:      "storage-engine.defrag-lwm-pct",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"storage-engine.defrag-queue-min": Info{
			Type:      Float,
			Name:      "storage-engine.defrag-queue-min",
			Help:      "storage-engine.defrag-queue-min",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"storage-engine.defrag-sleep": Info{
			Type:      Float,
			Name:      "storage-engine.defrag-sleep",
			Help:      "storage-engine.defrag-sleep",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"storage-engine.defrag-startup-minimum": Info{
			Type:      Float,
			Name:      "storage-engine.defrag-startup-minimum",
			Help:      "storage-engine.defrag-startup-minimum",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"storage-engine.disable-odirect": Info{
			Type:      Bool,
			Name:      "storage-engine.disable-odirect",
			Help:      "storage-engine.disable-odirect",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"storage-engine.enable-benchmarks-storage": Info{
			Type:      Bool,
			Name:      "storage-engine.enable-benchmarks-storage",
			Help:      "storage-engine.enable-benchmarks-storage",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"storage-engine.enable-osync": Info{
			Type:      Bool,
			Name:      "storage-engine.enable-osync",
			Help:      "storage-engine.enable-osync",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"storage-engine.flush-max-ms": Info{
			Type:      Float,
			Name:      "storage-engine.flush-max-ms",
			Help:      "storage-engine.flush-max-ms",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"storage-engine.fsync-max-sec": Info{
			Type:      Float,
			Name:      "storage-engine.fsync-max-sec",
			Help:      "storage-engine.fsync-max-sec",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"storage-engine.max-write-cache": Info{
			Type:      Float,
			Name:      "storage-engine.max-write-cache",
			Help:      "storage-engine.max-write-cache",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"storage-engine.min-avail-pct": Info{
			Type:      Float,
			Name:      "storage-engine.min-avail-pct",
			Help:      "storage-engine.min-avail-pct",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"storage-engine.post-write-queue": Info{
			Type:      Float,
			Name:      "storage-engine.post-write-queue",
			Help:      "storage-engine.post-write-queue",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"storage-engine.tomb-raider-sleep": Info{
			Type:      Float,
			Name:      "storage-engine.tomb-raider-sleep",
			Help:      "storage-engine.tomb-raider-sleep",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"storage-engine.write-threads": Info{
			Type:      Float,
			Name:      "storage-engine.write-threads",
			Help:      "storage-engine.write-threads",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"sindex.num-partitions": Info{
			Type:      Float,
			Name:      "sindex.num-partitions",
			Help:      "sindex.num-partitions",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"geo2dsphere-within.strict": Info{
			Type:      Bool,
			Name:      "geo2dsphere-within.strict",
			Help:      "geo2dsphere-within.strict",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"geo2dsphere-within.min-level": Info{
			Type:      Float,
			Name:      "geo2dsphere-within.min-level",
			Help:      "geo2dsphere-within.min-level",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"geo2dsphere-within.max-level": Info{
			Type:      Float,
			Name:      "geo2dsphere-within.max-level",
			Help:      "geo2dsphere-within.max-level",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"geo2dsphere-within.max-cells": Info{
			Type:      Float,
			Name:      "geo2dsphere-within.max-cells",
			Help:      "geo2dsphere-within.max-cells",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"geo2dsphere-within.level-mod": Info{
			Type:      Float,
			Name:      "geo2dsphere-within.level-mod",
			Help:      "geo2dsphere-within.level-mod",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},

		"geo2dsphere-within.earth-radius-meters": Info{
			Type:      Float,
			Name:      "geo2dsphere-within.earth-radius-meters",
			Help:      "geo2dsphere-within.earth-radius-meters",
			Labels:    []string{"addr", "alias", "namespace"},
			Subsystem: "ns",
		},
	}
)
