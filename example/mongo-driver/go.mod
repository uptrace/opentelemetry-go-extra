module github.com/uptrace/opentelemetry-go-extra/example/mongo-driver

go 1.16

replace github.com/uptrace/opentelemetry-go-extra/otelplay => ../../otelplay

require (
	github.com/uptrace/opentelemetry-go-extra/otelplay v0.0.0-00010101000000-000000000000
	go.mongodb.org/mongo-driver v1.7.3
	go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo v0.26.0
	go.opentelemetry.io/otel v1.1.0
)
