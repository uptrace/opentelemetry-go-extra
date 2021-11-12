[![PkgGoDev](https://pkg.go.dev/badge/github.com/uptrace/opentelemetry-go-extra/otelgraphql)](https://pkg.go.dev/github.com/uptrace/opentelemetry-go-extra/otelgraphql)

# OpenTelemetry Go instrumentation for the graphql-go GraphQL Server

This instrumentation records GraphQL operations using OpenTelemetry API.

## Installation

```shell
go get github.com/uptrace/opentelemetry-go-extra/otelgraphql
```

## Usage

graphql-go already provides a tracer interface along with its respective
[opentracing implementation](https://github.com/graph-gophers/graphql-go/tree/v1.1.0/trace) This
interface consists of the following methods, all of which get `context.Context` as a parameter:

`TraceValidation`: traces the schema validation step which precedes the actual operation.

`TraceQuery`: traces the actual operation, query or mutation, as a whole

`TraceField`: traces a field-specific operation; its span should typically be a sub-span of the
TraceQuery one;

otelgraphql provides an implementation of this interface, which is practically a port of the
opentracing one that comes pre-packaged with graphql-go.

Some other points:

1. graphql-go exposes a single HTTP Handler for all graphql operations. That makes it a natural fit
   for otelhttp and other router packages instrumentation, if propagating frontend baggage (e.g.
   headers) is required.
2. graphql-go resolver methods do get `context.Context` as a parameter, which allows for
   field-specific (sub-)span creation, if `TraceQuery` does not suffice.

```go
import (
  "context"
  "log"
  "net/http"

  "github.com/graph-gophers/graphql-go"
  "github.com/graph-gophers/graphql-go/relay"

  "github.com/uptrace/opentelemetry-go-extra/otelgraphql"
)

tracer := otelgraphql.NewTracer(otelgraphql.WithTracerProvider(traceProvider))

opts := []graphql.SchemaOpt{
  graphql.Tracer(tracer),
}
schema = graphql.MustParseSchema(schemaString, &RootResolver{}, opts...)

http.Handle("/graphql", &relay.Handler{Schema: schema})

log.Fatal(http.ListenAndServe(":8080", nil))
```

See [example](/example/) for details.
