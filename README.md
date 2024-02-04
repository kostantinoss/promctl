# promctl
A simple client written in Go and Cobra for querying Prometheus.
The goal of this project was to practice my Go skills and get familiar woth cobra.

Before you begin, make sure you have GOPTH/bin added to your PATH
Next compile and install
```
go build
go install
```
## Usage
To begin, configure the Prometheus server you want to interact with
```
promctl config --server=localhost30010
```

To get the available metrics to stdout:
```
promctl metrics
```
This will print the available metrics on stdout

```
...
node_vmstat_oom_kill{instance="localhost:9100", job="node"} => 0 @[1707059976.955]
node_vmstat_pgfault{instance="localhost:9100", job="node"} => 112427232 @[1707059976.955]
node_vmstat_pgmajfault{instance="localhost:9100", job="node"} => 69741 @[1707059976.955]
node_vmstat_pgpgin{instance="localhost:9100", job="node"} => 4102101 @[1707059976.955]
node_vmstat_pgpgout{instance="localhost:9100", job="node"} => 5997525 @[1707059976.955]
node_vmstat_pswpin{instance="localhost:9100", job="node"} => 100535 @[1707059976.955]
node_vmstat_pswpout{instance="localhost:9100", job="node"} => 337466 @[1707059976.955]
process_cpu_seconds_total{instance="localhost:9100", job="node"} => 10.08 @[1707059976.955]
process_max_fds{instance="localhost:9100", job="node"} => 1048576 @[1707059976.955]
process_open_fds{instance="localhost:9100", job="node"} => 8 @[1707059976.955]
process_resident_memory_bytes{instance="localhost:9100", job="node"} => 19939328 @[1707059976.955]
process_start_time_seconds{instance="localhost:9100", job="node"} => 1707055289.21 @[1707059976.955]
process_virtual_memory_bytes{instance="localhost:9100", job="node"} => 1271705600 @[1707059976.955]
process_virtual_memory_max_bytes{instance="localhost:9100", job="node"} => 18446744073709552000 @[1707059976.955]
promhttp_metric_handler_errors_total{cause="encoding", instance="localhost:9100", job="node"} => 0 @[1707059976.955]
promhttp_metric_handler_errors_total{cause="gathering", instance="localhost:9100", job="node"} => 0 @[1707059976.955]
promhttp_metric_handler_requests_in_flight{instance="localhost:9100", job="node"} => 1 @[1707059976.955]
promhttp_metric_handler_requests_total{code="200", instance="localhost:9100", job="node"} => 312 @[1707059976.955]
promhttp_metric_handler_requests_total{code="500", instance="localhost:9100", job="node"} => 0 @[1707059976.955]
promhttp_metric_handler_requests_total{code="503", instance="localhost:9100", job="node"} => 0 @[1707059976.955]
scrape_duration_seconds{instance="localhost:9100", job="node"} => 0.279202294 @[1707059976.955]
scrape_samples_post_metric_relabeling{instance="localhost:9100", job="node"} => 778 @[1707059976.955]
scrape_samples_scraped{instance="localhost:9100", job="node"} => 778 @[1707059976.955]
scrape_series_added{instance="localhost:9100", job="node"} => 0 @[1707059976.955]
up{instance="localhost:9100", job="node"} => 1 @[1707059976.955]
...
```

To make a simple query:
```
promctl query --query='process_cpu_seconds_total{instance="localhost:9100"}' --output=stdout
```
