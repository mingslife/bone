package bone

import "context"

type Logger interface {
	Component

	WithContext(ctx context.Context) Logger
	Debug(msg any, extra ...map[string]any)
	Info(msg any, extra ...map[string]any)
	Warn(msg any, extra ...map[string]any)
	Error(msg any, extra ...map[string]any)
	Panic(msg any, extra ...map[string]any)
	Fatal(msg any, extra ...map[string]any)
}
