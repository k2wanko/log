// +build !appengine

package log

import (
	"os"

	"golang.org/x/net/context"
)

func New(ctx context.Context) Logger {
	return &IOLogger{
		Out: os.Stdout,
	}
}
