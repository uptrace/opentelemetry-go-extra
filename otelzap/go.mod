module github.com/uptrace/opentelemetry-go-extra/otelzap

go 1.15

replace github.com/uptrace/opentelemetry-go-extra/otelutil => ../otelutil

require (
	github.com/stretchr/testify v1.7.0
	github.com/uptrace/opentelemetry-go-extra/otelutil v0.1.6
	go.opentelemetry.io/otel v1.2.0
	go.opentelemetry.io/otel/sdk v1.0.1
	go.opentelemetry.io/otel/trace v1.2.0
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	go.uber.org/zap v1.19.1
)
