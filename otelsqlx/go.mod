module github.com/uptrace/opentelemetry-go-extra/otelsqlx

go 1.17

replace github.com/uptrace/opentelemetry-go-extra/otelsql => ../otelsql

exclude go.opentelemetry.io/proto/otlp v0.15.0

require (
	github.com/jmoiron/sqlx v1.3.4
	github.com/uptrace/opentelemetry-go-extra/otelsql v0.1.11
)

require (
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	go.opentelemetry.io/otel v1.6.1 // indirect
	go.opentelemetry.io/otel/metric v0.28.0 // indirect
	go.opentelemetry.io/otel/trace v1.6.1 // indirect
)
