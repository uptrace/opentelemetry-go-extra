[![PkgGoDev](https://pkg.go.dev/badge/github.com/uptrace/opentelemetry-go-extra/otelzap)](https://pkg.go.dev/github.com/uptrace/opentelemetry-go-extra/otelzap)

# OpenTelemetry Go instrumentation for Zap logging library

This instrumentation records Zap log messages as events on the existing span that must be passed in
a `context.Context` as a first argument. It does not record anything if the context does not contain
a span.

## Installation

```shell
go get github.com/uptrace/opentelemetry-go-extra/otelzap
```

## Usage

You need to create an `otelzap.Logger` using this package and pass a
[context](https://docs.uptrace.dev/guide/go.html#context) to propage the active span.

```go
import (
    "go.uber.org/zap"
    "github.com/uptrace/opentelemetry-go-extra/otelzap"
)

// Wrap zap logger to extend Zap with API that accepts a context.Context.
log := otelzap.New(zap.L())

// And then pass ctx to propagate the span.
log.Ctx(ctx).Error("hello from zap",
	zap.Error(errors.New("hello world")),
	zap.String("foo", "bar"))

// Alternatively.
log.ErrorContext(ctx, "hello from zap",
	zap.Error(errors.New("hello world")),
	zap.String("foo", "bar"))
```

Both variants are fast and don't allocate. See [example](/example/) for details.

### Sugared logger

You can also use sugared logger API in a similar way:

```go
log := otelzap.New(zap.L())
sugar := log.Sugar()

sugar.Ctx(ctx).Infow("failed to fetch URL",
  // Structured context as loosely typed key-value pairs.
  "url", url,
  "attempt", 3,
  "backoff", time.Second,
)
sugar.InfowContext(ctx, "failed to fetch URL",
  // Structured context as loosely typed key-value pairs.
  "url", url,
  "attempt", 3,
  "backoff", time.Second,
)

sugar.Ctx(ctx).Infof("Failed to fetch URL: %s", url)
sugar.InfofContext(ctx, "Failed to fetch URL: %s", url)
```

## Options

[otelzap.New](https://pkg.go.dev/github.com/uptrace/opentelemetry-go-extra/otelzap#New) accepts a
couple of [options](https://pkg.go.dev/github.com/uptrace/opentelemetry-go-extra/otelzap#Option):

- `otelzap.WithMinLevel(zap.WarnLevel)` sets the minimal zap logging level on which the log message
  is recorded on the span.
- `otelzap.WithErrorStatusLevel(zap.ErrorLevel)` sets the minimal zap logging level on which the
  span status is set to codes.Error.
- `otelzap.WithCaller(true)` configures the logger to annotate each event with the filename, line
  number, and function name of the caller. Enabled by default.
- `otelzap.WithStackTrace(true)` configures the logger to capture logs with a stack trace. Disabled
  by default.
- `otelzap.WithTraceIDField(true)` configures the logger to add `trace_id` field to structured log
  messages. This option is only useful with backends that don't support OTLP and instead parse log
  messages to extract structured information.
