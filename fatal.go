package log

import (
	"bytes"
	"fmt"
	"os"
)

// Fatal prints error and call os.Exit(1)
// Output pattern: (?time) [FATAL] (?file:line) error
func (l Logger) Fatal(v ...interface{}) {
	buf := &bytes.Buffer{}
	buf.Write(l.getTime())
	buf.Write(l.getFatalMsg())
	buf.Write(l.getCaller())
	fmt.Fprint(buf, v...)

	l.mutex.Lock()
	l.output.Write(buf.Bytes())
	l.mutex.Unlock()

	os.Exit(1)
}

// Fatalf prints error and call os.Exit(1)
// Output pattern: (?time) [FATAL] (?file:line) error
func (l Logger) Fatalf(format string, v ...interface{}) {
	buf := &bytes.Buffer{}
	buf.Write(l.getTime())
	buf.Write(l.getFatalMsg())
	buf.Write(l.getCaller())
	fmt.Fprintf(buf, format, v...)

	l.mutex.Lock()
	l.output.Write(buf.Bytes())
	l.mutex.Unlock()

	os.Exit(1)
}

// Fatalln prints error and call os.Exit(1)
// Output pattern: (?time) [FATAL] (?file:line) error
func (l Logger) Fatalln(v ...interface{}) {
	buf := &bytes.Buffer{}
	buf.Write(l.getTime())
	buf.Write(l.getFatalMsg())
	buf.Write(l.getCaller())
	fmt.Fprintln(buf, v...)

	l.mutex.Lock()
	l.output.Write(buf.Bytes())
	l.mutex.Unlock()

	os.Exit(1)
}
