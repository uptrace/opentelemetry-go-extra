module github.com/uptrace/opentelemetry-go-extra/example/gorilla-mux

go 1.16

replace github.com/uptrace/opentelemetry-go-extra/otelplay => ../../otelplay

require (
	github.com/gorilla/mux v1.8.0
	github.com/stretchr/objx v0.1.1 // indirect
	github.com/uptrace/opentelemetry-go-extra/otelplay v0.1.4
	go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux v0.27.0
	go.opentelemetry.io/otel/trace v1.2.0
)
