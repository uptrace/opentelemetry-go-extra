module github.com/uptrace/opentelemetry-go-extra/otelgorm

go 1.18

replace github.com/uptrace/opentelemetry-go-extra/otelsql => ../otelsql

require (
	github.com/uptrace/opentelemetry-go-extra/otelsql v0.1.16
	go.opentelemetry.io/otel v1.10.0
	go.opentelemetry.io/otel/trace v1.10.0
	gorm.io/gorm v1.23.10
)

require (
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	go.opentelemetry.io/otel/metric v0.32.1 // indirect
)
