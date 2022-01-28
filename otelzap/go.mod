module github.com/uptrace/opentelemetry-go-extra/otelzap

go 1.17

replace github.com/uptrace/opentelemetry-go-extra/otelutil => ../otelutil

require (
	github.com/stretchr/testify v1.7.0
	github.com/uptrace/opentelemetry-go-extra/otelutil v0.1.7
	go.opentelemetry.io/otel v1.3.0
	go.opentelemetry.io/otel/sdk v1.0.1
	go.opentelemetry.io/otel/trace v1.3.0
	go.uber.org/zap v1.20.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-logr/logr v1.2.1 // indirect
	github.com/go-logr/stdr v1.2.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	golang.org/x/sys v0.0.0-20210510120138-977fb7262007 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)
