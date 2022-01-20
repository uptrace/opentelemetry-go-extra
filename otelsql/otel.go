package otelsql

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"io"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/global"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.opentelemetry.io/otel/trace"
)

const instrumName = "github.com/uptrace/opentelemetry-go-extra/otelsql"

var dbRowsAffected = attribute.Key("db.rows_affected")

type config struct {
	provider trace.TracerProvider
	tracer   trace.Tracer //nolint:structcheck
	meter    metric.Meter
	attrs    []attribute.KeyValue
}

func newConfig(opts []Option) *config {
	c := &config{}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func (c *config) formatQuery(query string) string {
	return query
}

type dbInstrum struct {
	*config

	queryHistogram metric.Int64Histogram
}

func newDBInstrum(opts []Option) *dbInstrum {
	t := &dbInstrum{
		config: newConfig(opts),
	}

	if t.provider == nil {
		t.provider = otel.GetTracerProvider()
	}
	if t.tracer == nil {
		t.tracer = t.provider.Tracer(instrumName)
	}

	if t.meter.MeterImpl() == nil {
		t.meter = global.Meter(instrumName)
	}

	meter := metric.Must(t.meter)
	t.queryHistogram = meter.NewInt64Histogram(
		"go.sql.query_timing",
		metric.WithDescription("Timing of processed queries"),
		metric.WithUnit("milliseconds"),
	)

	return t
}

func (t *dbInstrum) withSpan(
	ctx context.Context,
	spanName string,
	query string,
	fn func(ctx context.Context, span trace.Span) error,
) error {
	var startTime time.Time
	if query != "" {
		startTime = time.Now()
	}

	attrs := make([]attribute.KeyValue, 0, len(t.attrs)+1)
	attrs = append(attrs, t.attrs...)
	if query != "" {
		attrs = append(attrs, semconv.DBStatementKey.String(t.formatQuery(query)))
	}

	ctx, span := t.tracer.Start(ctx, spanName,
		trace.WithSpanKind(trace.SpanKindClient),
		trace.WithAttributes(attrs...))
	err := fn(ctx, span)
	span.End()

	if query != "" {
		t.queryHistogram.Record(ctx, time.Since(startTime).Milliseconds(), t.attrs...)
	}

	if !span.IsRecording() {
		return err
	}

	switch err {
	case nil,
		driver.ErrSkip,
		io.EOF, // end of rows iterator
		sql.ErrNoRows:
		// ignore
	default:
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}

	return err
}

type Option func(c *config)

// WithTracerProvider configures a tracer provider that is used to create a tracer.
func WithTracerProvider(provider trace.TracerProvider) Option {
	return func(c *config) {
		c.provider = provider
	}
}

// WithAttributes configures attributes that are used to create a span.
func WithAttributes(attrs ...attribute.KeyValue) Option {
	return func(c *config) {
		c.attrs = append(c.attrs, attrs...)
	}
}

// WithDBSystem configures a db.system attribute. You should prefer using
// WithAttributes and semconv, for example, `otelsql.WithAttributes(semconv.DBSystemSqlite)`.
func WithDBSystem(system string) Option {
	return func(c *config) {
		c.attrs = append(c.attrs, semconv.DBSystemKey.String(system))
	}
}

// WithDBName configures a db.name attribute.
func WithDBName(name string) Option {
	return func(c *config) {
		c.attrs = append(c.attrs, semconv.DBNameKey.String(name))
	}
}

// WithMeter configures a metric.Meter used to create instruments.
func WithMeter(meter metric.Meter) Option {
	return func(c *config) {
		c.meter = meter
	}
}

// ReportDBStatsMetrics reports DBStats metrics using OpenTelemetry Metrics API.
func ReportDBStatsMetrics(db *sql.DB, opts ...Option) {
	cfg := newConfig(opts)

	if cfg.meter.MeterImpl() == nil {
		cfg.meter = global.Meter(instrumName)
	}

	meter := metric.Must(cfg.meter)
	labels := cfg.attrs

	var maxOpenConns metric.Int64GaugeObserver
	var openConns metric.Int64GaugeObserver
	var inUseConns metric.Int64GaugeObserver
	var idleConns metric.Int64GaugeObserver
	var connsWaitCount metric.Int64CounterObserver
	var connsWaitDuration metric.Int64CounterObserver
	var connsClosedMaxIdle metric.Int64CounterObserver
	var connsClosedMaxIdleTime metric.Int64CounterObserver
	var connsClosedMaxLifetime metric.Int64CounterObserver

	batch := meter.NewBatchObserver(func(ctx context.Context, result metric.BatchObserverResult) {
		stats := db.Stats()

		result.Observe(labels,
			maxOpenConns.Observation(int64(stats.MaxOpenConnections)),

			openConns.Observation(int64(stats.OpenConnections)),
			inUseConns.Observation(int64(stats.InUse)),
			idleConns.Observation(int64(stats.Idle)),

			connsWaitCount.Observation(stats.WaitCount),
			connsWaitDuration.Observation(int64(stats.WaitDuration)),
			connsClosedMaxIdle.Observation(stats.MaxIdleClosed),
			connsClosedMaxIdleTime.Observation(stats.MaxIdleTimeClosed),
			connsClosedMaxLifetime.Observation(stats.MaxLifetimeClosed),
		)
	})

	maxOpenConns = batch.NewInt64GaugeObserver("go.sql.connections_max_open",
		metric.WithDescription("Maximum number of open connections to the database"),
	)
	openConns = batch.NewInt64GaugeObserver("go.sql.connections_open",
		metric.WithDescription("The number of established connections both in use and idle"),
	)
	inUseConns = batch.NewInt64GaugeObserver("go.sql.connections_in_use",
		metric.WithDescription("The number of connections currently in use"),
	)
	idleConns = batch.NewInt64GaugeObserver("go.sql.connections_idle",
		metric.WithDescription("The number of idle connections"),
	)
	connsWaitCount = batch.NewInt64CounterObserver("go.sql.connections_wait_count",
		metric.WithDescription("The total number of connections waited for"),
	)
	connsWaitDuration = batch.NewInt64CounterObserver("go.sql.connections_wait_duration",
		metric.WithDescription("The total time blocked waiting for a new connection"),
		metric.WithUnit("nanoseconds"),
	)
	connsClosedMaxIdle = batch.NewInt64CounterObserver("go.sql.connections_closed_max_idle",
		metric.WithDescription("The total number of connections closed due to SetMaxIdleConns"),
	)
	connsClosedMaxIdleTime = batch.NewInt64CounterObserver("go.sql.connections_closed_max_idle_time",
		metric.WithDescription("The total number of connections closed due to SetConnMaxIdleTime"),
	)
	connsClosedMaxLifetime = batch.NewInt64CounterObserver("go.sql.connections_closed_max_lifetime",
		metric.WithDescription("The total number of connections closed due to SetConnMaxLifetime"),
	)
}
