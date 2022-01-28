module github.com/uptrace/opentelemetry-go-extra/otelzap/example

go 1.17

replace github.com/uptrace/opentelemetry-go-extra/otelzap => ../

replace github.com/uptrace/opentelemetry-go-extra/otelutil => ../../otelutil

replace github.com/uptrace/opentelemetry-go-extra/otelplay => ../../otelplay

require (
	github.com/uptrace/opentelemetry-go-extra/otelzap v0.1.7
	go.opentelemetry.io/otel v1.3.0
	go.uber.org/zap v1.20.0
)

require github.com/uptrace/opentelemetry-go-extra/otelplay v0.1.7

require (
	github.com/cenkalti/backoff/v4 v4.1.2 // indirect
	github.com/go-logr/logr v1.2.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/uptrace/opentelemetry-go-extra/otelutil v0.1.7 // indirect
	github.com/uptrace/uptrace-go v1.3.1 // indirect
	go.opentelemetry.io/contrib/instrumentation/runtime v0.27.0 // indirect
	go.opentelemetry.io/otel/exporters/jaeger v1.3.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/internal/retry v1.3.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric v0.26.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc v0.26.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.3.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.3.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.3.0 // indirect
	go.opentelemetry.io/otel/internal/metric v0.26.0 // indirect
	go.opentelemetry.io/otel/metric v0.26.0 // indirect
	go.opentelemetry.io/otel/sdk v1.3.0 // indirect
	go.opentelemetry.io/otel/sdk/export/metric v0.26.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v0.26.0 // indirect
	go.opentelemetry.io/otel/trace v1.3.0 // indirect
	go.opentelemetry.io/proto/otlp v0.12.0 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	golang.org/x/sys v0.0.0-20220114195835-da31bd327af9 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220126215142-9970aeb2e350 // indirect
	google.golang.org/grpc v1.44.0 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)
