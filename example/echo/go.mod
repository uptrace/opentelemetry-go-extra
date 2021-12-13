module github.com/uptrace/opentelemetry-go-extra/example/echo

go 1.16

replace github.com/uptrace/opentelemetry-go-extra/otelplay => ../../otelplay

require (
	github.com/labstack/echo/v4 v4.6.1
	github.com/labstack/gommon v0.3.1 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/stretchr/objx v0.1.1 // indirect
	github.com/uptrace/opentelemetry-go-extra/otelplay v0.1.7
	go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho v0.28.0
	go.opentelemetry.io/otel/trace v1.3.0
	golang.org/x/crypto v0.0.0-20211209193657-4570a0811e8b // indirect
	golang.org/x/time v0.0.0-20211116232009-f0f3c7e86c11 // indirect
)
