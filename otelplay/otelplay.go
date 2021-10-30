package otelplay

import (
	"context"
	"fmt"
	"os"

	"github.com/uptrace/uptrace-go/uptrace"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

func PrintTraceID(ctx context.Context) {
	span := trace.SpanFromContext(ctx)

	switch {
	case os.Getenv("UPTRACE_DSN") != "":
		fmt.Println("trace:", uptrace.TraceURL(span))
	case os.Getenv("OTEL_EXPORTER_JAEGER_ENDPOINT") != "":
		fmt.Printf("trace: http://localhost:16686/trace/%s\n", span.SpanContext().TraceID())
	default:
		// nothing
	}
}

func ConfigureOpentelemetry(ctx context.Context) func() {
	switch {
	case os.Getenv("UPTRACE_DSN") != "":
		uptrace.ConfigureOpentelemetry()
		return func() {
			uptrace.Shutdown(ctx)
		}
	case os.Getenv("OTEL_EXPORTER_JAEGER_ENDPOINT") != "":
		return configureJaeger(ctx)
	default:
		return configureStdout(ctx)
	}
}

func configureJaeger(ctx context.Context) func() {
	provider := sdktrace.NewTracerProvider()
	otel.SetTracerProvider(provider)

	exp, err := jaeger.New(jaeger.WithCollectorEndpoint())
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

func configureStdout(ctx context.Context) func() {
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
