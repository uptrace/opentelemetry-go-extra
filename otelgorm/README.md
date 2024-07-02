[![PkgGoDev](https://pkg.go.dev/badge/github.com/uptrace/opentelemetry-go-extra/otelgorm)](https://pkg.go.dev/github.com/uptrace/opentelemetry-go-extra/otelgorm)

# GORM OpenTelemetry instrumentation

[OpenTelemetry GORM instrumentation](https://uptrace.dev/get/instrument/opentelemetry-gorm.html)
records database queries and reports `DBStats` metrics.

## Installation

```shell
go get github.com/uptrace/opentelemetry-go-extra/otelgorm
```

## Usage

To instrument GORM, you need to install the plugin provided by otelgorm:

```go
import (
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
if err != nil {
	panic(err)
}

if err := db.Use(otelgorm.NewPlugin()); err != nil {
	panic(err)
}
```

And then use `db.WithContext(ctx)` to propagate the active span via
[context](https://uptrace.dev/opentelemetry/go-tracing.html#context):

```go
var num int
if err := db.WithContext(ctx).Raw("SELECT 42").Scan(&num).Error; err != nil {
	panic(err)
}
```

See [example](/example/) for details.

## Usage: exceptions

In case you want to have certain database-related queries not initiate starting a new trace, you can set use this key to do so.
```go
db.Set(otelgorm.EnabledTraceDatabaseKey, false)
```
This will make sure no new trace is started.
Do note that if a trace was already started earlier within the context that was passed, the queries will still be visible inside.

## Options

You can customize the plugin using configuration
[options](https://pkg.go.dev/github.com/uptrace/opentelemetry-go-extra/otelgorm#Option):

- [WithAttributes](https://pkg.go.dev/github.com/uptrace/opentelemetry-go-extra/otelgorm#WithAttributes)
  configures attributes that are used to create a span.
- [WithDBName](https://pkg.go.dev/github.com/uptrace/opentelemetry-go-extra/otelgorm#WithDBName)
  configures a `db.name` attribute.

For example:

```go
otelPlugin := otelgorm.NewPlugin(otelgorm.WithDBName("mydb"))

if err := db.Use(otelPlugin); err != nil {
	panic(err)
}
```
