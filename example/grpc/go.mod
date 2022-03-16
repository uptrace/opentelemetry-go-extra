module github.com/uptrace/opentelemetry-go-extra/example/grpc

go 1.17

replace github.com/uptrace/opentelemetry-go-extra/otelplay => ../../otelplay

require (
	github.com/golang/protobuf v1.5.2
	github.com/uptrace/opentelemetry-go-extra/otelplay v0.1.10
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.26.0
	go.opentelemetry.io/otel/trace v1.4.1
	google.golang.org/grpc v1.45.0
	google.golang.org/protobuf v1.27.1
)

require (
	github.com/cenkalti/backoff/v4 v4.1.2 // indirect
	github.com/go-logr/logr v1.2.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/stretchr/objx v0.1.1 // indirect
	github.com/uptrace/uptrace-go v1.4.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/runtime v0.29.0 // indirect
	go.opentelemetry.io/otel v1.4.1 // indirect
	go.opentelemetry.io/otel/exporters/jaeger v1.4.1 // indirect
	go.opentelemetry.io/otel/exporters/otlp/internal/retry v1.4.1 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric v0.27.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc v0.27.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.4.1 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.4.1 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.4.1 // indirect
	go.opentelemetry.io/otel/internal/metric v0.27.0 // indirect
	go.opentelemetry.io/otel/metric v0.27.0 // indirect
	go.opentelemetry.io/otel/sdk v1.4.1 // indirect
	go.opentelemetry.io/otel/sdk/metric v0.27.0 // indirect
	go.opentelemetry.io/proto/otlp v0.12.0 // indirect
	golang.org/x/net v0.0.0-20220225172249-27dd8689420f // indirect
	golang.org/x/sys v0.0.0-20220315194320-039c03cc5b86 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220314164441-57ef72a4c106 // indirect
)
