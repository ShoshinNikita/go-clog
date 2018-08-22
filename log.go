package log

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

var (
	showTime   bool
	printColor = true
	timeLayout = "01.02.2006 15:04:05"

	printChan = make(chan string, 500)

	// For [ERR]
	red = color.New(color.FgRed).Sprint

	// For [INFO]
	cyan = color.New(color.FgCyan).Sprint

	// For time
	yellowf = color.New(color.FgYellow).Sprintf

	// For fatal
	bgRed = color.New(color.BgRed).Sprint
)

func init() {
	go func() {
		for text := range printChan {
			fmt.Fprint(color.Output, text)
		}
	}()
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
