package log

import (
	"bytes"
	"fmt"
)

// Print prints msg
// Output pattern: (?time) msg
func (l Logger) Print(v ...interface{}) {
	buf := &bytes.Buffer{}
	buf.Write(l.getTime())
	fmt.Fprint(buf, v...)

	l.mutex.Lock()
	l.output.Write(buf.Bytes())
	l.mutex.Unlock()
}

// Printf prints msg
// Output pattern: (?time) msg
func (l Logger) Printf(format string, v ...interface{}) {
	buf := &bytes.Buffer{}
	buf.Write(l.getTime())
	fmt.Fprintf(buf, format, v...)

	l.mutex.Lock()
	l.output.Write(buf.Bytes())
	l.mutex.Unlock()
}

// Println prints msg
// Output pattern: (?time) msg
func (l Logger) Println(v ...interface{}) {
	buf := &bytes.Buffer{}
	buf.Write(l.getTime())
	fmt.Fprintln(buf, v...)

	l.mutex.Lock()
	l.output.Write(buf.Bytes())
	l.mutex.Unlock()
}
