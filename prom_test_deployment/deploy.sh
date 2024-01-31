docker run \
    --net="host" \
    -p localhost:9090:9090 \
    -d \
    -v $(pwd)/prometheus.yml:/etc/prometheus/prometheus.yml \
    prom/prometheus