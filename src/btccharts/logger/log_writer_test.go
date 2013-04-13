package logger

import "testing"

func TestNewLogWriter(t *testing.T) {
	lw := NewLogWriter()
	if lw.path != "./data/test.log" {
		t.Errorf("Fail %#v", lw)
	}
}
