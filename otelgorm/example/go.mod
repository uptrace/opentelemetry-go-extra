module github.com/uptrace/opentelemetry-go-extra/otelgorm/example

go 1.16

replace github.com/uptrace/opentelemetry-go-extra/otelsql => ../../otelsql

replace github.com/uptrace/opentelemetry-go-extra/otelgorm => ./..

replace github.com/uptrace/opentelemetry-go-extra/otelplay => ../../otelplay

require (
	github.com/uptrace/opentelemetry-go-extra/otelgorm v0.1.4
	gorm.io/driver/sqlite v1.2.4
	gorm.io/gorm v1.22.3
)

require (
	github.com/uptrace/opentelemetry-go-extra/otelplay v0.1.4
	go.opentelemetry.io/otel v1.2.0
)
