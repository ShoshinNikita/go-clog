package clog

import (
	"fmt"
	"time"
)

// Debug prints debug message if Debug mode is on
// Output pattern: (?time) [DBG] msg
func (l Logger) Debug(v ...interface{}) {
	print := func() (int, error) {
		return fmt.Fprintln(l.buff, v...)
	}

	l.debug(print)
}

// Debugf prints debug message if Debug mode is on
// Output pattern: (?time) [DBG] msg
func (l Logger) Debugf(format string, v ...interface{}) {
	print := func() (int, error) {
		return fmt.Fprintf(l.buff, format, v...)
	}

	l.debug(print)
}

// debug is an internal function for printing debug messages
// Output pattern: (?time) [DBG] msg
func (l Logger) debug(print messagePrintFunction) {
	if !l.shouldPrint(LevelDebug) {
		return
	}

	now := time.Now()

	l.mutex.Lock()
	defer l.mutex.Unlock()

	l.buff.Reset()

	l.writeIntoBuffer(l.getTime(now))
	l.writeIntoBuffer(l.getDebugPrefix())

	print()

	l.output.Write(l.buff.Bytes())
}
