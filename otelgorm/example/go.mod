module github.com/uptrace/opentelemetry-go-extra/otelgorm/example

go 1.16

replace github.com/uptrace/opentelemetry-go-extra/otelsql => ../../otelsql

replace github.com/uptrace/opentelemetry-go-extra/otelgorm => ./..

require (
	github.com/uptrace/opentelemetry-go-extra/otelgorm v0.1.2
	go.opentelemetry.io/otel v1.0.1
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.0.1
	go.opentelemetry.io/otel/sdk v1.0.1
	gorm.io/driver/sqlite v1.1.6
	gorm.io/gorm v1.21.16
)

require (
	github.com/mattn/go-sqlite3 v1.14.9 // indirect
	golang.org/x/sys v0.0.0-20211025201205-69cdffdb9359 // indirect
)
