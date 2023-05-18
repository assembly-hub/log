package log

import "context"

type Log interface {
	Debug(ctx context.Context, msg ...any)
	Trace(ctx context.Context, msg ...any)
	Notice(ctx context.Context, msg ...any)
	Warning(ctx context.Context, msg ...any)
	Error(ctx context.Context, msg ...any)
	Fatal(ctx context.Context, msg ...any)
}
