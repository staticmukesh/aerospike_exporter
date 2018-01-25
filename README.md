# Aerospike Metrics Exporter
[![Circle CI](https://circleci.com/gh/staticmukesh/aerospike_exporter.svg?style=shield)](https://circleci.com/gh/staticmukesh/aerospike_exporter)

A Prometheus exporter for collecting Aerospike metrics.

## Installation

Download latest executable from [releases](https://github.com/staticmukesh/aerospike_exporter/releases), and extract and run it.

```
$ ./aerospike_exporter <flags>
```

Add a block to the `scrape_configs` of your prometheus.yml config file:

```
scrape_configs:

...
  - job_name: aerospike_exporter
    static_configs:
      - targets: ['localhost:9145']
...
```
and adjust the host name accordingly.

### Flags

Name               | Description
-------------------|------------
aerospike.addr     | Address of aerospike node, defaults to `localhost:3000`.
aerospike.alias    | Alias for aerospike node address.
web.listen-address | Address to listen on for web interface and telemetry, defaults to `0.0.0.0:9145`.
web.telemetry-path | Path under which to expose metrics, defaults to `metrics`.

These settings take precedence over any configurations provided by [environment variables](#environment-variables).

### Environment Variables

Name               | Description
-------------------|------------
AEROSPIKE_ADDR     | Address of Aerospike node
AEROSPIKE_ALIAS    | Alias name of Aerospike node

### What's exported ?

The exporter collects metrics related to following items from `Aerospike`'s [Info](https://www.aerospike.com/docs/reference/info/index.html) command. 

- [x] Basic Info
- [x] Statistics
- [ ] Latency
- [ ] Namespace

### Local building and running

If you want to build package yourself, or want to contribute. You can follow the guide to setup the project.

```
$ go get github.com/staticmukesh/aerospike_exporter
$ go build
$ ./aerospike_exporter <flags>
```

To rebuild the package on source change, 
```
$ go get github.com/Unknwon/bra
$ bra run
```

### Contributing

Open an issue or PR if you have more suggestions or ideas about what to add or improve.

### Thanks

This project has been inspired from [redis_exporter](https://github.com/oliver006/redis_exporter) by [oliver006](https://github.com/oliver006).

