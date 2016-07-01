package log

import "golang.org/x/net/context"

var ctxKey = "log"

type Logger interface {
	Debugf(string, ...interface{})
	Infof(string, ...interface{})
	Warningf(string, ...interface{})
	Errorf(string, ...interface{})
	Criticalf(string, ...interface{})
}

func WithContext(ctx context.Context, l Logger) context.Context {
	return context.WithValue(ctx, &ctxKey, l)
}

func FromContext(ctx context.Context) Logger {
	if l, ok := ctx.Value(&ctxKey).(Logger); ok {
		return l
	}
	return nil
}
