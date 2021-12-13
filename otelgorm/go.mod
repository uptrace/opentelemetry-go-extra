module github.com/uptrace/opentelemetry-go-extra/otelgorm

go 1.16

replace github.com/uptrace/opentelemetry-go-extra/otelsql => ../otelsql

require (
	github.com/jinzhu/now v1.1.4 // indirect
	github.com/uptrace/opentelemetry-go-extra/otelsql v0.1.6
	go.opentelemetry.io/otel v1.3.0
	go.opentelemetry.io/otel/internal/metric v0.26.0 // indirect
	go.opentelemetry.io/otel/trace v1.3.0
	gorm.io/gorm v1.22.4
)
