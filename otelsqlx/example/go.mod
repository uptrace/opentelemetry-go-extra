module github.com/uptrace/opentelemetry-go-extra/otelsqlx/example

go 1.16

replace github.com/uptrace/opentelemetry-go-extra/otelsqlx => ../

replace github.com/uptrace/opentelemetry-go-extra/otelsql => ../../otelsql

replace github.com/uptrace/opentelemetry-go-extra/otelplay => ../../otelplay

require (
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/uptrace/opentelemetry-go-extra/otelplay v0.1.6
	github.com/uptrace/opentelemetry-go-extra/otelsql v0.1.6
	github.com/uptrace/opentelemetry-go-extra/otelsqlx v0.1.6
	go.opentelemetry.io/otel v1.3.0
	golang.org/x/tools v0.1.8 // indirect
	modernc.org/ccgo/v3 v3.12.95 // indirect
	modernc.org/sqlite v1.14.2
)
