module github.com/uptrace/opentelemetry-go-extra/otelgorm/e2e

go 1.16

replace github.com/uptrace/opentelemetry-go-extra/otelgorm => ./..

replace github.com/uptrace/opentelemetry-go-extra/otelsql => ./../../otelsql

require (
	github.com/stretchr/testify v1.7.0
	github.com/uptrace/opentelemetry-go-extra/otelgorm v0.1.3
	go.opentelemetry.io/otel v1.1.0
	go.opentelemetry.io/otel/sdk v1.0.1
	gorm.io/driver/sqlite v1.1.6
	gorm.io/gorm v1.22.2
)
