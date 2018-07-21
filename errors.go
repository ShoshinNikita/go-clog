package log

import (
	"fmt"
)

func LogError(v ...interface{}) {
	if showTime {
		printTime()
	}
	printErrMsg()
	fmt.Print(v...)
}

func LogErrorf(format string, v ...interface{}) {
	if showTime {
		printTime()
	}
	printErrMsg()
	fmt.Printf(format, v...)
}

func LogErrorln(v ...interface{}) {
	if showTime {
		printTime()
	}
	printErrMsg()
	fmt.Println(v...)
}
