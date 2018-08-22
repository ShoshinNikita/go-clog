package log

import (
	"fmt"
	"os"
)

func Fatal(v ...interface{}) {
	text := ""
	if showTime {
		text = getTime()
	}
	text += getFatalMsg()
	printText(text + fmt.Sprint(v...))
	os.Exit(1)

}

func Fatalf(format string, v ...interface{}) {
	text := ""
	if showTime {
		text = getTime()
	}
	text += getFatalMsg()
	printText(text + fmt.Sprintf(format, v...))
	os.Exit(1)
}

func Fatalln(v ...interface{}) {
	text := ""
	if showTime {
		text = getTime()
	}
	text += getFatalMsg()
	printText(text + fmt.Sprint(v...))
	os.Exit(1)
}
