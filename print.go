package log

import (
	"bytes"
	"fmt"
	"time"
)

// Print prints msg
// Output pattern: (?time) msg
func (l Logger) Print(v ...interface{}) {
	now := time.Now()

	buf := &bytes.Buffer{}
	buf.Write(l.getTime(now))
	fmt.Fprint(buf, v...)

	l.mutex.Lock()
	l.output.Write(buf.Bytes())
	l.mutex.Unlock()
}

// Printf prints msg
// Output pattern: (?time) msg
func (l Logger) Printf(format string, v ...interface{}) {
	now := time.Now()

	buf := &bytes.Buffer{}
	buf.Write(l.getTime(now))
	fmt.Fprintf(buf, format, v...)

	l.mutex.Lock()
	l.output.Write(buf.Bytes())
	l.mutex.Unlock()
}

// Println prints msg
// Output pattern: (?time) msg
func (l Logger) Println(v ...interface{}) {
	now := time.Now()

	buf := &bytes.Buffer{}
	buf.Write(l.getTime(now))
	fmt.Fprintln(buf, v...)

	l.mutex.Lock()
	l.output.Write(buf.Bytes())
	l.mutex.Unlock()
}
