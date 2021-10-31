module github.com/uptrace/opentelemetry-go-extra/otellogrus/example

go 1.16

replace github.com/uptrace/opentelemetry-go-extra/otellogrus => ./..

replace github.com/uptrace/opentelemetry-go-extra/otelplay => ../../otelplay

require (
	github.com/sirupsen/logrus v1.8.1
	github.com/uptrace/opentelemetry-go-extra/otellogrus v0.1.3
	github.com/uptrace/opentelemetry-go-extra/otelplay v0.1.3
	go.opentelemetry.io/contrib/instrumentation/runtime v0.26.0 // indirect
	go.opentelemetry.io/otel v1.1.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.1.0 // indirect
	golang.org/x/net v0.0.0-20211029224645-99673261e6eb // indirect
	google.golang.org/genproto v0.0.0-20211029142109-e255c875f7c7 // indirect
)
