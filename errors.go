package log

import (
	"fmt"
)

// Error prints error
// Output pattern: (?time) [ERR] (?file:line) error
func (l Logger) Error(v ...interface{}) {
	text := ""
	text = l.getTime()

	text += l.getErrMsg()
	text += l.getCaller()

	l.printText(text + fmt.Sprint(v...))
}

// Errorf prints error
// Output pattern: (?time) [ERR] (?file:line) error
func (l Logger) Errorf(format string, v ...interface{}) {
	text := ""
	text = l.getTime()

	text += l.getErrMsg()
	text += l.getCaller()

	l.printText(text + fmt.Sprintf(format, v...))
}

// Errorln prints error
// Output pattern: (?time) [ERR] (?file:line) error
func (l Logger) Errorln(v ...interface{}) {
	text := ""
	text = l.getTime()

	text += l.getErrMsg()
	text += l.getCaller()

	l.printText(text + fmt.Sprintln(v...))
}
