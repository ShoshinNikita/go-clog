package clog

import (
	"fmt"
	"time"
)

// Info prints info message
// Output pattern: (?time) [INF] msg
func (l Logger) Info(v ...interface{}) {
	print := func() (int, error) {
		return fmt.Fprintln(l.buff, v...)
	}

	l.info(print)
}

// Infof prints info message
// Output pattern: (?time) [INF] msg
func (l Logger) Infof(format string, v ...interface{}) {
	print := func() (int, error) {
		return fmt.Fprintf(l.buff, format, v...)
	}

	l.info(print)
}

// Now is an internal function for printing info messages
// Output pattern: (?time) [INF] msg
func (l Logger) info(print messagePrintFunction) {
	if !l.shouldPrint(LevelInfo) {
		return
	}

	now := time.Now()

	l.mutex.Lock()
	defer l.mutex.Unlock()

	l.buff.Reset()

	l.buff.Write(l.getTime(now))
	l.buff.Write(l.getInfoMsg())
	print()

	l.output.Write(l.buff.Bytes())
}
