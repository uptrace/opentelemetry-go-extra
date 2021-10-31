package otelgraphql

import (
	"errors"
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/stretchr/testify/require"

	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
)

const schemaString = `
	schema {
		query: Query
	}
	type Query {
		echo(message: String!): String!
		echo2(message: String!): String!
		echoError(): String!
	}
	`

type RootResolver struct{}

func (*RootResolver) Echo(args struct{ Message string }) string {
	return args.Message
}

func (r *RootResolver) Echo2(args struct{ Message string }) string {
	return r.Echo(args)
}

func (r *RootResolver) EchoError() (string, error) {
	return "", errors.New("echo error")
}

type fixture struct {
	SpanRecorder   *tracetest.SpanRecorder
	GraphQLHandler *relay.Handler
}

func newFixture() *fixture {
	sr := tracetest.NewSpanRecorder()
	tp := tracesdk.NewTracerProvider()
	tp.RegisterSpanProcessor(sr)

	tracer := NewTracer(WithTracerProvider(tp))

	opts := []graphql.SchemaOpt{
		graphql.Tracer(tracer),
		graphql.UseFieldResolvers(),
	}
	schema := graphql.MustParseSchema(schemaString, &RootResolver{}, opts...)

	handler := &relay.Handler{Schema: schema}

	return &fixture{
		SpanRecorder:   sr,
		GraphQLHandler: handler,
	}
}

func (f *fixture) getSpans(query string, vars string) []tracesdk.ReadOnlySpan {
	request := fmt.Sprintf(`{"query":"%s",
		"variables":%s}`, query, vars)
	body := strings.NewReader(request)
	r := httptest.NewRequest("GET", "/graphql", body)
	w := httptest.NewRecorder()

	f.GraphQLHandler.ServeHTTP(w, r)

	return f.SpanRecorder.Ended()
}

func TestForSingleFieldTrace(t *testing.T) {
	query := "query Echo($message: String!) { echo (message: $message) }"
	vars := "{\"message\": \"Hello\"}"
	spans := newFixture().getSpans(query, vars)

	require.Len(t, spans, 3)
	require.Equal(t, "graphql.Validate", spans[0].Name())
	require.Equal(t, "GraphQL field: Query.echo", spans[1].Name())
	require.Equal(t, "graphql.Echo", spans[2].Name())
}

func TestForTwoFieldTraces(t *testing.T) {
	query := "query Echo($message1: String!, $message2: String!) { echo (message: $message1)\\necho2 (message: $message2) }"
	vars := "{\"message1\": \"Hello\", \"message2\": \"World\"}}"
	spans := newFixture().getSpans(query, vars)

	require.Len(t, spans, 4)
	require.Equal(t, "graphql.Validate", spans[0].Name())
	require.ElementsMatch(t,
		[]string{"GraphQL field: Query.echo", "GraphQL field: Query.echo2"},
		[]string{spans[1].Name(), spans[2].Name()})
	require.Equal(t, "graphql.Echo", spans[3].Name())
}

func TestForValidationTraceWithError(t *testing.T) {
	query := "query { nonExistingFieldToTriggerValidationError }"
	vars := "{}"
	spans := newFixture().getSpans(query, vars)

	require.Len(t, spans, 1)
	span := spans[0]
	require.Equal(t, "graphql.Validate", span.Name())

	events := span.Events()
	require.Equal(t, 1, len(events))
	event := events[0]
	require.Equal(t, "exception", event.Name)
}

func TestForRequestTraceWithError(t *testing.T) {
	query := "query { echoError }"
	vars := "{}"
	spans := newFixture().getSpans(query, vars)

	require.Len(t, spans, 3)
	require.Equal(t, "graphql.Validate", spans[0].Name())
	require.Equal(t, "GraphQL field: Query.echoError", spans[1].Name())
	require.Equal(t, "graphql.Request", spans[2].Name())

	events := spans[2].Events()
	require.Len(t, events, 1)
	event := events[0]
	require.Equal(t, "exception", event.Name)
}
