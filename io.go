package log

import (
	"fmt"
	"io"
)

type IOLogger struct {
	Out io.Writer
}

func (l *IOLogger) Debugf(format string, args ...interface{}) {
	fmt.Fprintf(l.Out, "[Debug] "+format, args...)
}

func (l *IOLogger) Infof(format string, args ...interface{}) {
	fmt.Fprintf(l.Out, "[Info] "+format, args...)
}

func (l *IOLogger) Warningf(format string, args ...interface{}) {
	fmt.Fprintf(l.Out, "[Warning] "+format, args...)
}

func (l *IOLogger) Errorf(format string, args ...interface{}) {
	fmt.Fprintf(l.Out, "[Error] "+format, args...)
}

func (l *IOLogger) Criticalf(format string, args ...interface{}) {
	fmt.Fprintf(l.Out, "[Critical] "+format, args...)
}
