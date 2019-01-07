// Package log provides functions for pretty print
//
// Patterns of functions print:
// * Print(), Printf(), Println():
//   (?time) msg
// * Info(), Infof(), Infoln():
//   (?time) [INFO] msg
// * Warn(), Warnf(), Warnln():
//   (?time) [WARN] warning
// * Error(), Errorf(), Errorln():
//   (?time) [ERR] (?file:line) error
// * Fatal(), Fatalf(), Fatalln():
//   (?time) [FATAL] (?file:line) error
//
// Time pattern: MM.dd.yyyy hh:mm:ss (01.30.2018 05:5:59)
//
package log

import (
	"fmt"

	"github.com/fatih/color"
)

const (
	timeLayout = "01.02.2006 15:04:05"
)

// init inits globalLogger with NewLogger()
func init() {
	globalLogger = NewLogger()

	globalLogger.PrintTime(false)
	globalLogger.PrintColor(false)
	globalLogger.PrintErrorLine(false)

	globalLogger.global = true
}

type textStruct struct {
	text string
	ch   chan struct{}
}

func newText(text string) textStruct {
	return textStruct{text: text, ch: make(chan struct{})}
}

func (t *textStruct) done() {
	close(t.ch)
}

type Logger struct {
	printTime      bool
	printColor     bool
	printErrorLine bool

	printChan chan textStruct
	global    bool
}

// NewLogger creates *Logger and run goroutine (Logger.printer())
func NewLogger() *Logger {
	l := new(Logger)
	l.printChan = make(chan textStruct, 200)
	go l.printer()
	return l
}

func (l Logger) printer() {
	for text := range l.printChan {
		fmt.Fprint(color.Output, text.text)
		text.done()
	}
}

func (l Logger) printText(text string) {
	t := newText(text)
	l.printChan <- t
	<-t.ch
}

// PrintTime sets Logger.printTime to b
func (l *Logger) PrintTime(b bool) {
	l.printTime = b
}

// PrintColor sets Logger.printColor to b
func (l *Logger) PrintColor(b bool) {
	l.printColor = b
}

// PrintErrorLine sets Logger.printErrorLine to b
func (l *Logger) PrintErrorLine(b bool) {
	l.printErrorLine = b
}
