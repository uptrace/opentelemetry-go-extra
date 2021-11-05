module github.com/uptrace/opentelemetry-go-extra/otelgorm/example

go 1.16

replace github.com/uptrace/opentelemetry-go-extra/otelsql => ../../otelsql

replace github.com/uptrace/opentelemetry-go-extra/otelgorm => ./..

replace github.com/uptrace/opentelemetry-go-extra/otelplay => ../../otelplay

require (
	github.com/uptrace/opentelemetry-go-extra/otelgorm v0.1.3
	gorm.io/driver/sqlite v1.2.3
	gorm.io/gorm v1.22.2
)

require (
	github.com/uptrace/opentelemetry-go-extra/otelplay v0.1.3
	go.opentelemetry.io/otel v1.1.0
)
