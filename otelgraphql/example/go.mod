module github.com/uptrace/opentelemetry-go-extra/otelgraphql/example

go 1.15

replace github.com/uptrace/opentelemetry-go-extra/otelgraphql => ../

replace github.com/uptrace/opentelemetry-go-extra/otelplay => ../../otelplay

require (
	github.com/google/uuid v1.3.0
	github.com/graph-gophers/graphql-go v1.2.0
	github.com/uptrace/opentelemetry-go-extra/otelgraphql v0.1.5
	github.com/uptrace/opentelemetry-go-extra/otelplay v0.1.5
)
