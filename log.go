package log

import "context"

type Log interface {
	Debug(ctx context.Context, msg string)
	Trace(ctx context.Context, msg string)
	Notice(ctx context.Context, msg string)
	Warning(ctx context.Context, msg string)
	Error(ctx context.Context, msg string)
	Fatal(ctx context.Context, msg string)
}
