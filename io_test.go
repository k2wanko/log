package log

import (
	"bytes"
	"testing"
)

func TestIOLogger(t *testing.T) {
	out := new(bytes.Buffer)
	l := &IOLogger{Out: out}

	l.Debugf("Test %s", "Word")

	if o, want := out.String(), "[Debug] Test Word"; o != want {
		t.Errorf("out = %s; want %s", o, want)
	}
}
