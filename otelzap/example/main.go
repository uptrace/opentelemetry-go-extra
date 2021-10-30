package main

import (
	"context"
	"errors"
	"sync"

	"go.opentelemetry.io/otel"
	"go.uber.org/zap"

	"github.com/uptrace/opentelemetry-go-extra/otelplay"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
)

func main() {
	ctx := context.Background()

	shutdown := otelplay.ConfigureOpentelemetry(ctx)
	defer shutdown()

	tracer := otel.Tracer("app_or_package_name")

	ctx, span := tracer.Start(ctx, "root")
	defer span.End()

	// Use Ctx to propagate the active span.
	Logger(ctx).Error("hello from zap",
		zap.Error(errors.New("hello world")),
		zap.String("foo", "bar"))

	otelplay.PrintTraceID(ctx)
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
