package log

import (
	"fmt"
)

func Error(v ...interface{}) {
	text := ""
	if showTime {
		text = getTime()
	}
	text += getErrMsg()
	printChan <- text + fmt.Sprint(v...)
}

func Errorf(format string, v ...interface{}) {
	text := ""
	if showTime {
		text = getTime()
	}
	text += getErrMsg()
	printChan <- text + fmt.Sprintf(format, v...)
}

func Errorln(v ...interface{}) {
	text := ""
	if showTime {
		text = getTime()
	}
	text += getErrMsg()
	printChan <- text + fmt.Sprintln(v...)
}
