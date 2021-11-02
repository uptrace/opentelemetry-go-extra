# Gomemcache example for OpenTelemetry

[![PkgGoDev](https://pkg.go.dev/badge/go.opentelemetry.io/contrib/instrumentation/github.com/bradfitz/gomemcache/memcache/otelmemcache)](https://pkg.go.dev/go.opentelemetry.io/contrib/instrumentation/github.com/bradfitz/gomemcache/memcache/otelmemcache)

To run this example you need a memcached server. You can start one with Docker:

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
UPTRACE_DSN="https://<token>@api.uptrace.dev/<project_id>" go run .
```

## Links

- [Find instrumentations](https://opentelemetry.uptrace.dev/instrumentations/?lang=go)
- [OpenTelemetry Tracing API](https://opentelemetry.uptrace.dev/guide/go-tracing.html)

