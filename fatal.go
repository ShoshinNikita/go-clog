package log

import (
	"fmt"
	"os"
)

func Fatal(v ...interface{}) {
	if showTime {
		printTime()
	}
	printFatalMsg()
	fmt.Print(v...)
	os.Exit(1)

}

func Fatalf(format string, v ...interface{}) {
	if showTime {
		printTime()
	}
	printFatalMsg()
	fmt.Printf(format, v...)
	os.Exit(1)
}

func Fatalln(v ...interface{}) {
	if showTime {
		printTime()
	}
	printFatalMsg()
	fmt.Println(v...)
	os.Exit(1)
}
