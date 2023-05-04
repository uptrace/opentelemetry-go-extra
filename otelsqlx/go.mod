module github.com/uptrace/opentelemetry-go-extra/otelsqlx

go 1.18

replace github.com/uptrace/opentelemetry-go-extra/otelsql => ../otelsql

require (
	github.com/jmoiron/sqlx v1.3.5
	github.com/uptrace/opentelemetry-go-extra/otelsql v0.1.21
)

require (
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	go.opentelemetry.io/otel v1.15.1 // indirect
	go.opentelemetry.io/otel/metric v0.38.1 // indirect
	go.opentelemetry.io/otel/trace v1.15.1 // indirect
)
