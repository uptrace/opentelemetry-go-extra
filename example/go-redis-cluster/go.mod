module github.com/uptrace/opentelemetry-go-extra/example/go-redis-cluster

go 1.16

replace github.com/uptrace/opentelemetry-go-extra/otelplay => ../../otelplay

require (
	github.com/go-redis/redis/extra/redisotel/v8 v8.11.4
	github.com/go-redis/redis/v8 v8.11.4
	github.com/uptrace/opentelemetry-go-extra/otelplay v0.1.6
	go.opentelemetry.io/otel v1.2.0
	golang.org/x/net v0.0.0-20211123203042-d83791d6bcd9 // indirect
	golang.org/x/sys v0.0.0-20211124211545-fe61309f8881 // indirect
	google.golang.org/genproto v0.0.0-20211118181313-81c1377c94b1 // indirect
)
