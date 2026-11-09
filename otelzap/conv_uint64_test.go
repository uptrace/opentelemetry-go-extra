package otelzap

import (
	"context"
	"math"
	"testing"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/log"
	"go.opentelemetry.io/otel/log/embedded"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
)

type recordingProvider struct {
	embedded.LoggerProvider
	logger *recordingLogger
}

func (p *recordingProvider) Logger(name string, _ ...log.LoggerOption) log.Logger {
	return p.logger
}

type recordingLogger struct {
	embedded.Logger
	records []log.Record
}

func (l *recordingLogger) Enabled(context.Context, log.EnabledParameters) bool { return true }

func (l *recordingLogger) Emit(_ context.Context, r log.Record) {
	l.records = append(l.records, r)
}

func attrString(t *testing.T, r log.Record, key string) string {
	t.Helper()
	var got string
	var found bool
	r.WalkAttributes(func(kv attribute.KeyValue) bool {
		if string(kv.Key) == key {
			got = kv.Value.String()
			found = true
			return false
		}
		return true
	})
	if !found {
		t.Fatalf("attribute %q not emitted", key)
	}
	return got
}

func newRecordingLogger(t *testing.T) (*Logger, *recordingLogger) {
	t.Helper()
	rec := &recordingLogger{}
	l := New(zaptest.NewLogger(t),
		WithLoggerProvider(&recordingProvider{logger: rec}),
		WithMinLevel(zap.InfoLevel),
	)
	return l, rec
}

func TestUint64FieldAboveMaxInt64IsNotEmittedAsNegative(t *testing.T) {
	l, rec := newRecordingLogger(t)
	l.InfoContext(context.Background(), "hello", zap.Uint64("hash", math.MaxUint64))

	if len(rec.records) != 1 {
		t.Fatalf("want 1 record, got %d", len(rec.records))
	}
	if got, want := attrString(t, rec.records[0], "hash"), "18446744073709551615"; got != want {
		t.Fatalf("hash = %s, want %s", got, want)
	}
}

func TestUint64FieldBelowMaxInt64StillEmittedAsInt(t *testing.T) {
	l, rec := newRecordingLogger(t)
	l.InfoContext(context.Background(), "hello", zap.Uint64("n", 42))

	if got, want := attrString(t, rec.records[0], "n"), "42"; got != want {
		t.Fatalf("n = %s, want %s", got, want)
	}
}

func TestSignedIntFieldIsUnchanged(t *testing.T) {
	l, rec := newRecordingLogger(t)
	l.InfoContext(context.Background(), "hello", zap.Int64("n", -7))

	if got, want := attrString(t, rec.records[0], "n"), "-7"; got != want {
		t.Fatalf("n = %s, want %s", got, want)
	}
}
