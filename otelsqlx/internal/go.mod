module github.com/uptrace/opentelemetry-go-extra/otelsqlx/internal

go 1.16

replace github.com/uptrace/opentelemetry-go-extra/otelsqlx => ../

replace github.com/uptrace/opentelemetry-go-extra/otelsql => ../../otelsql

require (
	github.com/stretchr/testify v1.7.0
	github.com/uptrace/opentelemetry-go-extra/otelsql v0.1.1
	github.com/uptrace/opentelemetry-go-extra/otelsqlx v0.0.0-00010101000000-000000000000
	go.opentelemetry.io/otel v1.0.1
	go.opentelemetry.io/otel/sdk v1.0.1
	modernc.org/sqlite v1.13.3
)
