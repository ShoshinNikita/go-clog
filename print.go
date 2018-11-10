package log

import (
	"fmt"
)

func (l Logger) Print(v ...interface{}) {
	l.printText(fmt.Sprint(v...))
}

func (l Logger) Printf(format string, v ...interface{}) {
	l.printText(fmt.Sprintf(format, v...))
}

func (l Logger) Println(v ...interface{}) {
	l.printText(fmt.Sprintln(v...))
}
