package e2etest

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
)

type Test struct {
	do      func(ctx context.Context, db *gorm.DB)
	require func(t *testing.T, spans []sdktrace.ReadOnlySpan)
}

func TestEndToEnd(t *testing.T) {
	tests := []Test{
		{
			do: func(ctx context.Context, db *gorm.DB) {
				var num int
				err := db.WithContext(ctx).Raw("SELECT 42").Scan(&num).Error
				require.NoError(t, err)
			},
			require: func(t *testing.T, spans []sdktrace.ReadOnlySpan) {
				require.Equal(t, 1, len(spans))
				require.Equal(t, "gorm.Row", spans[0].Name())
				require.Equal(t, trace.SpanKindClient, spans[0].SpanKind())

				m := attrMap(spans[0].Attributes())

				sys, ok := m[semconv.DBSystemKey]
				require.True(t, ok)
				require.Equal(t, "sqlite", sys.AsString())

				stmt, ok := m[semconv.DBStatementKey]
				require.True(t, ok)
				require.Equal(t, "SELECT 42", stmt.AsString())
			},
		},
		{
			do: func(ctx context.Context, db *gorm.DB) {
				var num int
				_ = db.WithContext(ctx).Raw("SELECT foo_bar").Scan(&num).Error
			},
			require: func(t *testing.T, spans []sdktrace.ReadOnlySpan) {
				require.Equal(t, 1, len(spans))
				require.Equal(t, "gorm.Row", spans[0].Name())

				span := spans[0]
				status := span.Status()
				require.Equal(t, codes.Error, status.Code)
				require.Equal(t, "no such column: foo_bar", status.Description)

				m := attrMap(span.Attributes())

				sys, ok := m[semconv.DBSystemKey]
				require.True(t, ok)
				require.Equal(t, "sqlite", sys.AsString())

				stmt, ok := m[semconv.DBStatementKey]
				require.True(t, ok)
				require.Equal(t, "SELECT foo_bar", stmt.AsString())
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			sr := tracetest.NewSpanRecorder()
			provider := sdktrace.NewTracerProvider(sdktrace.WithSpanProcessor(sr))

			db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
			require.NoError(t, err)

			err = db.Use(otelgorm.NewPlugin(otelgorm.WithTracerProvider(provider)))
			require.NoError(t, err)

			test.do(context.TODO(), db)

			spans := sr.Ended()
			test.require(t, spans)
		})
	}
}

func attrMap(attrs []attribute.KeyValue) map[attribute.Key]attribute.Value {
	m := make(map[attribute.Key]attribute.Value, len(attrs))
	for _, kv := range attrs {
		m[kv.Key] = kv.Value
	}
	return m
}
