package log

import (
	"bytes"
	"fmt"
)

// Error prints error
// Output pattern: (?time) [ERR] (?file:line) error
func (l Logger) Error(v ...interface{}) {
	buf := &bytes.Buffer{}
	buf.Write(l.getTime())
	buf.Write(l.getErrMsg())
	buf.Write(l.getCaller())
	fmt.Fprint(buf, v...)

	l.mutex.Lock()
	l.output.Write(buf.Bytes())
	l.mutex.Unlock()
}

// Errorf prints error
// Output pattern: (?time) [ERR] (?file:line) error
func (l Logger) Errorf(format string, v ...interface{}) {
	buf := &bytes.Buffer{}
	buf.Write(l.getTime())
	buf.Write(l.getErrMsg())
	buf.Write(l.getCaller())
	fmt.Fprintf(buf, format, v...)

	l.mutex.Lock()
	l.output.Write(buf.Bytes())
	l.mutex.Unlock()
}

// Errorln prints error
// Output pattern: (?time) [ERR] (?file:line) error
func (l Logger) Errorln(v ...interface{}) {
	buf := &bytes.Buffer{}
	buf.Write(l.getTime())
	buf.Write(l.getErrMsg())
	buf.Write(l.getCaller())
	fmt.Fprintln(buf, v...)

	l.mutex.Lock()
	l.output.Write(buf.Bytes())
	l.mutex.Unlock()
}
