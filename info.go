package log

import (
	"bytes"
	"fmt"
)

// Info prints info message
// Output pattern: (?time) [INFO] msg
func (l Logger) Info(v ...interface{}) {
	buf := &bytes.Buffer{}
	buf.Write(l.getTime())
	buf.Write(l.getInfoMsg())
	fmt.Fprint(buf, v...)

	l.mutex.Lock()
	l.output.Write(buf.Bytes())
	l.mutex.Unlock()
}

// Infof prints info message
// Output pattern: (?time) [INFO] msg
func (l Logger) Infof(format string, v ...interface{}) {
	buf := &bytes.Buffer{}
	buf.Write(l.getTime())
	buf.Write(l.getInfoMsg())
	fmt.Fprintf(buf, format, v...)

	l.mutex.Lock()
	l.output.Write(buf.Bytes())
	l.mutex.Unlock()
}

// Infoln prints info message
// Output pattern: (?time) [INFO] msg
func (l Logger) Infoln(v ...interface{}) {
	buf := &bytes.Buffer{}
	buf.Write(l.getTime())
	buf.Write(l.getInfoMsg())
	fmt.Fprintln(buf, v...)

	l.mutex.Lock()
	l.output.Write(buf.Bytes())
	l.mutex.Unlock()
}
