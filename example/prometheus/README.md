# Prometheus exporter example for OpenTelemetry

This is an example for
[Sending OpenTelemetry Metrics to Prometheus](https://uptrace.dev/opentelemetry/prometheus-metrics.html).
To run this example:

```shell
docker-compose up -d
```

Dockers starts 2 services:

- https://localhost:8088/metrics - Prometheus target that exports OpenTelemetry metrics.
- http://localhost:9090/graph - Prometheus configured to scrape our target.

You can use Prometheus to query the counter:

```
test_my_counter
```

## See also

- [OpenTelemetry Metrics](https://uptrace.dev/opentelemetry/metrics.html)
- [OpenTelemetry Metrics API for Go](https://uptrace.dev/opentelemetry/go-metrics.html)
- [Best distributed tracing tools](https://uptrace.dev/get/compare/distributed-tracing-tools.html)
