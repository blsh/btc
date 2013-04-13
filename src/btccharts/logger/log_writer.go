package logger

import (
	"fmt"
	"io"
	"os"
)

type LogWriter struct {
	path string
	c    chan string
}

func (lw LogWriter) Write(msg string) {
	lw.c <- msg
}

func NewLogWriter() LogWriter {
	lw := LogWriter{path: "./data/test.log", c: make(chan string)}
	go lw.run(os.Stdout)
	return lw
}

func (lw LogWriter) run(out io.Writer) {
	for line := range lw.c {
		fmt.Fprint(out, line)
	}
}
