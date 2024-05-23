module github.com/uptrace/opentelemetry-go-extra/otelgraphql

go 1.22

toolchain go1.22.3

replace github.com/uptrace/opentelemetry-go-extra/otelutil => ../otelutil

require (
	github.com/graph-gophers/graphql-go v1.5.0
	github.com/stretchr/testify v1.9.0
	github.com/uptrace/opentelemetry-go-extra/otelutil v0.3.0
	go.opentelemetry.io/contrib v1.27.0
	go.opentelemetry.io/otel v1.27.0
	go.opentelemetry.io/otel/sdk v1.24.0
	go.opentelemetry.io/otel/trace v1.27.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.opentelemetry.io/otel/log v0.3.0 // indirect
	go.opentelemetry.io/otel/metric v1.27.0 // indirect
	golang.org/x/sys v0.17.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
