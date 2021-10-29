package otelgraphql

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/graph-gophers/graphql-go/errors"
	"github.com/graph-gophers/graphql-go/introspection"
	"github.com/graph-gophers/graphql-go/trace"

	otelcontrib "go.opentelemetry.io/contrib"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	oteltrace "go.opentelemetry.io/otel/trace"
)

type Tracer struct {
	tracer oteltrace.Tracer
}

const tracerName = "github.com/uptrace/opentelemetry-go-extra/otelgraphql-go"

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

func (t Tracer) TraceQuery(ctx context.Context,
	queryString string, operationName string,
	variables map[string]interface{},
	varTypes map[string]*introspection.Type) (context.Context, trace.TraceQueryFinishFunc) {
	spanCtx, span := t.tracer.Start(ctx, "GraphQL request",
		oteltrace.WithSpanKind(oteltrace.SpanKindServer),
	)
	span.SetAttributes(attribute.String("trace.operation", "request"))
	span.SetAttributes(attribute.String("graphql.query", queryString))

	if operationName != "" {
		span.SetAttributes(attribute.String("graphql.operationName", operationName))
	}

	if len(variables) != 0 {
		for name, value := range variables {
			span.SetAttributes(attrAny("graphql.variables."+name, value))
		}
	}

	return spanCtx, func(errs []*errors.QueryError) {
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

func (t Tracer) TraceField(ctx context.Context,
	label,
	typeName,
	fieldName string,
	trivial bool,
	args map[string]interface{}) (context.Context, trace.TraceFieldFinishFunc) {
	if trivial {
		return ctx, func(*errors.QueryError) {}
	}

	spanCtx, span := t.tracer.Start(ctx, label)
	span.SetAttributes(attribute.String("trace.operation", "field"))
	span.SetAttributes(attribute.String("graphql.type", typeName))
	span.SetAttributes(attribute.String("graphql.field", fieldName))
	for name, value := range args {
		span.SetAttributes(attrAny("graphql.args."+name, value))
	}

	return spanCtx, func(err *errors.QueryError) {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}
}

func (t Tracer) TraceValidation(ctx context.Context) trace.TraceValidationFinishFunc {

	_, span := t.tracer.Start(ctx, "Validate query")
	span.SetAttributes(attribute.String("trace.operation", "validation"))

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

// stolen from otellogrus
func attrAny(key string, value interface{}) attribute.KeyValue {
	switch value := value.(type) {
	case nil:
		return attribute.String(key, "<nil>")
	case string:
		return attribute.String(key, value)
	case int:
		return attribute.Int(key, value)
	case int64:
		return attribute.Int64(key, value)
	case uint64:
		return attribute.Int64(key, int64(value))
	case float64:
		return attribute.Float64(key, value)
	case bool:
		return attribute.Bool(key, value)
	case fmt.Stringer:
		return attribute.String(key, value.String())
	}

	rv := reflect.ValueOf(value)

	switch rv.Kind() {
	case reflect.Array:
		rv = rv.Slice(0, rv.Len())
		fallthrough
	case reflect.Slice:
		switch reflect.TypeOf(value).Elem().Kind() {
		case reflect.Bool:
			return attribute.BoolSlice(key, rv.Interface().([]bool))
		case reflect.Int:
			return attribute.IntSlice(key, rv.Interface().([]int))
		case reflect.Int64:
			return attribute.Int64Slice(key, rv.Interface().([]int64))
		case reflect.Float64:
			return attribute.Float64Slice(key, rv.Interface().([]float64))
		case reflect.String:
			return attribute.StringSlice(key, rv.Interface().([]string))
		default:
			return attribute.KeyValue{Key: attribute.Key(key)}
		}
	case reflect.Bool:
		return attribute.Bool(key, rv.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return attribute.Int64(key, rv.Int())
	case reflect.Float64:
		return attribute.Float64(key, rv.Float())
	case reflect.String:
		return attribute.String(key, rv.String())
	}
	if b, err := json.Marshal(value); b != nil && err == nil {
		return attribute.String(key, string(b))
	}
	return attribute.String(key, fmt.Sprint(value))
}
