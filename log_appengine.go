// +build appengine

package log

import "golang.org/x/net/context"

func New(ctx context.Context) Logger {
	return &AELogger{c: ctx}
}
