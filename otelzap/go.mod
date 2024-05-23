module github.com/uptrace/opentelemetry-go-extra/otelzap

go 1.22

toolchain go1.22.3

replace github.com/uptrace/opentelemetry-go-extra/otelutil => ../otelutil

require (
	github.com/uptrace/opentelemetry-go-extra/otelutil v0.3.0
	go.opentelemetry.io/otel v1.27.0
	go.opentelemetry.io/otel/log v0.3.0
	go.opentelemetry.io/otel/trace v1.27.0
	go.uber.org/zap v1.27.0
)

require (
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	go.opentelemetry.io/otel/metric v1.27.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
)
