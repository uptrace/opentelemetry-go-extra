module github.com/uptrace/opentelemetry-go-extra/example/echo

go 1.16

replace github.com/uptrace/opentelemetry-go-extra/otelplay => ../../otelplay

require (
	github.com/labstack/echo/v4 v4.6.1
	github.com/labstack/gommon v0.3.1 // indirect
	github.com/stretchr/objx v0.1.1 // indirect
	github.com/uptrace/opentelemetry-go-extra/otelplay v0.1.4
	go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho v0.26.1
	go.opentelemetry.io/otel/trace v1.1.0
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
	golang.org/x/time v0.0.0-20210723032227-1f47c861a9ac // indirect
)
