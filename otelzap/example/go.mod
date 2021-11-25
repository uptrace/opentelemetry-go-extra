module github.com/uptrace/opentelemetry-go-extra/otelzap/example

go 1.16

replace github.com/uptrace/opentelemetry-go-extra/otelzap => ../

replace github.com/uptrace/opentelemetry-go-extra/otelutil => ../../otelutil

replace github.com/uptrace/opentelemetry-go-extra/otelplay => ../../otelplay

require (
	github.com/uptrace/opentelemetry-go-extra/otelzap v0.1.5
	go.opentelemetry.io/otel v1.2.0
	go.uber.org/zap v1.19.1
)

require github.com/uptrace/opentelemetry-go-extra/otelplay v0.1.5
