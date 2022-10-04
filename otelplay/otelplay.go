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
	fmt.Println("trace:", TraceURL(trace.SpanFromContext(ctx)))
}

func TraceURL(span trace.Span) string {
	switch {
	case os.Getenv("UPTRACE_DSN") != "":
		return uptrace.TraceURL(span)
	case os.Getenv("OTEL_EXPORTER_JAEGER_ENDPOINT") != "":
		return fmt.Sprintf("http://localhost:16686/trace/%s", span.SpanContext().TraceID())
	default:
		return fmt.Sprintf("http://localhost:16686/trace/%s", span.SpanContext().TraceID())
	}
}

// ConfigureOpentelemetry configures Opentelemetry to export spans to Uptrace, Jaeger,
// or console depending on environment variables.
//
// You can use it to run examples, but don't use it in your applications. Instead, use
// uptrace-go or opentelemetry-go directly. See https://uptrace.dev/get/uptrace-go.html
func ConfigureOpentelemetry(ctx context.Context) func() {
	switch {
	case os.Getenv("UPTRACE_DSN") != "":
		uptrace.ConfigureOpentelemetry(uptrace.WithServiceName("myservicename"))
		return func() {
			_ = uptrace.Shutdown(ctx)
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
