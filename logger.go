package bone

import "context"

// Logger defines the common log methods.
type Logger interface {
	Component

	// WithContext returns Logger, which has its own context, convenient for
	// tracking.
	WithContext(ctx context.Context) Logger
	// Debug logs to the DEBUG log.
	Debug(msg any, extra ...map[string]any)
	// Info logs to the INFO log.
	Info(msg any, extra ...map[string]any)
	// Warn logs to the WARN log.
	Warn(msg any, extra ...map[string]any)
	// Error logs to the Error log.
	Error(msg any, extra ...map[string]any)
	// Panic logs to the PANIC log. And application should panic, which is
	// recoverable, after calling this method.
	Panic(msg any, extra ...map[string]any)
	// Fatal logs to the FATAL log. And application should fatal, which is NOT
	// recoverable, after calling this method.
	Fatal(msg any, extra ...map[string]any)
}
