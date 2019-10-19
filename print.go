package clog

import (
	"fmt"
	"time"
)

// Print prints msg
// Output pattern: (?time) msg
func (l Logger) Print(v ...interface{}) {
	print := func() (int, error) {
		return fmt.Fprintln(l.buff, v...)
	}

	l.print(print)
}

// Printf prints msg
// Output pattern: (?time) msg
func (l Logger) Printf(format string, v ...interface{}) {
	print := func() (int, error) {
		return fmt.Fprintf(l.buff, format, v...)
	}

	l.print(print)
}

func (l Logger) print(print messagePrintFunction) {
	now := time.Now()

	l.mutex.Lock()
	defer l.mutex.Unlock()

	l.buff.Reset()

	l.buff.Write(l.getTime(now))

	print()

	l.output.Write(l.buff.Bytes())
}
