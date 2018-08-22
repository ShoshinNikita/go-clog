package log

import (
	"fmt"
)

func Print(v ...interface{}) {
	printChan <- fmt.Sprint(v...)
}

func Printf(format string, v ...interface{}) {
	printChan <- fmt.Sprintf(format, v...)
}

func Println(v ...interface{}) {
	printChan <- fmt.Sprintln(v...)
}
