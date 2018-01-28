package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/staticmukesh/aerospike_exporter/collector"
)

var (
	asAddr        = flag.String("aerospike.addr", GetEnv("AEROSPIKE_ADDR", "localhost:3000"), "Address of aerospike node")
	asAlias       = flag.String("aerospike.alias", GetEnv("AEROSPIKE_ALIAS", ""), "Alias of aerospike node")
	listenAddr    = flag.String("web.listen-address", "0.0.0.0:9145", "Address to listen on for web interface and telemetry.")
	metricPath    = flag.String("web.telemetry-path", "/metrics", "Path under which to expose metrics.")
	asMetricsOnly = flag.Bool("as-only-metrics", false, "Whether to avoid metrics other than aerospike.")

	version = "<<<auto-filled by ldflags>>>"
	build   = "<<<auto-filled by ldflags>>>"
	commit  = "<<<auto-filled by ldflags>>>"

	logger *log.Logger
)

func init() {
	logger = log.New(os.Stdout, "[aerospike_exporter] ", log.LstdFlags)
}

func main() {
	flag.Parse()

	logger.Println("Initializing Aerospike Exporter", version)

	options := collector.Options{
		Addr:  *asAddr,
		Alias: *asAlias,
	}

	collector, err := collector.NewAerospike(options)
	if err != nil {
		logger.Panic(err)
	}

	logger.Println("Connecting to Aerospike Node:", *asAddr)
	logger.Println("Using node alias: ", *asAlias)

	// GET /metrics
	if *asMetricsOnly {
		registry := prometheus.NewRegistry()
		registry.Register(collector)
		handler := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
		http.Handle(*metricPath, handler)
	} else {
		prometheus.MustRegister(collector)
		http.Handle(*metricPath, prometheus.Handler())
	}

	// GET /
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
			<html>
				<head>
					<title>Aerospike Exporter ` + version + `</title>
				</head>
				<body>
					<h1>Aerospike Exporter ` + version + `</h1>
					<p><a href='` + *metricPath + `'>Metrics</a></p>
				</body>
			</html>
		`))
	})

	logger.Printf("Providing metrics at %s%s", *listenAddr, *metricPath)
	logger.Fatal(http.ListenAndServe(*listenAddr, nil))
}
