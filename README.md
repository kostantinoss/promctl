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
process_cpu_seconds_total{instance="localhost:9100", job="node"} => 10.08 @[1707059976.955]
process_max_fds{instance="localhost:9100", job="node"} => 1048576 @[1707059976.955]
process_open_fds{instance="localhost:9100", job="node"} => 8 @[1707059976.955]
process_resident_memory_bytes{instance="localhost:9100", job="node"} => 19939328 @[1707059976.955]
up{instance="localhost:9100", job="node"} => 1 @[1707059976.955]
...
```

To make a simple query:
```
promctl query --query='process_cpu_seconds_total{instance="localhost:9100"}' --output=stdout
```
