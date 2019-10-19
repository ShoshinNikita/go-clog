package clog

import (
	"fmt"
	"os"
	"time"
)

// Fatal prints error and call os.Exit(1)
// Output pattern: (?time) [FAT] (?file:line) error
func (l Logger) Fatal(v ...interface{}) {
	print := func() (int, error) {
		return fmt.Fprintln(l.buff, v...)
	}

	l.fatal(print)
}

// Fatalf prints error and call os.Exit(1)
// Output pattern: (?time) [FAT] (?file:line) error
func (l Logger) Fatalf(format string, v ...interface{}) {
	print := func() (int, error) {
		return fmt.Fprintf(l.buff, format, v...)
	}

	l.fatal(print)
}

// fatal is an internal function for printing fatal messages. Is also calls os.Exit(1)
// Output pattern: (?time) [FAT] (?file:line) error
func (l Logger) fatal(print messagePrintFunction) {
	now := time.Now()

	l.mutex.Lock()
	defer l.mutex.Unlock()

	l.buff.Reset()

	l.buff.Write(l.getTime(now))
	l.buff.Write(l.getFatalMsg())
	l.buff.Write(l.getCaller())

	print()

	l.output.Write(l.buff.Bytes())

	os.Exit(1)
}
