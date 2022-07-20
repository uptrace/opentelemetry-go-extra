# go-redis cluster example for OpenTelemetry

[![PkgGoDev](https://pkg.go.dev/badge/github.com/go-redis/redis/tree/master/extra/redisotel)](https://pkg.go.dev/github.com/go-redis/redis/tree/master/extra/redisotel)

To run this example you need a Redis cluster. You can start one with Docker:

```shell
docker-compose up -d
```

You can run this example with different exporters by providing environment variables.

**Stdout** exporter (default):

```shell
go run .
```

**Jaeger** exporter:

```shell
OTEL_EXPORTER_JAEGER_ENDPOINT=http://localhost:14268/api/traces go run .
```

[Uptrace](https://github.com/uptrace/uptrace/) exporter:

```shell
UPTRACE_DSN="https://<token>@uptrace.dev/<project_id>" go run .
```

## Links

- [OpenTelemetry Go instrumentations](https://uptrace.dev/opentelemetry/instrumentations/?lang=go)
- [OpenTelemetry Tracing API](https://uptrace.dev/opentelemetry/go-tracing.html)
- [Best distributed tracing tools](https://uptrace.dev/get/compare/distributed-tracing-tools.html)
