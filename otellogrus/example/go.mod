module github.com/uptrace/opentelemetry-go-extra/otellogrus/example

go 1.16

replace github.com/uptrace/opentelemetry-go-extra/otellogrus => ./..

replace github.com/uptrace/opentelemetry-go-extra/otelutil => ../../otelutil

replace github.com/uptrace/opentelemetry-go-extra/otelplay => ../../otelplay

require (
	github.com/sirupsen/logrus v1.8.1
	github.com/uptrace/opentelemetry-go-extra/otellogrus v0.1.6
	github.com/uptrace/opentelemetry-go-extra/otelplay v0.1.6
	go.opentelemetry.io/otel v1.2.0
	golang.org/x/net v0.0.0-20211123203042-d83791d6bcd9 // indirect
	google.golang.org/genproto v0.0.0-20211118181313-81c1377c94b1 // indirect
)
