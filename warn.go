package clog

import (
	"fmt"
	"time"
)

// Warn prints warning
// Output pattern: (?time) [WRN] warning
func (l Logger) Warn(v ...interface{}) {
	print := func() (int, error) {
		return fmt.Fprintln(l.buff, v...)
	}

	l.warn(print)
}

// Warnf prints warning
// Output pattern: (?time) [WRN] warning
func (l Logger) Warnf(format string, v ...interface{}) {
	print := func() (int, error) {
		return fmt.Fprintf(l.buff, format, v...)
	}

	l.warn(print)
}

// warn is an internal function for printing warning messages
// Output pattern: (?time) [WRN] warning
func (l Logger) warn(print messagePrintFunction) {
	if !l.shouldPrint(LevelWarn) {
		return
	}

	now := time.Now()

	l.mutex.Lock()
	defer l.mutex.Unlock()

	l.buff.Reset()

	l.buff.Write(l.getTime(now))
	l.buff.Write(l.getWarnMsg())

	print()

	l.output.Write(l.buff.Bytes())
}
