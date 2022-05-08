package otellogrus

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/attribute"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
)

type Test struct {
	log     func(context.Context)
	require func(sdktrace.Event)
}

func TestOtelLogrus(t *testing.T) {
	tests := []Test{
		{
			log: func(ctx context.Context) {
				logrus.WithContext(ctx).Info("hello")
			},
			require: func(event sdktrace.Event) {
				m := attrMap(event.Attributes)

				sev, ok := m[logSeverityKey]
				require.True(t, ok)
				require.Equal(t, "INFO", sev.AsString())

				msg, ok := m[logMessageKey]
				require.True(t, ok)
				require.Equal(t, "hello", msg.AsString())
			},
		},
		{
			log: func(ctx context.Context) {
				logrus.WithContext(ctx).WithField("foo", "bar").Warn("hello")
			},
			require: func(event sdktrace.Event) {
				m := attrMap(event.Attributes)

				sev, ok := m[logSeverityKey]
				require.True(t, ok)
				require.Equal(t, "WARN", sev.AsString())

				msg, ok := m[logMessageKey]
				require.True(t, ok)
				require.Equal(t, "hello", msg.AsString())

				foo, ok := m["foo"]
				require.True(t, ok)
				require.Equal(t, "bar", foo.AsString())
			},
		},
		{
			log: func(ctx context.Context) {
				err := errors.New("some error")
				logrus.WithContext(ctx).WithError(err).Error("hello")
			},
			require: func(event sdktrace.Event) {
				m := attrMap(event.Attributes)

				sev, ok := m[logSeverityKey]
				require.True(t, ok)
				require.Equal(t, "ERROR", sev.AsString())

				msg, ok := m[logMessageKey]
				require.True(t, ok)
				require.Equal(t, "hello", msg.AsString())

				excTyp, ok := m[semconv.ExceptionTypeKey]
				require.True(t, ok)
				require.Equal(t, "*errors.errorString", excTyp.AsString())

				excMsg, ok := m[semconv.ExceptionMessageKey]
				require.True(t, ok)
				require.Equal(t, "some error", excMsg.AsString())
			},
		},
		{
			log: func(ctx context.Context) {
				logrus.SetReportCaller(true)
				logrus.WithContext(ctx).Info("hello")
				logrus.SetReportCaller(false)
			},
			require: func(event sdktrace.Event) {
				m := attrMap(event.Attributes)

				fn, ok := m[semconv.CodeFunctionKey]
				require.True(t, ok)
				require.Contains(t, fn.AsString(), "github.com/uptrace/opentelemetry-go-extra/otellogrus.TestOtelLogrus")

				file, ok := m[semconv.CodeFilepathKey]
				require.True(t, ok)
				require.Contains(t, file.AsString(), "otellogrus/otellogrus_test.go")

				_, ok = m[semconv.CodeLineNumberKey]
				require.True(t, ok)
			},
		},
	}

	logrus.AddHook(NewHook(WithLevels(
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
	)))

	for i, test := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			sr := tracetest.NewSpanRecorder()
			provider := sdktrace.NewTracerProvider(sdktrace.WithSpanProcessor(sr))
			tracer := provider.Tracer("test")

			ctx := context.Background()
			ctx, span := tracer.Start(ctx, "main")

			test.log(ctx)

			span.End()

			spans := sr.Ended()
			require.Equal(t, 1, len(spans))

			events := spans[0].Events()
			require.Equal(t, 1, len(events))

			event := events[0]
			require.Equal(t, "log", event.Name)
			test.require(event)
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
