module github.com/uptrace/opentelemetry-go-extra/example/net-http

go 1.17

replace github.com/uptrace/opentelemetry-go-extra/otelplay => ../../otelplay

require (
	github.com/uptrace/opentelemetry-go-extra/otelplay v0.1.11
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.31.0
	go.opentelemetry.io/otel/trace v1.6.3
)

require (
	github.com/cenkalti/backoff/v4 v4.1.3 // indirect
	github.com/felixge/httpsnoop v1.0.2 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.10.0 // indirect
	github.com/stretchr/objx v0.1.1 // indirect
	github.com/uptrace/uptrace-go v1.6.3 // indirect
	go.opentelemetry.io/contrib/instrumentation/runtime v0.31.0 // indirect
	go.opentelemetry.io/otel v1.6.3 // indirect
	go.opentelemetry.io/otel/exporters/jaeger v1.6.3 // indirect
	go.opentelemetry.io/otel/exporters/otlp/internal/retry v1.6.3 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric v0.29.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc v0.29.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.6.3 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.6.3 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.6.3 // indirect
	go.opentelemetry.io/otel/metric v0.29.0 // indirect
	go.opentelemetry.io/otel/sdk v1.6.3 // indirect
	go.opentelemetry.io/otel/sdk/metric v0.29.0 // indirect
	go.opentelemetry.io/proto/otlp v0.15.0 // indirect
	golang.org/x/net v0.0.0-20220412020605-290c469a71a5 // indirect
	golang.org/x/sys v0.0.0-20220412071739-889880a91fd5 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220407144326-9054f6ed7bac // indirect
	google.golang.org/grpc v1.45.0 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)
