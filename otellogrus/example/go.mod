module github.com/uptrace/opentelemetry-go-extra/otellogrus/example

go 1.16

replace github.com/uptrace/opentelemetry-go-extra/otellogrus => ./..

replace github.com/uptrace/opentelemetry-go-extra/otelplay => ../../otelplay

require (
	github.com/sirupsen/logrus v1.8.1
	github.com/uptrace/opentelemetry-go-extra/otellogrus v0.1.4
	github.com/uptrace/opentelemetry-go-extra/otelplay v0.1.4
	go.opentelemetry.io/otel v1.1.0
)
