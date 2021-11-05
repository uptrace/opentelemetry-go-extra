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

**Uptrace** exporter:

```shell
UPTRACE_DSN="https://<token>@uptrace.dev/<project_id>" go run .
```

## Links

- [Find instrumentations](https://opentelemetry.uptrace.dev/instrumentations/?lang=go)
- [OpenTelemetry Tracing API](https://opentelemetry.uptrace.dev/guide/go-tracing.html)
