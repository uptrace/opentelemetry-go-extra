module github.com/uptrace/opentelemetry-go-extra/otelzap

go 1.17

replace github.com/uptrace/opentelemetry-go-extra/otelutil => ../otelutil

require (
	github.com/stretchr/testify v1.7.1
	github.com/uptrace/opentelemetry-go-extra/otelutil v0.1.14
	go.opentelemetry.io/otel v1.7.0
	go.opentelemetry.io/otel/sdk v1.0.1
	go.opentelemetry.io/otel/trace v1.7.0
	go.uber.org/zap v1.21.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	golang.org/x/sys v0.0.0-20210510120138-977fb7262007 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)
