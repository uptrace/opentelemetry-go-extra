# gRPC OpenTelemetry instrumentation example

[![PkgGoDev](https://pkg.go.dev/badge/go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc)](https://pkg.go.dev/go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc)

[![PkgGoDev](https://pkg.go.dev/badge/go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho)](https://pkg.go.dev/go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho)

You can run this example with different exporters by providing environment variables.

**Stdout** exporter (default):

```shell
go run .
```

Start the gRPC server:

**Jaeger** exporter:

```shell
OTEL_EXPORTER_JAEGER_ENDPOINT=http://localhost:14268/api/traces go run server/server.go
```

**Uptrace** exporter:

```shell
UPTRACE_DSN="https://<token>@api.uptrace.dev/<project_id>" go run server/server.go
```

Start the client:

**Jaeger** exporter:

```shell
OTEL_EXPORTER_JAEGER_ENDPOINT=http://localhost:14268/api/traces go run client/client.go
```

**Uptrace** exporter:

```shell
UPTRACE_DSN="https://<token>@api.uptrace.dev/<project_id>" go run client/client.go
```

## Links

- [Find instrumentations](https://opentelemetry.uptrace.dev/instrumentations/?lang=go)
- [OpenTelemetry Tracing API](https://opentelemetry.uptrace.dev/guide/go-tracing.html)

