package log

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

type textStruct struct {
	text string
	ch   chan struct{}
}

func (t *textStruct) done() {
	close(t.ch)
}

func newText(text string) textStruct {
	return textStruct{text: text, ch: make(chan struct{})}
}

const (
	timeLayout = "01.02.2006 15:04:05"
)

var (
	showTime   bool
	printColor = true
	printChan  = make(chan textStruct, 500)

	// For [ERR]
	red = color.New(color.FgRed).SprintFunc()

	// For [INFO]
	cyan = color.New(color.FgCyan).SprintFunc()

	// For time
	yellowf = color.New(color.FgYellow).SprintfFunc()

	// For fatal
	bgRed = color.New(color.BgRed).SprintFunc()
)

// init runs goroutine, which prints text from channel
func init() {
	go func() {
		for text := range printChan {
			fmt.Fprint(color.Output, text.text)
			text.done()
		}
	}()
}

func printText(text string) {
	t := newText(text)
	printChan <- t
	<-t.ch
}

// ShowTime sets showTime
// Time isn't printed by default
func ShowTime(b bool) {
	showTime = b
}

// PrintColor sets printColor
// printColor is true by default
func PrintColor(b bool) {
	printColor = b
}

func getTime() string {
	if printColor {
		return yellowf("%s ", time.Now().Format(timeLayout))
	}
	return fmt.Sprintf("%s ", time.Now().Format(timeLayout))
}

func getErrMsg() string {
	if printColor {
		return red("[ERR] ")
	}
	return "[ERR] "
}

func getInfoMsg() string {
	if printColor {
		return cyan("[INFO] ")
	}
	return "[INFO] "
}

func getFatalMsg() (s string) {
	if printColor {
		return bgRed("[FATAL]") + " "
	}
	return "[FATAL] "
}
