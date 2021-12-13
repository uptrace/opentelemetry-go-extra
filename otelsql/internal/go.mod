module github.com/uptrace/opentelemetry-go-extra/otelsql/internal

go 1.16

replace github.com/uptrace/opentelemetry-go-extra/otelsql => ../

require (
	github.com/stretchr/testify v1.7.0
	github.com/uptrace/opentelemetry-go-extra/otelsql v0.1.7
	go.opentelemetry.io/otel v1.3.0
	go.opentelemetry.io/otel/sdk v1.0.1
	go.opentelemetry.io/otel/trace v1.3.0
	modernc.org/sqlite v1.13.1
)
