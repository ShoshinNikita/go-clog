package log

import (
	"fmt"
	"os"
)

func (l Logger) Fatal(v ...interface{}) {
	text := ""
	if l.PrintTime {
		text = l.getTime()
	}
	text += l.getFatalMsg()
	if l.PrintErrorLine {
		text += getCaller()
	}
	l.printText(text + fmt.Sprint(v...))
	os.Exit(1)

}

func (l Logger) Fatalf(format string, v ...interface{}) {
	text := ""
	if l.PrintTime {
		text = l.getTime()
	}
	text += l.getFatalMsg()
	if l.PrintErrorLine {
		text += getCaller()
	}
	l.printText(text + fmt.Sprintf(format, v...))
	os.Exit(1)
}

func (l Logger) Fatalln(v ...interface{}) {
	text := ""
	if l.PrintTime {
		text = l.getTime()
	}
	text += l.getFatalMsg()
	if l.PrintErrorLine {
		text += getCaller()
	}
	l.printText(text + fmt.Sprint(v...))
	os.Exit(1)
}
