package log

import (
	"fmt"
)

func LogInfo(v ...interface{}) {
	if showTime {
		printTime()
	}
	printInfoMsg()
	fmt.Print(v...)
}

func LogInfof(format string, v ...interface{}) {
	if showTime {
		printTime()
	}
	printInfoMsg()
	fmt.Printf(format, v...)
}

func LogInfoln(v ...interface{}) {
	if showTime {
		printTime()
	}
	printInfoMsg()
	fmt.Println(v...)
}
