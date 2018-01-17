package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/staticmukesh/aerospike_exporter/exporter"
)

var (
	asAddr     = flag.String("aerospike.addr", GetEnv("AEROSPIKE_ADDR", "localhost:3000"), "Address of aerospike node")
	listenAddr = flag.String("web.listen-address", GetEnv("LISTEN_ADDR", ":9090"), "Address to listen on for web interface and telemetry.")
	metricPath = flag.String("web.telemetry-path", "/metrics", "Path under which to expose metrics.")

	version = "<<< filled in by build >>>"
	build   = "<<< filled in by build >>>"
	commit  = "<<< filled in by build >>>"
)

func main() {
	flag.Parse()

	exporter := exporter.New(exporter.Options{
		Addr: "localhost:3000",
	})

	prometheus.Register(exporter)

	// GET /
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
			<html>
				<head>
					<title>Aerospike Exporter v` + version + `</title>
				</head>
				<body>
					<h1>Aerospike Exporter v` + version + `</h1>
					<p><a href='` + *metricPath + `'>Metrics</a></p>
				</body>
			</html>
		`))
	})

	// GET /metrics
	http.Handle(*metricPath, prometheus.Handler())

	log.Printf("Providing metrics at %s%s", *listenAddr, *metricPath)
	log.Fatal(http.ListenAndServe(*listenAddr, nil))
}
