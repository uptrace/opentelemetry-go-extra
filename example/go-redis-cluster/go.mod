module github.com/uptrace/opentelemetry-go-extra/example/go-redis-cluster

go 1.16

replace github.com/uptrace/opentelemetry-go-extra/otelplay => ../../otelplay

require (
	github.com/go-redis/redis/extra/redisotel/v8 v8.11.4
	github.com/go-redis/redis/v8 v8.11.4
	github.com/uptrace/opentelemetry-go-extra/otelplay v0.1.7
	go.opentelemetry.io/otel v1.3.0
)
