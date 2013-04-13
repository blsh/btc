package logger

import (
	"io"
)

type LogWriter struct {
	out io.Writer
	c   chan string
}

func (lw LogWriter) Write(msg string) {
	lw.c <- msg
}

func NewLogWriter(out io.Writer) LogWriter {
	lw := LogWriter{out, make(chan string)}
	go lw.run()
	return lw
}

func (lw LogWriter) run() {
	for line := range lw.c {
		io.WriteString(lw.out, line)
	}
}
