module github.com/uptrace/opentelemetry-go-extra/otelgorm

go 1.16

replace github.com/uptrace/opentelemetry-go-extra/otelsql => ../otelsql

require (
	github.com/uptrace/opentelemetry-go-extra/otelsql v0.1.3
	go.opentelemetry.io/otel v1.1.0
	go.opentelemetry.io/otel/trace v1.1.0
	gorm.io/gorm v1.22.2
)
