package log

import (
	"google.golang.org/appengine/log"

	"golang.org/x/net/context"
)

type AELogger struct {
	c context.Context
}

func (l *AELogger) Debugf(format string, args ...interface{}) {
	log.Debugf(l.c, format, args...)
}

func (l *AELogger) Infof(format string, args ...interface{}) {
	log.Infof(l.c, format, args...)
}

func (l *AELogger) Warningf(format string, args ...interface{}) {
	log.Warningf(l.c, format, args...)
}

func (l *AELogger) Errorf(format string, args ...interface{}) {
	log.Errorf(l.c, format, args...)
}

func (l *AELogger) Criticalf(format string, args ...interface{}) {
	log.Criticalf(l.c, format, args...)
}
