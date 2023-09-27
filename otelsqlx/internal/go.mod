module github.com/uptrace/opentelemetry-go-extra/otelsqlx/internal

go 1.18

replace github.com/uptrace/opentelemetry-go-extra/otelsqlx => ../

replace github.com/uptrace/opentelemetry-go-extra/otelsql => ../../otelsql

require (
	github.com/stretchr/testify v1.8.4
	github.com/uptrace/opentelemetry-go-extra/otelsql v0.2.3
	github.com/uptrace/opentelemetry-go-extra/otelsqlx v0.2.3
	go.opentelemetry.io/otel v1.18.0
	go.opentelemetry.io/otel/sdk v1.15.1
	modernc.org/sqlite v1.20.3
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/jmoiron/sqlx v1.3.5 // indirect
	github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20200410134404-eec4a21b6bb0 // indirect
	go.opentelemetry.io/otel/metric v1.18.0 // indirect
	go.opentelemetry.io/otel/trace v1.18.0 // indirect
	golang.org/x/mod v0.3.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/tools v0.0.0-20201124115921-2c860bdd6e78 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	lukechampine.com/uint128 v1.2.0 // indirect
	modernc.org/cc/v3 v3.40.0 // indirect
	modernc.org/ccgo/v3 v3.16.13 // indirect
	modernc.org/libc v1.22.2 // indirect
	modernc.org/mathutil v1.5.0 // indirect
	modernc.org/memory v1.4.0 // indirect
	modernc.org/opt v0.1.3 // indirect
	modernc.org/strutil v1.1.3 // indirect
	modernc.org/token v1.0.1 // indirect
)
