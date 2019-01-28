package log

import (
	"bytes"
	"fmt"
)

// Warn prints warning
// Output pattern: (?time) [WARN] warning
func (l Logger) Warn(v ...interface{}) {
	buf := &bytes.Buffer{}
	buf.Write(l.getTime())
	buf.Write(l.getWarnMsg())
	fmt.Fprint(buf, v...)

	l.mutex.Lock()
	l.output.Write(buf.Bytes())
	l.mutex.Unlock()
}

// Warnf prints warning
// Output pattern: (?time) [WARN] warning
func (l Logger) Warnf(format string, v ...interface{}) {
	buf := &bytes.Buffer{}
	buf.Write(l.getTime())
	buf.Write(l.getWarnMsg())
	fmt.Fprintf(buf, format, v...)

	l.mutex.Lock()
	l.output.Write(buf.Bytes())
	l.mutex.Unlock()
}

// Warnln prints warning
// Output pattern: (?time) [WARN] warning
func (l Logger) Warnln(v ...interface{}) {
	buf := &bytes.Buffer{}
	buf.Write(l.getTime())
	buf.Write(l.getWarnMsg())
	fmt.Fprintln(buf, v...)

	l.mutex.Lock()
	l.output.Write(buf.Bytes())
	l.mutex.Unlock()
}
