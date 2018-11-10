package log

import (
	"fmt"
)

func (l Logger) Info(v ...interface{}) {
	text := ""
	if l.PrintTime {
		text = l.getTime()
	}
	text += l.getInfoMsg()
	l.printText(text + fmt.Sprint(v...))
}

func (l Logger) Infof(format string, v ...interface{}) {
	text := ""
	if l.PrintTime {
		text = l.getTime()
	}
	text += l.getInfoMsg()
	l.printText(text + fmt.Sprintf(format, v...))
}

func (l Logger) Infoln(v ...interface{}) {
	text := ""
	if l.PrintTime {
		text = l.getTime()
	}
	text += l.getInfoMsg()
	l.printText(text + fmt.Sprintln(v...))
}
