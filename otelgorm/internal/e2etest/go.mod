module github.com/uptrace/opentelemetry-go-extra/otelgorm/internal/e2e

go 1.21

toolchain go1.22.1

replace github.com/uptrace/opentelemetry-go-extra/otelgorm => ./../..

replace github.com/uptrace/opentelemetry-go-extra/otelsql => ./../../../otelsql

require (
	github.com/stretchr/testify v1.8.4
	github.com/uptrace/opentelemetry-go-extra/otelgorm v0.2.3
	go.opentelemetry.io/otel v1.24.0
	go.opentelemetry.io/otel/sdk v1.24.0
	go.opentelemetry.io/otel/trace v1.24.0
	gorm.io/driver/sqlite v1.4.4
	gorm.io/gorm v1.25.7
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-sqlite3 v1.14.15 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/uptrace/opentelemetry-go-extra/otelsql v0.2.3 // indirect
	go.opentelemetry.io/otel/metric v1.24.0 // indirect
	golang.org/x/sys v0.17.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
