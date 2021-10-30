module github.com/uptrace/opentelemetry-go-extra/otelgraphql/example

go 1.15

replace github.com/uptrace/opentelemetry-go-extra/otelgraphql => ../

replace github.com/uptrace/opentelemetry-go-extra/otelplay => ../../otelplay

require (
	github.com/google/uuid v1.3.0
	github.com/graph-gophers/graphql-go v1.2.0
	github.com/uptrace/opentelemetry-go-extra/otelgraphql v0.0.0-20211030063627-9b916d325530
	github.com/uptrace/opentelemetry-go-extra/otelplay v0.0.0-00010101000000-000000000000
	golang.org/x/sys v0.0.0-20211029165221-6e7872819dc8 // indirect
)
