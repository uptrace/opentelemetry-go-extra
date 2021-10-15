[![PkgGoDev](https://pkg.go.dev/badge/github.com/uptrace/opentelemetry-go-extra/otelsql)](https://pkg.go.dev/github.com/uptrace/opentelemetry-go-extra/otelsql)

# OpenTelemetry Go instrumentation for database/sql package

otelsql instrumentation records database queries (including `Tx` and `Stmt` queries) and reports
`DBStats` metrics.

## Installation

```shell
go get github.com/uptrace/opentelemetry-go-extra/otelsql
```

## Usage

To instrument database/sql, you need to connect to a database using the API provided by this
package:

- `sql.Open(driverName, dsn)` becomes `otelsql.Open(driverName, dsn)`.
- `sql.OpenDB(connector)` becomes `otelsql.OpenDB(connector)`.

```go
import (
	"github.com/uptrace/opentelemetry-go-extra/otelsql"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

db, err := otelsql.Open("sqlite", "file::memory:?cache=shared",
	otelsql.WithAttributes(semconv.DBSystemSqlite),
	otelsql.WithDBName("mydb"))
if err != nil {
	panic(err)
}
```

And then use context-aware API to propagate the active span via
[context](https://docs.uptrace.dev/guide/go.html#context):

```go
var num int
if err := db.QueryRowContext(ctx, "SELECT 42").Scan(&num); err != nil {
	panic(err)
}
```

See [example](/example/) for details.

## Options

Both [otelsql.Open](https://pkg.go.dev/github.com/uptrace/opentelemetry-go-extra/otelsql#Open) and
[otelsql.OpenDB](https://pkg.go.dev/github.com/uptrace/opentelemetry-go-extra/otelsql#OpenDB) accept
the same [options](https://pkg.go.dev/github.com/uptrace/opentelemetry-go-extra/otelsql#Option):

- [WithAttributes](https://pkg.go.dev/github.com/uptrace/opentelemetry-go-extra/otelsql#WithAttributes)
  configures attributes that are used to create a span.
- [WithDBName](https://pkg.go.dev/github.com/uptrace/opentelemetry-go-extra/otelsql#WithDBName)
  configures a `db.name` attribute.
- [WithDBSystem](https://pkg.go.dev/github.com/uptrace/opentelemetry-go-extra/otelsql#WithDBSystem)
  configures a `db.system` attribute. When possible, you should prefer using WithAttributes and
  [semconv](https://pkg.go.dev/go.opentelemetry.io/otel/semconv/v1.4.0), for example,
  `otelsql.WithAttributes(semconv.DBSystemSqlite)`.

## Alternatives

- https://github.com/XSAM/otelsql - different driver registration and no metrics.
- https://github.com/j2gg0s/otsql - like XSAM/otelsql but with Prometheus metrics.
