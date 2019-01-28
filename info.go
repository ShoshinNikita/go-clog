package log

import (
	"bytes"
	"fmt"
	"time"
)

// Info prints info message
// Output pattern: (?time) [INF] msg
func (l Logger) Info(v ...interface{}) {
	now := time.Now()

	buf := &bytes.Buffer{}
	buf.Write(l.getTime(now))
	buf.Write(l.getInfoMsg())
	fmt.Fprint(buf, v...)

	l.mutex.Lock()
	l.output.Write(buf.Bytes())
	l.mutex.Unlock()
}

// Infof prints info message
// Output pattern: (?time) [INF] msg
func (l Logger) Infof(format string, v ...interface{}) {
	now := time.Now()

	buf := &bytes.Buffer{}
	buf.Write(l.getTime(now))
	buf.Write(l.getInfoMsg())
	fmt.Fprintf(buf, format, v...)

	l.mutex.Lock()
	l.output.Write(buf.Bytes())
	l.mutex.Unlock()
}

// Infoln prints info message
// Output pattern: (?time) [INF] msg
func (l Logger) Infoln(v ...interface{}) {
	now := time.Now()

	buf := &bytes.Buffer{}
	buf.Write(l.getTime(now))
	buf.Write(l.getInfoMsg())
	fmt.Fprintln(buf, v...)

	l.mutex.Lock()
	l.output.Write(buf.Bytes())
	l.mutex.Unlock()
}
