module github.com/uptrace/opentelemetry-go-extra/otelzap/example

go 1.16

replace github.com/uptrace/opentelemetry-go-extra/otelzap => ../

replace github.com/uptrace/opentelemetry-go-extra/otelplay => ../../otelplay

require (
	github.com/uptrace/opentelemetry-go-extra/otelzap v0.1.2
	go.opentelemetry.io/otel v1.1.0
	go.uber.org/zap v1.19.1
)

require (
	github.com/uptrace/opentelemetry-go-extra/otelplay v0.0.0-00010101000000-000000000000
	golang.org/x/sys v0.0.0-20211029165221-6e7872819dc8 // indirect
)
