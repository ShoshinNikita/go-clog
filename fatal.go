package log

import (
	"bytes"
	"fmt"
	"os"
	"time"
)

// Fatal prints error and call os.Exit(1)
// Output pattern: (?time) [FAT] (?file:line) error
func (l Logger) Fatal(v ...interface{}) {
	now := time.Now()

	buf := &bytes.Buffer{}
	buf.Write(l.getTime(now))
	buf.Write(l.getFatalMsg())
	buf.Write(l.getCaller())
	fmt.Fprint(buf, v...)

	l.mutex.Lock()
	l.output.Write(buf.Bytes())
	l.mutex.Unlock()

	os.Exit(1)
}

// Fatalf prints error and call os.Exit(1)
// Output pattern: (?time) [FAT] (?file:line) error
func (l Logger) Fatalf(format string, v ...interface{}) {
	now := time.Now()

	buf := &bytes.Buffer{}
	buf.Write(l.getTime(now))
	buf.Write(l.getFatalMsg())
	buf.Write(l.getCaller())
	fmt.Fprintf(buf, format, v...)

	l.mutex.Lock()
	l.output.Write(buf.Bytes())
	l.mutex.Unlock()

	os.Exit(1)
}

// Fatalln prints error and call os.Exit(1)
// Output pattern: (?time) [FAT] (?file:line) error
func (l Logger) Fatalln(v ...interface{}) {
	now := time.Now()

	buf := &bytes.Buffer{}
	buf.Write(l.getTime(now))
	buf.Write(l.getFatalMsg())
	buf.Write(l.getCaller())
	fmt.Fprintln(buf, v...)

	l.mutex.Lock()
	l.output.Write(buf.Bytes())
	l.mutex.Unlock()

	os.Exit(1)
}
