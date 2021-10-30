module github.com/uptrace/opentelemetry-go-extra/otelgraphql/example

go 1.15

replace github.com/uptrace/opentelemetry-go-extra/otelgraphql => ../

require (
	github.com/google/uuid v1.3.0
	github.com/graph-gophers/graphql-go v1.2.0
	github.com/uptrace/opentelemetry-go-extra/otelgraphql v0.0.0-00010101000000-000000000000
	go.opentelemetry.io/otel v1.1.0
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.1.0
	go.opentelemetry.io/otel/sdk v1.1.0
)
