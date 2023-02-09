module github.com/uptrace/opentelemetry-go-extra/otelgraphql

go 1.18

replace github.com/uptrace/opentelemetry-go-extra/otelutil => ../otelutil

require (
	github.com/graph-gophers/graphql-go v1.5.0
	github.com/stretchr/testify v1.8.1
	github.com/uptrace/opentelemetry-go-extra/otelutil v0.1.21
	go.opentelemetry.io/contrib v1.14.0
	go.opentelemetry.io/otel v1.13.0
	go.opentelemetry.io/otel/sdk v1.1.0
	go.opentelemetry.io/otel/trace v1.13.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/sys v0.0.0-20210423185535-09eb48e85fd7 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
