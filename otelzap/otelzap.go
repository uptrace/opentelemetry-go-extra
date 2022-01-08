package otelzap

import (
	"context"
	"fmt"
	"math"
	"reflect"
	"runtime"
	"strconv"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/uptrace/opentelemetry-go-extra/otelutil"
)

const numAttr = 5

var (
	logSeverityKey = attribute.Key("log.severity")
	logMessageKey  = attribute.Key("log.message")
	logTemplateKey = attribute.Key("log.template")
)

// Logger is a thin wrapper for zap.Logger that adds Ctx method.
type Logger struct {
	*zap.Logger
	skipCaller *zap.Logger

	withTraceID bool

	minLevel         zapcore.Level
	errorStatusLevel zapcore.Level

	caller     bool
	stackTrace bool
}

func New(logger *zap.Logger, opts ...Option) *Logger {
	l := &Logger{
		Logger:     logger,
		skipCaller: logger.WithOptions(zap.AddCallerSkip(1)),

		minLevel:         zap.WarnLevel,
		errorStatusLevel: zap.ErrorLevel,
		caller:           true,
	}
	for _, opt := range opts {
		opt(l)
	}
	return l
}

// WithOptions clones the current Logger, applies the supplied Options,
// and returns the resulting Logger. It's safe to use concurrently.
func (l *Logger) WithOptions(opts ...zap.Option) *Logger {
	clone := *l
	clone.Logger = l.Logger.WithOptions(opts...)
	return &clone
}

// Sugar wraps the Logger to provide a more ergonomic, but slightly slower,
// API. Sugaring a Logger is quite inexpensive, so it's reasonable for a
// single application to use both Loggers and SugaredLoggers, converting
// between them on the boundaries of performance-sensitive code.
func (l *Logger) Sugar() *SugaredLogger {
	return &SugaredLogger{
		SugaredLogger: l.Logger.Sugar(),
		skipCaller:    l.skipCaller.Sugar(),
		l:             l,
	}
}

// Clone clones the current logger applying the supplied options.
func (l *Logger) Clone(opts ...Option) *Logger {
	clone := *l
	for _, opt := range opts {
		opt(&clone)
	}
	return &clone
}

// Ctx returns a new logger with the context.
func (l *Logger) Ctx(ctx context.Context) LoggerWithCtx {
	return LoggerWithCtx{
		ctx: ctx,
		l:   l,
	}
}

func (l *Logger) DebugContext(ctx context.Context, msg string, fields ...zapcore.Field) {
	fields = l.logFields(ctx, zap.DebugLevel, msg, fields)
	l.Debug(msg, fields...)
}

func (l *Logger) InfoContext(ctx context.Context, msg string, fields ...zapcore.Field) {
	fields = l.logFields(ctx, zap.InfoLevel, msg, fields)
	l.Info(msg, fields...)
}

func (l *Logger) WarnContext(ctx context.Context, msg string, fields ...zapcore.Field) {
	fields = l.logFields(ctx, zap.WarnLevel, msg, fields)
	l.Warn(msg, fields...)
}

func (l *Logger) ErrorContext(ctx context.Context, msg string, fields ...zapcore.Field) {
	fields = l.logFields(ctx, zap.ErrorLevel, msg, fields)
	l.Error(msg, fields...)
}

func (l *Logger) DPanicContext(ctx context.Context, msg string, fields ...zapcore.Field) {
	fields = l.logFields(ctx, zap.DPanicLevel, msg, fields)
	l.DPanic(msg, fields...)
}

func (l *Logger) PanicContext(ctx context.Context, msg string, fields ...zapcore.Field) {
	fields = l.logFields(ctx, zap.PanicLevel, msg, fields)
	l.Panic(msg, fields...)
}

func (l *Logger) FatalContext(ctx context.Context, msg string, fields ...zapcore.Field) {
	fields = l.logFields(ctx, zap.FatalLevel, msg, fields)
	l.Fatal(msg, fields...)
}

func (l *Logger) logFields(
	ctx context.Context, lvl zapcore.Level, msg string, fields []zapcore.Field,
) []zapcore.Field {
	if lvl < l.minLevel {
		return fields
	}
	span := trace.SpanFromContext(ctx)
	if !span.IsRecording() {
		return fields
	}

	attrs := make([]attribute.KeyValue, 0, numAttr+len(fields))

	for _, f := range fields {
		if f.Type == zapcore.NamespaceType {
			// should this be a prefix?
			continue
		}
		attrs = appendField(attrs, f)
	}

	l.log(span, lvl, msg, attrs)

	if l.withTraceID {
		traceID := span.SpanContext().TraceID().String()
		fields = append(fields, zap.String("trace_id", traceID))
	}

	return fields
}

func (l *Logger) log(
	span trace.Span, lvl zapcore.Level, msg string, attrs []attribute.KeyValue,
) {
	attrs = append(attrs, logSeverityKey.String(levelString(lvl)))
	attrs = append(attrs, logMessageKey.String(msg))

	if l.caller {
		if fn, file, line, ok := runtimeCaller(4); ok {
			if fn != "" {
				attrs = append(attrs, semconv.CodeFunctionKey.String(fn))
			}
			if file != "" {
				attrs = append(attrs, semconv.CodeFilepathKey.String(file))
				attrs = append(attrs, semconv.CodeLineNumberKey.Int(line))
			}
		}
	}

	if l.stackTrace {
		stackTrace := make([]byte, 2048)
		n := runtime.Stack(stackTrace, false)
		attrs = append(attrs, semconv.ExceptionStacktraceKey.String(string(stackTrace[0:n])))
	}

	span.AddEvent("log", trace.WithAttributes(attrs...))

	if lvl >= l.errorStatusLevel {
		span.SetStatus(codes.Error, msg)
	}
}

func runtimeCaller(skip int) (fn, file string, line int, ok bool) {
	rpc := make([]uintptr, 1)
	n := runtime.Callers(skip+1, rpc[:])
	if n < 1 {
		return
	}
	frame, _ := runtime.CallersFrames(rpc).Next()
	return frame.Function, frame.File, frame.Line, frame.PC != 0
}

//------------------------------------------------------------------------------

// LoggerWithCtx is a wrapper for Logger that also carries a context.Context.
type LoggerWithCtx struct {
	ctx context.Context
	l   *Logger
}

// Context returns logger's context.
func (l LoggerWithCtx) Context() context.Context {
	return l.ctx
}

// Logger returns the underlying logger.
func (l LoggerWithCtx) Logger() *Logger {
	return l.l
}

// ZapLogger returns the underlying zap logger.
func (l LoggerWithCtx) ZapLogger() *zap.Logger {
	return l.l.Logger
}

// Sugar returns a sugared logger with the context.
func (l LoggerWithCtx) Sugar() SugaredLoggerWithCtx {
	return SugaredLoggerWithCtx{
		ctx: l.ctx,
		s:   l.l.Sugar(),
	}
}

// WithOptions clones the current Logger, applies the supplied Options,
// and returns the resulting Logger. It's safe to use concurrently.
func (l LoggerWithCtx) WithOptions(opts ...zap.Option) LoggerWithCtx {
	return LoggerWithCtx{
		ctx: l.ctx,
		l:   l.l.WithOptions(opts...),
	}
}

// Clone clones the current logger applying the supplied options.
func (l LoggerWithCtx) Clone(opts ...Option) LoggerWithCtx {
	return LoggerWithCtx{
		ctx: l.ctx,
		l:   l.l.Clone(opts...),
	}
}

// Debug logs a message at DebugLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (l LoggerWithCtx) Debug(msg string, fields ...zapcore.Field) {
	fields = l.l.logFields(l.ctx, zap.DebugLevel, msg, fields)
	l.l.skipCaller.Debug(msg, fields...)
}

// Info logs a message at InfoLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (l LoggerWithCtx) Info(msg string, fields ...zapcore.Field) {
	fields = l.l.logFields(l.ctx, zap.InfoLevel, msg, fields)
	l.l.skipCaller.Info(msg, fields...)
}

// Warn logs a message at WarnLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (l LoggerWithCtx) Warn(msg string, fields ...zapcore.Field) {
	fields = l.l.logFields(l.ctx, zap.WarnLevel, msg, fields)
	l.l.skipCaller.Warn(msg, fields...)
}

// Error logs a message at ErrorLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (l LoggerWithCtx) Error(msg string, fields ...zapcore.Field) {
	fields = l.l.logFields(l.ctx, zap.ErrorLevel, msg, fields)
	l.l.skipCaller.Error(msg, fields...)
}

// DPanic logs a message at DPanicLevel. The message includes any fields
// passed at the log site, as well as any fields accumulated on the logger.
//
// If the logger is in development mode, it then panics (DPanic means
// "development panic"). This is useful for catching errors that are
// recoverable, but shouldn't ever happen.
func (l LoggerWithCtx) DPanic(msg string, fields ...zapcore.Field) {
	fields = l.l.logFields(l.ctx, zap.DPanicLevel, msg, fields)
	l.l.skipCaller.DPanic(msg, fields...)
}

// Panic logs a message at PanicLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// The logger then panics, even if logging at PanicLevel is disabled.
func (l LoggerWithCtx) Panic(msg string, fields ...zapcore.Field) {
	fields = l.l.logFields(l.ctx, zap.PanicLevel, msg, fields)
	l.l.skipCaller.Panic(msg, fields...)
}

// Fatal logs a message at FatalLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// The logger then calls os.Exit(1), even if logging at FatalLevel is
// disabled.
func (l LoggerWithCtx) Fatal(msg string, fields ...zapcore.Field) {
	fields = l.l.logFields(l.ctx, zap.FatalLevel, msg, fields)
	l.l.skipCaller.Fatal(msg, fields...)
}

//------------------------------------------------------------------------------

// A SugaredLogger wraps the base Logger functionality in a slower, but less
// verbose, API. Any Logger can be converted to a SugaredLogger with its Sugar
// method.
//
// Unlike the Logger, the SugaredLogger doesn't insist on structured logging.
// For each log level, it exposes three methods: one for loosely-typed
// structured logging, one for println-style formatting, and one for
// printf-style formatting. For example, SugaredLoggers can produce InfoLevel
// output with Infow ("info with" structured context), Info, or Infof.
type SugaredLogger struct {
	*zap.SugaredLogger
	skipCaller *zap.SugaredLogger

	l *Logger
}

// Desugar unwraps a SugaredLogger, exposing the original Logger. Desugaring
// is quite inexpensive, so it's reasonable for a single application to use
// both Loggers and SugaredLoggers, converting between them on the boundaries
// of performance-sensitive code.
func (s *SugaredLogger) Desugar() *Logger {
	return s.l
}

// With adds a variadic number of fields to the logging context. It accepts a
// mix of strongly-typed Field objects and loosely-typed key-value pairs. When
// processing pairs, the first element of the pair is used as the field key
// and the second as the field value.
//
// For example,
//   sugaredLogger.With(
//     "hello", "world",
//     "failure", errors.New("oh no"),
//     Stack(),
//     "count", 42,
//     "user", User{Name: "alice"},
//  )
// is the equivalent of
//   unsugared.With(
//     String("hello", "world"),
//     String("failure", "oh no"),
//     Stack(),
//     Int("count", 42),
//     Object("user", User{Name: "alice"}),
//   )
//
// Note that the keys in key-value pairs should be strings. In development,
// passing a non-string key panics. In production, the logger is more
// forgiving: a separate error is logged, but the key-value pair is skipped
// and execution continues. Passing an orphaned key triggers similar behavior:
// panics in development and errors in production.
func (s *SugaredLogger) With(args ...interface{}) *SugaredLogger {
	return &SugaredLogger{
		SugaredLogger: s.SugaredLogger.With(args...),
		l:             s.l,
	}
}

// Ctx returns a new sugared logger with the context.
func (s *SugaredLogger) Ctx(ctx context.Context) SugaredLoggerWithCtx {
	return SugaredLoggerWithCtx{
		ctx: ctx,
		s:   s,
	}
}

// Debugf uses fmt.Sprintf to log a templated message.
func (s *SugaredLogger) DebugfContext(ctx context.Context, template string, args ...interface{}) {
	s.logArgs(ctx, zap.DebugLevel, template, args)
	s.Debugf(template, args...)
}

// Infof uses fmt.Sprintf to log a templated message.
func (s *SugaredLogger) InfofContext(ctx context.Context, template string, args ...interface{}) {
	s.logArgs(ctx, zap.InfoLevel, template, args)
	s.Infof(template, args...)
}

// Warnf uses fmt.Sprintf to log a templated message.
func (s *SugaredLogger) WarnfContext(ctx context.Context, template string, args ...interface{}) {
	s.logArgs(ctx, zap.WarnLevel, template, args)
	s.Warnf(template, args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func (s *SugaredLogger) ErrorfContext(ctx context.Context, template string, args ...interface{}) {
	s.logArgs(ctx, zap.ErrorLevel, template, args)
	s.Errorf(template, args...)
}

// DPanicf uses fmt.Sprintf to log a templated message. In development, the
// logger then panics. (See DPanicLevel for details.)
func (s *SugaredLogger) DPanicfContext(ctx context.Context, template string, args ...interface{}) {
	s.logArgs(ctx, zap.DPanicLevel, template, args)
	s.DPanicf(template, args...)
}

// Panicf uses fmt.Sprintf to log a templated message, then panics.
func (s *SugaredLogger) PanicfContext(ctx context.Context, template string, args ...interface{}) {
	s.logArgs(ctx, zap.PanicLevel, template, args)
	s.Panicf(template, args...)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func (s *SugaredLogger) FatalfContext(ctx context.Context, template string, args ...interface{}) {
	s.logArgs(ctx, zap.FatalLevel, template, args)
	s.Fatalf(template, args...)
}

func (s *SugaredLogger) logArgs(
	ctx context.Context, lvl zapcore.Level, template string, args []interface{},
) {
	if lvl < s.l.minLevel {
		return
	}
	span := trace.SpanFromContext(ctx)
	if !span.IsRecording() {
		return
	}

	attrs := make([]attribute.KeyValue, 0, numAttr+1)
	attrs = append(attrs, logTemplateKey.String(template))

	s.l.log(span, lvl, fmt.Sprintf(template, args...), attrs)
}

// Infow logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (s *SugaredLogger) InfowContext(
	ctx context.Context, msg string, keysAndValues ...interface{},
) {
	keysAndValues = s.logKVs(ctx, zap.InfoLevel, msg, keysAndValues)
	s.Infow(msg, keysAndValues...)
}

// Warnw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (s *SugaredLogger) WarnwContext(
	ctx context.Context, msg string, keysAndValues ...interface{},
) {
	keysAndValues = s.logKVs(ctx, zap.WarnLevel, msg, keysAndValues)
	s.Warnw(msg, keysAndValues...)
}

// Errorw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (s *SugaredLogger) ErrorwContext(
	ctx context.Context, msg string, keysAndValues ...interface{},
) {
	keysAndValues = s.logKVs(ctx, zap.ErrorLevel, msg, keysAndValues)
	s.Errorw(msg, keysAndValues...)
}

// DPanicw logs a message with some additional context. In development, the
// logger then panics. (See DPanicLevel for details.) The variadic key-value
// pairs are treated as they are in With.
func (s *SugaredLogger) DPanicwContext(
	ctx context.Context, msg string, keysAndValues ...interface{},
) {
	keysAndValues = s.logKVs(ctx, zap.DPanicLevel, msg, keysAndValues)
	s.DPanicw(msg, keysAndValues...)
}

// Panicw logs a message with some additional context, then panics. The
// variadic key-value pairs are treated as they are in With.
func (s *SugaredLogger) PanicwContext(
	ctx context.Context, msg string, keysAndValues ...interface{},
) {
	keysAndValues = s.logKVs(ctx, zap.PanicLevel, msg, keysAndValues)
	s.Panicw(msg, keysAndValues...)
}

// Fatalw logs a message with some additional context, then calls os.Exit. The
// variadic key-value pairs are treated as they are in With.
func (s *SugaredLogger) FatalwContext(
	ctx context.Context, msg string, keysAndValues ...interface{},
) {
	keysAndValues = s.logKVs(ctx, zap.FatalLevel, msg, keysAndValues)
	s.Fatalw(msg, keysAndValues...)
}

func (s *SugaredLogger) logKVs(
	ctx context.Context, lvl zapcore.Level, msg string, kvs []interface{},
) []interface{} {
	if lvl < s.l.minLevel {
		return kvs
	}
	span := trace.SpanFromContext(ctx)
	if !span.IsRecording() {
		return kvs
	}

	attrs := make([]attribute.KeyValue, 0, numAttr+len(kvs))

	for i := 0; i < len(kvs); i += 2 {
		if key, ok := kvs[i].(string); ok {
			attrs = append(attrs, otelutil.Attribute(key, kvs[i+1]))
		}
	}

	s.l.log(span, lvl, msg, attrs)

	if s.l.withTraceID {
		traceID := span.SpanContext().TraceID().String()
		kvs = append(kvs, "trace_id", traceID)
	}

	return kvs
}

//------------------------------------------------------------------------------

type SugaredLoggerWithCtx struct {
	ctx context.Context
	s   *SugaredLogger
}

// Desugar unwraps a SugaredLogger, exposing the original Logger. Desugaring
// is quite inexpensive, so it's reasonable for a single application to use
// both Loggers and SugaredLoggers, converting between them on the boundaries
// of performance-sensitive code.
func (s SugaredLoggerWithCtx) Desugar() *Logger {
	return s.s.Desugar()
}

// Debugf uses fmt.Sprintf to log a templated message.
func (s SugaredLoggerWithCtx) Debugf(template string, args ...interface{}) {
	s.s.logArgs(s.ctx, zap.DebugLevel, template, args)
	s.s.skipCaller.Debugf(template, args...)
}

// Infof uses fmt.Sprintf to log a templated message.
func (s SugaredLoggerWithCtx) Infof(template string, args ...interface{}) {
	s.s.logArgs(s.ctx, zap.InfoLevel, template, args)
	s.s.skipCaller.Infof(template, args...)
}

// Warnf uses fmt.Sprintf to log a templated message.
func (s SugaredLoggerWithCtx) Warnf(template string, args ...interface{}) {
	s.s.logArgs(s.ctx, zap.WarnLevel, template, args)
	s.s.skipCaller.Warnf(template, args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func (s SugaredLoggerWithCtx) Errorf(template string, args ...interface{}) {
	s.s.logArgs(s.ctx, zap.ErrorLevel, template, args)
	s.s.skipCaller.Errorf(template, args...)
}

// DPanicf uses fmt.Sprintf to log a templated message. In development, the
// logger then panics. (See DPanicLevel for details.)
func (s SugaredLoggerWithCtx) DPanicf(template string, args ...interface{}) {
	s.s.logArgs(s.ctx, zap.DPanicLevel, template, args)
	s.s.skipCaller.DPanicf(template, args...)
}

// Panicf uses fmt.Sprintf to log a templated message, then panics.
func (s SugaredLoggerWithCtx) Panicf(template string, args ...interface{}) {
	s.s.logArgs(s.ctx, zap.PanicLevel, template, args)
	s.s.skipCaller.Panicf(template, args...)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func (s SugaredLoggerWithCtx) Fatalf(template string, args ...interface{}) {
	s.s.logArgs(s.ctx, zap.FatalLevel, template, args)
	s.s.skipCaller.Fatalf(template, args...)
}

// Debugw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
//
// When debug-level logging is disabled, this is much faster than
//  s.With(keysAndValues).Debug(msg)
func (s SugaredLoggerWithCtx) Debugw(msg string, keysAndValues ...interface{}) {
	s.s.logKVs(s.ctx, zap.DebugLevel, msg, keysAndValues)
	s.s.skipCaller.Debugw(msg, keysAndValues...)
}

// Infow logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (s SugaredLoggerWithCtx) Infow(msg string, keysAndValues ...interface{}) {
	s.s.logKVs(s.ctx, zap.InfoLevel, msg, keysAndValues)
	s.s.skipCaller.Infow(msg, keysAndValues...)
}

// Warnw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (s SugaredLoggerWithCtx) Warnw(msg string, keysAndValues ...interface{}) {
	s.s.logKVs(s.ctx, zap.WarnLevel, msg, keysAndValues)
	s.s.skipCaller.Warnw(msg, keysAndValues...)
}

// Errorw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (s SugaredLoggerWithCtx) Errorw(msg string, keysAndValues ...interface{}) {
	s.s.logKVs(s.ctx, zap.ErrorLevel, msg, keysAndValues)
	s.s.skipCaller.Errorw(msg, keysAndValues...)
}

// DPanicw logs a message with some additional context. In development, the
// logger then panics. (See DPanicLevel for details.) The variadic key-value
// pairs are treated as they are in With.
func (s SugaredLoggerWithCtx) DPanicw(msg string, keysAndValues ...interface{}) {
	s.s.logKVs(s.ctx, zap.DPanicLevel, msg, keysAndValues)
	s.s.skipCaller.DPanicw(msg, keysAndValues...)
}

// Panicw logs a message with some additional context, then panics. The
// variadic key-value pairs are treated as they are in With.
func (s SugaredLoggerWithCtx) Panicw(msg string, keysAndValues ...interface{}) {
	s.s.logKVs(s.ctx, zap.PanicLevel, msg, keysAndValues)
	s.s.skipCaller.Panicw(msg, keysAndValues...)
}

// Fatalw logs a message with some additional context, then calls os.Exit. The
// variadic key-value pairs are treated as they are in With.
func (s SugaredLoggerWithCtx) Fatalw(msg string, keysAndValues ...interface{}) {
	s.s.logKVs(s.ctx, zap.FatalLevel, msg, keysAndValues)
	s.s.skipCaller.Fatalw(msg, keysAndValues...)
}

//------------------------------------------------------------------------------

func appendField(attrs []attribute.KeyValue, f zapcore.Field) []attribute.KeyValue {
	switch f.Type {
	case zapcore.BoolType:
		attr := attribute.Bool(f.Key, f.Integer == 1)
		return append(attrs, attr)

	case zapcore.Int8Type, zapcore.Int16Type, zapcore.Int32Type, zapcore.Int64Type,
		zapcore.Uint32Type, zapcore.Uint8Type, zapcore.Uint16Type, zapcore.Uint64Type,
		zapcore.UintptrType:
		attr := attribute.Int64(f.Key, f.Integer)
		return append(attrs, attr)

	case zapcore.Float32Type, zapcore.Float64Type:
		attr := attribute.Float64(f.Key, math.Float64frombits(uint64(f.Integer)))
		return append(attrs, attr)

	case zapcore.Complex64Type:
		s := strconv.FormatComplex(complex128(f.Interface.(complex64)), 'E', -1, 64)
		attr := attribute.String(f.Key, s)
		return append(attrs, attr)
	case zapcore.Complex128Type:
		s := strconv.FormatComplex(f.Interface.(complex128), 'E', -1, 128)
		attr := attribute.String(f.Key, s)
		return append(attrs, attr)

	case zapcore.StringType:
		attr := attribute.String(f.Key, f.String)
		return append(attrs, attr)
	case zapcore.BinaryType, zapcore.ByteStringType:
		attr := attribute.String(f.Key, string(f.Interface.([]byte)))
		return append(attrs, attr)
	case zapcore.StringerType:
		attr := attribute.String(f.Key, f.Interface.(fmt.Stringer).String())
		return append(attrs, attr)

	case zapcore.DurationType, zapcore.TimeType:
		attr := attribute.Int64(f.Key, f.Integer)
		return append(attrs, attr)
	case zapcore.TimeFullType:
		attr := attribute.Int64(f.Key, f.Interface.(time.Time).UnixNano())
		return append(attrs, attr)
	case zapcore.ErrorType:
		err := f.Interface.(error)
		typ := reflect.TypeOf(err).String()
		attrs = append(attrs, semconv.ExceptionTypeKey.String(typ))
		attrs = append(attrs, semconv.ExceptionMessageKey.String(err.Error()))
		return attrs
	case zapcore.ReflectType:
		attr := otelutil.Attribute(f.Key, f.Interface)
		return append(attrs, attr)
	case zapcore.SkipType:
		return attrs

	case zapcore.ArrayMarshalerType:
		var attr attribute.KeyValue
		arrayEncoder := &bufferArrayEncoder{
			stringsSlice: []string{},
		}
		err := f.Interface.(zapcore.ArrayMarshaler).MarshalLogArray(arrayEncoder)
		if err != nil {
			attr = attribute.String(f.Key+"_error", fmt.Sprintf("otelzap: unable to marshal array: %v", err))
		} else {
			attr = attribute.StringSlice(f.Key, arrayEncoder.stringsSlice)
		}
		return append(attrs, attr)

	case zapcore.ObjectMarshalerType:
		attr := attribute.String(f.Key+"_error", "otelzap: zapcore.ObjectMarshalerType is not implemented")
		return append(attrs, attr)

	default:
		attr := attribute.String(f.Key+"_error", fmt.Sprintf("otelzap: unknown field type: %v", f))
		return append(attrs, attr)
	}
}

func levelString(lvl zapcore.Level) string {
	if lvl == zapcore.DPanicLevel {
		return "PANIC"
	}
	return lvl.CapitalString()
}
