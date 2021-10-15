package main

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	_ "modernc.org/sqlite"

	"github.com/uptrace/opentelemetry-go-extra/otelsql"
)

func main() {
	ctx := context.Background()

	stop := configureOpentelemetry(ctx)
	defer stop()

	db, err := otelsql.Open("sqlite", "file::memory:?cache=shared",
		otelsql.WithAttributes(semconv.DBSystemSqlite),
		otelsql.WithDBName("mydb"))
	if err != nil {
		panic(err)
	}

	if err := db.PingContext(ctx); err != nil {
		panic(err)
	}

	var num int
	if err := db.QueryRowContext(ctx, "SELECT 42").Scan(&num); err != nil {
		panic(err)
	}
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
