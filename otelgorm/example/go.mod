module github.com/uptrace/opentelemetry-go-extra/otelgorm/example

go 1.16

replace github.com/uptrace/opentelemetry-go-extra/otelsql => ../../otelsql

replace github.com/uptrace/opentelemetry-go-extra/otelgorm => ./..

require (
	github.com/uptrace/opentelemetry-go-extra/otelgorm v0.1.2
	go.opentelemetry.io/otel v1.1.0
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.1.0
	go.opentelemetry.io/otel/sdk v1.1.0
	gorm.io/driver/sqlite v1.2.3
	gorm.io/gorm v1.22.2
)

require golang.org/x/sys v0.0.0-20211029165221-6e7872819dc8 // indirect
