module github.com/uptrace/opentelemetry-go-extra/otelgraphqlgo/example

go 1.15

replace github.com/uptrace/opentelemetry-go-extra/otelgraphql-go => ./..

require (
	github.com/google/uuid v1.3.0
	github.com/graph-gophers/graphql-go v1.2.0
	github.com/uptrace/opentelemetry-go-extra/otelgraphql-go v0.0.0-00010101000000-000000000000
	go.opentelemetry.io/otel v1.1.0
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.1.0
	go.opentelemetry.io/otel/sdk v1.1.0
)
