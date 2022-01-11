package otelzap

import "go.uber.org/zap/zapcore"

// Option applies a configuration to the given config.
type Option func(l *Logger)

// WithMinLevel sets the minimal zap logging level on which the log message
// is recorded on the span.
//
// The default is >= zap.WarnLevel.
func WithMinLevel(lvl zapcore.Level) Option {
	return func(l *Logger) {
		l.minLevel = lvl
	}
}

// WithErrorStatusLevel sets the minimal zap logging level on which
// the span status is set to codes.Error.
//
// The default is >= zap.ErrorLevel.
func WithErrorStatusLevel(lvl zapcore.Level) Option {
	return func(l *Logger) {
		l.errorStatusLevel = lvl
	}
}

// WithCaller configures the logger to annotate each event with the filename,
// line number, and function name of the caller.
//
// It is enabled by default.
func WithCaller(on bool) Option {
	return func(l *Logger) {
		l.caller = on
	}
}

// WithStackTrace configures the logger to capture logs with a stack trace.
func WithStackTrace(on bool) Option {
	return func(l *Logger) {
		l.stackTrace = on
	}
}

// WithTraceIDField configures the logger to add `trace_id` field to structured log messages.
//
// This option is only useful with backends that don't support OTLP and instead parse log
// messages to extract structured information.
func WithTraceIDField(on bool) Option {
	return func(l *Logger) {
		l.withTraceID = on
	}
}
