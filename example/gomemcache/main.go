package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
	"go.opentelemetry.io/contrib/instrumentation/github.com/bradfitz/gomemcache/memcache/otelmemcache"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"

	"github.com/uptrace/opentelemetry-go-extra/otelplay"
)

var tracer = otel.Tracer("app_or_package_name")

func main() {
	flag.Parse()

	ctx := context.Background()

	shutdown := otelplay.ConfigureOpentelemetry(ctx)
	defer shutdown()

	mc := otelmemcache.NewClientWithTracing(
		memcache.New(":11211"),
	)

	ctx, span := tracer.Start(ctx, "test-operations")
	defer span.End()

	doMemcacheOperations(ctx, mc)

	fmt.Println("trace", otelplay.TraceURL(span))
}

func doMemcacheOperations(ctx context.Context, mc *otelmemcache.Client) {
	mc = mc.WithContext(ctx)

	err := mc.Add(&memcache.Item{
		Key:   "foo",
		Value: []byte("bar"),
	})
	if err != nil {
		trace.SpanFromContext(ctx).RecordError(err)
	}

	_, err = mc.Get("foo")
	if err != nil {
		trace.SpanFromContext(ctx).RecordError(err)
	}

	_, err = mc.Get("hello")
	if err != nil {
		trace.SpanFromContext(ctx).RecordError(err)
	}

	err = mc.Delete("foo")
	if err != nil {
		trace.SpanFromContext(ctx).RecordError(err)
	}
}
