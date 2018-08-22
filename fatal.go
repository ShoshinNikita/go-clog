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
	text += fmt.Sprint(v...)
	// Don't send text into channel, because we can exit before printing of the text
	fmt.Print(v...)
	os.Exit(1)

}

func Fatalf(format string, v ...interface{}) {
	text := ""
	if showTime {
		text = getTime()
	}
	text += getFatalMsg()
	text += fmt.Sprintf(format, v...)
	// Don't send text into channel, because we can exit before printing of the text
	fmt.Print(text)
	os.Exit(1)
}

func Fatalln(v ...interface{}) {
	text := ""
	if showTime {
		text = getTime()
	}
	text += getFatalMsg()
	text += fmt.Sprint(v...)
	// Don't send text into channel, because we can exit before printing of the text
	fmt.Println(text)
	os.Exit(1)
}
