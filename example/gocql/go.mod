module github.com/uptrace/opentelemetry-go-extra/example/gocql

go 1.16

replace github.com/uptrace/opentelemetry-go-extra/otelplay => ../../otelplay

require (
	github.com/gocql/gocql v0.0.0-20211015133455-b225f9b53fa1
	github.com/uptrace/opentelemetry-go-extra/otelplay v0.0.0-00010101000000-000000000000
	go.opentelemetry.io/contrib/instrumentation/github.com/gocql/gocql/otelgocql v0.26.0
	go.opentelemetry.io/otel v1.1.0
	go.opentelemetry.io/otel/trace v1.1.0
)
