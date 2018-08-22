package log

import (
	"fmt"
)

func Info(v ...interface{}) {
	text := ""
	if showTime {
		text = getTime()
	}
	text += getInfoMsg()
	printChan <- text + fmt.Sprint(v...)
}

func Infof(format string, v ...interface{}) {
	text := ""
	if showTime {
		text = getTime()
	}
	text += getInfoMsg()
	printChan <- text + fmt.Sprintf(format, v...)
}

func Infoln(v ...interface{}) {
	text := ""
	if showTime {
		text = getTime()
	}
	text += getInfoMsg()
	printChan <- text + fmt.Sprintln(v...)
}
