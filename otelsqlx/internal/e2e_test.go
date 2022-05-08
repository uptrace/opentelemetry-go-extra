package internal_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
	_ "modernc.org/sqlite"

	"github.com/uptrace/opentelemetry-go-extra/otelsql"
	"github.com/uptrace/opentelemetry-go-extra/otelsqlx"
)

type Person struct {
	ID   int64
	Name string
}

func TestE2E(t *testing.T) {
	sr := tracetest.NewSpanRecorder()
	provider := sdktrace.NewTracerProvider(sdktrace.WithSpanProcessor(sr))

	db, err := otelsqlx.Open("sqlite", "file::memory:?cache=shared",
		otelsql.WithAttributes(semconv.DBSystemSqlite),
		otelsql.WithDBName("mydb"),
		otelsql.WithTracerProvider(provider))
	require.NoError(t, err)

	person := new(Person)
	err = db.Get(person, "SELECT 123 AS id, 'hello' AS name")
	require.NoError(t, err)
	require.Equal(t, int64(123), person.ID)
	require.Equal(t, "hello", person.Name)

	spans := sr.Ended()
	require.Equal(t, 2, len(spans))
	require.Equal(t, "db.Connect", spans[0].Name())
	require.Equal(t, "db.Query", spans[1].Name())
}
