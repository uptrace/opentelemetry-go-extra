module github.com/uptrace/opentelemetry-go-extra/otelgorm/internal/e2e

go 1.17

replace github.com/uptrace/opentelemetry-go-extra/otelgorm => ./../..

replace github.com/uptrace/opentelemetry-go-extra/otelsql => ./../../../otelsql

require (
	github.com/stretchr/testify v1.7.1
	github.com/uptrace/opentelemetry-go-extra/otelgorm v0.1.14
	go.opentelemetry.io/otel v1.7.0
	go.opentelemetry.io/otel/sdk v1.0.1
	go.opentelemetry.io/otel/trace v1.7.0
	gorm.io/driver/sqlite v1.1.6
	gorm.io/gorm v1.23.5
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-sqlite3 v1.14.8 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/uptrace/opentelemetry-go-extra/otelsql v0.1.14 // indirect
	go.opentelemetry.io/otel/metric v0.30.0 // indirect
	golang.org/x/sys v0.0.0-20210423185535-09eb48e85fd7 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c // indirect
)
