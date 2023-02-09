module github.com/uptrace/opentelemetry-go-extra/example/grpc

go 1.18

replace github.com/uptrace/opentelemetry-go-extra/otelplay => ../../otelplay

require (
	github.com/golang/protobuf v1.5.2
	github.com/uptrace/opentelemetry-go-extra/otelplay v0.1.20
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.26.0
	go.opentelemetry.io/otel/trace v1.13.0
	google.golang.org/grpc v1.53.0
	google.golang.org/protobuf v1.28.1
)

require (
	github.com/cenkalti/backoff/v4 v4.2.0 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.15.0 // indirect
	github.com/uptrace/uptrace-go v1.13.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/runtime v0.39.0 // indirect
	go.opentelemetry.io/otel v1.13.0 // indirect
	go.opentelemetry.io/otel/exporters/jaeger v1.13.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/internal/retry v1.13.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric v0.36.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc v0.36.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.13.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.13.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.13.0 // indirect
	go.opentelemetry.io/otel/metric v0.36.0 // indirect
	go.opentelemetry.io/otel/sdk v1.13.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v0.36.0 // indirect
	go.opentelemetry.io/proto/otlp v0.19.0 // indirect
	golang.org/x/net v0.6.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	google.golang.org/genproto v0.0.0-20230202175211-008b39050e57 // indirect
)
