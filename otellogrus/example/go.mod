module github.com/uptrace/opentelemetry-go-extra/otellogrus/example

go 1.17

replace github.com/uptrace/opentelemetry-go-extra/otellogrus => ./..

require (
	github.com/sirupsen/logrus v1.8.1
	github.com/uptrace/opentelemetry-go-extra/otellogrus v1.0.5
	go.opentelemetry.io/otel v1.0.1
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.0.1
	go.opentelemetry.io/otel/sdk v1.0.1
)

require (
	go.opentelemetry.io/otel/trace v1.0.1 // indirect
	golang.org/x/sys v0.0.0-20211013075003-97ac67df715c // indirect
)
