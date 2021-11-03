module github.com/uptrace/opentelemetry-go-extra/example/echo

go 1.16

replace github.com/uptrace/opentelemetry-go-extra/otelplay => ../../otelplay

require (
	github.com/labstack/echo/v4 v4.6.1
	github.com/stretchr/objx v0.1.1 // indirect
	github.com/uptrace/opentelemetry-go-extra/otelplay v0.0.0-00010101000000-000000000000
	go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho v0.26.0
	go.opentelemetry.io/otel/trace v1.1.0
)
