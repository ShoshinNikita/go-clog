// Package log provides functions for pretty print
//
// Patterns of functions print:
// * Print(), Printf(), Println():
//   (?time) msg
// * Info(), Infof(), Infoln():
//   (?time) [INF] msg
// * Warn(), Warnf(), Warnln():
//   (?time) [WRN] warning
// * Error(), Errorf(), Errorln():
//   (?time) [ERR] (?file:line) error
// * Fatal(), Fatalf(), Fatalln():
//   (?time) [FAT] (?file:line) error
//
// Time pattern: MM.dd.yyyy hh:mm:ss (01.30.2018 05:5:59)
//
package log

import (
	"io"
	"sync"

	"github.com/fatih/color"
)

const (
	DefaultTimeLayout = "01.02.2006 15:04:05"
)

type Logger struct {
	output io.Writer
	mutex  *sync.Mutex

	global bool

	printTime      bool
	printColor     bool
	printErrorLine bool
	timeLayout     string
}

// NewLogger creates *Logger
func NewLogger() *Logger {
	l := new(Logger)
	l.output = color.Output
	l.timeLayout = DefaultTimeLayout
	l.mutex = new(sync.Mutex)
	return l
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

// ChangeOutput changes Logger.output writer.
// Default Logger.output is github.com/fatih/color.Output
func (l *Logger) ChangeOutput(w io.Writer) {
	l.output = w
}

// ChangeTimeLayout changes Logger.timeLayout
// Default Logger.timeLayout is DefaultTimeLayout
func (l *Logger) ChangeTimeLayout(layout string) {
	l.timeLayout = layout
}
