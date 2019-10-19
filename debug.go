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

	l.debugPrint(print)
}

// Debugf prints debug message if Debug mode is on
// Output pattern: (?time) [DBG] msg
func (l Logger) Debugf(format string, v ...interface{}) {
	print := func() (int, error) {
		return fmt.Fprintf(l.buff, format, v...)
	}

	l.debugPrint(print)
}

// TODO: update name
// debugPrint is an internal function for printing debug messages
// Output pattern: (?time) [DBG] msg
func (l Logger) debugPrint(print messagePrintFunction) {
	if !l.debug {
		return
	}

	now := time.Now()

	l.mutex.Lock()
	defer l.mutex.Unlock()

	l.buff.Reset()

	l.buff.Write(l.getTime(now))
	l.buff.Write(l.getDebugMsg())

	print()

	l.output.Write(l.buff.Bytes())
}
