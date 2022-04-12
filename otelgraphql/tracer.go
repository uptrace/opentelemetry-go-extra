package otelgraphql

import (
	"context"
	"fmt"

	"github.com/graph-gophers/graphql-go/errors"
	"github.com/graph-gophers/graphql-go/introspection"
	"github.com/graph-gophers/graphql-go/trace/tracer"

	otelcontrib "go.opentelemetry.io/contrib"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	oteltrace "go.opentelemetry.io/otel/trace"

	"github.com/uptrace/opentelemetry-go-extra/otelutil"
)

type Tracer struct {
	tracer oteltrace.Tracer
}

const tracerName = "github.com/uptrace/opentelemetry-go-extra/otelgraphql"

func NewTracer(opts ...Option) *Tracer {
	cfg := config{}
	for _, opt := range opts {
		opt.apply(&cfg)
	}
	if cfg.TracerProvider == nil {
		cfg.TracerProvider = otel.GetTracerProvider()
	}
	tracer := cfg.TracerProvider.Tracer(
		tracerName,
		oteltrace.WithInstrumentationVersion(otelcontrib.SemVersion()),
	)
	return &Tracer{tracer: tracer}
}

func (t Tracer) TraceQuery(
	ctx context.Context,
	queryString string, operationName string,
	variables map[string]interface{},
	varTypes map[string]*introspection.Type,
) (context.Context, tracer.QueryFinishFunc) {
	var spanName string
	if operationName != "" {
		spanName = "graphql." + operationName
	} else {
		spanName = "graphql.Request"
	}

	ctx, span := t.tracer.Start(ctx, spanName,
		oteltrace.WithSpanKind(oteltrace.SpanKindServer))
	span.SetAttributes(attribute.String("graphql.query", queryString))

	for name, value := range variables {
		span.SetAttributes(otelutil.Attribute("graphql.variables."+name, value))
	}

	return ctx, func(errs []*errors.QueryError) {
		if len(errs) > 0 {
			msg := errs[0].Error()
			if len(errs) > 1 {
				msg += fmt.Sprintf(" (and %d more errors)", len(errs)-1)
			}
			span.RecordError(errs[0])
			span.SetStatus(codes.Error, msg)
		}
		span.End()
	}
}

func (t Tracer) TraceField(
	ctx context.Context,
	label,
	typeName,
	fieldName string,
	trivial bool,
	args map[string]interface{},
) (context.Context, tracer.FieldFinishFunc) {
	if trivial {
		return ctx, func(*errors.QueryError) {}
	}

	ctx, span := t.tracer.Start(ctx, label)
	span.SetAttributes(attribute.String("graphql.type", typeName))
	span.SetAttributes(attribute.String("graphql.field", fieldName))
	for name, value := range args {
		span.SetAttributes(otelutil.Attribute("graphql.args."+name, value))
	}

	return ctx, func(err *errors.QueryError) {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}
}

func (t Tracer) TraceValidation(ctx context.Context) tracer.ValidationFinishFunc {
	_, span := t.tracer.Start(ctx, "graphql.Validate")

	return func(errs []*errors.QueryError) {
		if len(errs) > 0 {
			msg := errs[0].Error()
			if len(errs) > 1 {
				msg += fmt.Sprintf(" (and %d more errors)", len(errs)-1)
			}
			span.RecordError(errs[0])
			span.SetStatus(codes.Error, msg)
		}
		span.End()
	}
}
