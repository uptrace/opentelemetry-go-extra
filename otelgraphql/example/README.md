# Example for graphql-go OpenTelemetry instrumentation

This example consists of a server that runs a GraphQL handler and a client that makes requests to
the server. First, you should run the server and then run the client.

## Server

You can run the server with different exporters by providing environment variables.

**Stdout** exporter (default):

```shell
go run server.go
```

**Jaeger** exporter:

```
OTEL_EXPORTER_JAEGER_ENDPOINT=http://localhost:14268/api/traces go run server.go
```

**Uptrace** exporter:

```shell
UPTRACE_DSN="https://<token>@api.uptrace.dev/<project_id>" go run server.go
```

## Client

To run the client:

```shell
./make_calls.sh
```

## Documentation

See [otelgraphql](../) for documentation.
