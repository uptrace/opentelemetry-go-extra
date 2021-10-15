package main

import (
	"context"
	"errors"
	"sync"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/zap"

	"github.com/uptrace/opentelemetry-go-extra/otelzap"
)

func main() {
	ctx := context.Background()

	stop := configureOpentelemetry(ctx)
	defer stop()

	tracer := otel.Tracer("app_or_package_name")

	ctx, span := tracer.Start(ctx, "main")
	defer span.End()

	// Use Ctx to propagate the active span.
	Logger(ctx).Error("hello from zap",
		zap.Error(errors.New("hello world")),
		zap.String("foo", "bar"))
}

var (
	once   sync.Once
	logger *otelzap.Logger
)

// Logger ensures that the caller does not forget to pass the context.
func Logger(ctx context.Context) otelzap.LoggerWithCtx {
	once.Do(func() {
		l, err := zap.NewDevelopment()
		if err != nil {
			panic(err)
		}
		logger = otelzap.New(l)
	})
	return logger.Ctx(ctx)
}

func configureOpentelemetry(ctx context.Context) func() {
	provider := sdktrace.NewTracerProvider()
	otel.SetTracerProvider(provider)

	exp, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		panic(err)
	}

	bsp := sdktrace.NewBatchSpanProcessor(exp)
	provider.RegisterSpanProcessor(bsp)

	return func() {
		if err := provider.Shutdown(ctx); err != nil {
			panic(err)
		}
	}
}
