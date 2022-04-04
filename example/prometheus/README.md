# Prometheus exporter example for OpenTelemetry

This is an example for
[Sending OpenTelemetry Metrics to Prometheus](https://blog.uptrace.dev/posts/prometheus-opentelemetry-metrics.html).
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

- [OpenTelemetry Metrics](https://opentelemetry.uptrace.dev/guide/metrics.html)
- [OpenTelemetry Metrics API for Go](https://opentelemetry.uptrace.dev/guide/go-metrics.html)
- [Best distributed tracing tools](https://get.uptrace.dev/compare/distributed-tracing-tools.html)
