[![PkgGoDev](https://pkg.go.dev/badge/github.com/uptrace/opentelemetry-go-extra/otelsql)](https://pkg.go.dev/github.com/uptrace/opentelemetry-go-extra/otelsqlx)

# sqlx instrumentation for OpenTelemetry Go

otelsqlx instrumentation records [sqlx](https://github.com/jmoiron/sqlx) queries (including `Tx` and
`Stmt` queries) and reports `DBStats` metrics.

```shell
go get github.com/uptrace/opentelemetry-go-extra/otelsqlx
```

## Usage

To instrument sqlx, you need to connect to a database using the API provided by this package:

| sqlx                  | otelsqlx                  |
| --------------------- | ------------------------- |
| `sqlx.Connect`        | `otelsqlx.Connect`        |
| `sqlx.ConnectContext` | `otelsqlx.ConnectContext` |
| `sqlx.MustConnect`    | `otelsqlx.MustConnect`    |
| `sqlx.Open`           | `otelsqlx.Open`           |
| `sqlx.MustOpen`       | `otelsqlx.MustOpen`       |
| `sqlx.NewDb`          | not supported             |

```go
import (
    "github.com/uptrace/opentelemetry-go-extra/otelsqlx"
    semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
    _ "modernc.org/sqlite"
)

db, err := otelsqlx.Open("sqlite", "file::memory:?cache=shared",
	otelsql.WithAttributes(semconv.DBSystemSqlite))
if err != nil {
	panic(err)
}

// db is *sqlx.DB
```

And then use context-aware API to propagate the active span via
[context](https://uptrace.dev/opentelemetry/go-tracing.html#context):

```go
var num int
if err := db.QueryRowContext(ctx, "SELECT 42").Scan(&num); err != nil {
	panic(err)
}
```

See [example](/example/) for details.

## Options

otelsqlx accepts all the options from
[otelsql](https://github.com/uptrace/opentelemetry-go-extra/tree/main/otelsql) package, for example:

```go
import (
    "github.com/uptrace/opentelemetry-go-extra/otelsqlx"
    semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
    _ "modernc.org/sqlite"
)

db, err := otelsqlx.Open("sqlite", "file::memory:?cache=shared",
	otelsql.WithAttributes(semconv.DBSystemSqlite),
	otelsql.WithDBName("mydb"))
```
