module github.com/uptrace/opentelemetry-go-extra/otelsqlx

go 1.16

replace github.com/uptrace/opentelemetry-go-extra/otelsql => ../otelsql

require (
	github.com/jmoiron/sqlx v1.3.4
	github.com/uptrace/opentelemetry-go-extra/otelsql v0.1.1
)
