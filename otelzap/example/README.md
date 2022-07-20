# Example for otelzap instrumentation

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

See [otelzap](../) for documentation.

## Links

- [OpenTelemetry Go instrumentations](https://uptrace.dev/opentelemetry/instrumentations/?lang=go)
- [OpenTelemetry Tracing API](https://uptrace.dev/opentelemetry/go-tracing.html)
- [Distributed tracing tools](https://uptrace.dev/get/compare/distributed-tracing-tools.html)
