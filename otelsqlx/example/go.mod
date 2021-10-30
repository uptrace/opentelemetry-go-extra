module github.com/uptrace/opentelemetry-go-extra/otelsqlx/example

go 1.16

replace github.com/uptrace/opentelemetry-go-extra/otelsqlx => ../

replace github.com/uptrace/opentelemetry-go-extra/otelsql => ../../otelsql

replace github.com/uptrace/opentelemetry-go-extra/otelplay => ../../otelplay

require (
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/uptrace/opentelemetry-go-extra/otelplay v0.0.0-00010101000000-000000000000
	github.com/uptrace/opentelemetry-go-extra/otelsql v0.1.2
	github.com/uptrace/opentelemetry-go-extra/otelsqlx v0.1.2
	go.opentelemetry.io/otel v1.1.0
	golang.org/x/mod v0.5.1 // indirect
	golang.org/x/sys v0.0.0-20211029165221-6e7872819dc8 // indirect
	golang.org/x/tools v0.1.7 // indirect
	modernc.org/ccgo/v3 v3.12.49 // indirect
	modernc.org/sqlite v1.13.3
)
