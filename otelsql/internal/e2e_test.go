package internal_test

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/attribute"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
	"go.opentelemetry.io/otel/trace"
	_ "modernc.org/sqlite"

	"github.com/uptrace/opentelemetry-go-extra/otelsql"
)

var dbRowsAffected = attribute.Key("db.rows_affected")

type Test struct {
	do      func(ctx context.Context, db *sql.DB)
	require func(t *testing.T, spans []sdktrace.ReadOnlySpan)
}

func TestConn(t *testing.T) {
	tests := []Test{
		{
			do: func(ctx context.Context, db *sql.DB) {
				err := db.PingContext(ctx)
				require.NoError(t, err)
			},
			require: func(t *testing.T, spans []sdktrace.ReadOnlySpan) {
				require.Equal(t, 2, len(spans))
				require.Equal(t, "db.Connect", spans[0].Name())
				require.Equal(t, "db.Ping", spans[1].Name())
			},
		},
		{
			do: func(ctx context.Context, db *sql.DB) {
				_, err := db.ExecContext(ctx, "SELECT 1")
				require.NoError(t, err)
			},
			require: func(t *testing.T, spans []sdktrace.ReadOnlySpan) {
				require.Equal(t, 2, len(spans))
				require.Equal(t, "db.Connect", spans[0].Name())
				require.Equal(t, "db.Exec", spans[1].Name())

				span := spans[1]
				require.Equal(t, "db.Exec", span.Name())

				m := attrMap(span.Attributes())

				stmt, ok := m[semconv.DBStatementKey]
				require.True(t, ok)
				require.Equal(t, "SELECT 1", stmt.AsString())

				rows, ok := m[dbRowsAffected]
				require.True(t, ok)
				require.Equal(t, int64(0), rows.AsInt64())
			},
		},
		{
			do: func(ctx context.Context, db *sql.DB) {
				var num int
				err := db.QueryRowContext(ctx, "SELECT 1").Scan(&num)
				require.NoError(t, err)
				require.Equal(t, 1, num)
			},
			require: func(t *testing.T, spans []sdktrace.ReadOnlySpan) {
				require.Equal(t, 2, len(spans))
				require.Equal(t, "db.Connect", spans[0].Name())
				require.Equal(t, "db.Query", spans[1].Name())

				span := spans[1]
				require.Equal(t, "db.Query", span.Name())

				m := attrMap(span.Attributes())

				stmt, ok := m[semconv.DBStatementKey]
				require.True(t, ok)
				require.Equal(t, "SELECT 1", stmt.AsString())
			},
		},
		{
			do: func(ctx context.Context, db *sql.DB) {
				stmt, err := db.PrepareContext(ctx, "SELECT 1")
				require.NoError(t, err)

				_, err = stmt.ExecContext(ctx)
				require.NoError(t, err)

				var num int
				err = stmt.QueryRowContext(ctx).Scan(&num)
				require.NoError(t, err)
				require.Equal(t, 1, num)
			},
			require: func(t *testing.T, spans []sdktrace.ReadOnlySpan) {
				wanted := []struct {
					name string
					stmt string
				}{
					{name: "db.Connect", stmt: ""},
					{name: "db.Prepare", stmt: "SELECT 1"},
					{name: "stmt.Exec", stmt: "SELECT 1"},
					{name: "stmt.Query", stmt: "SELECT 1"},
				}
				for i, wanted := range wanted {
					span := spans[i]
					require.Equal(t, wanted.name, span.Name())

					m := attrMap(span.Attributes())

					stmt, ok := m[semconv.DBStatementKey]
					require.Equal(t, wanted.stmt != "", ok)
					require.Equal(t, wanted.stmt, stmt.AsString())
				}
			},
		},
		{
			do: func(ctx context.Context, db *sql.DB) {
				tx, err := db.BeginTx(ctx, nil)
				require.NoError(t, err)

				_, err = tx.ExecContext(ctx, "SELECT 1")
				require.NoError(t, err)

				var num int
				err = tx.QueryRowContext(ctx, "SELECT 1").Scan(&num)
				require.NoError(t, err)
				require.Equal(t, 1, num)

				err = tx.Rollback()
				require.NoError(t, err)
			},
			require: func(t *testing.T, spans []sdktrace.ReadOnlySpan) {
				require.Equal(t, 5, len(spans))

				wanted := []struct {
					name string
					stmt string
				}{
					{name: "db.Connect", stmt: ""},
					{name: "db.Begin", stmt: ""},
					{name: "db.Exec", stmt: "SELECT 1"},
					{name: "db.Query", stmt: "SELECT 1"},
					{name: "tx.Rollback", stmt: ""},
				}
				for i, wanted := range wanted {
					span := spans[i]
					require.Equal(t, wanted.name, span.Name())

					m := attrMap(span.Attributes())

					stmt, ok := m[semconv.DBStatementKey]
					require.Equal(t, wanted.stmt != "", ok)
					require.Equal(t, wanted.stmt, stmt.AsString())
				}
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			sr := tracetest.NewSpanRecorder()
			provider := sdktrace.NewTracerProvider(sdktrace.WithSpanProcessor(sr))

			db, err := otelsql.Open("sqlite", "file::memory:?cache=shared",
				otelsql.WithTracerProvider(provider))
			require.NoError(t, err)

			test.do(context.TODO(), db)

			spans := sr.Ended()
			for _, span := range spans {
				require.Equal(t, trace.SpanKindClient, span.SpanKind())
			}
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
