module github.com/uptrace/opentelemetry-go-extra/otelzap

go 1.18

replace github.com/uptrace/opentelemetry-go-extra/otelutil => ../otelutil

require (
	github.com/stretchr/testify v1.8.3
	github.com/uptrace/opentelemetry-go-extra/otelutil v0.2.2
	go.opentelemetry.io/otel v1.16.0
	go.opentelemetry.io/otel/sdk v1.15.1
	go.opentelemetry.io/otel/trace v1.16.0
	go.uber.org/zap v1.24.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.opentelemetry.io/otel/metric v1.16.0 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
