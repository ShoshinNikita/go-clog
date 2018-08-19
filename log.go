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

	// For [ERR]
	red = color.New(color.FgRed).PrintFunc()

	// For [INFO]
	cyan = color.New(color.FgCyan).PrintFunc()

	// For time
	yellowf = color.New(color.FgYellow).PrintfFunc()

	// For fatal
	bgRed = color.New(color.BgRed).PrintFunc()
)

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

func printTime() {
	if printColor {
		yellowf("%s ", time.Now().Format(timeLayout))
	} else {
		fmt.Printf("%s ", time.Now().Format(timeLayout))
	}
}

func printErrMsg() {
	if printColor {
		red("[ERR] ")
	} else {
		fmt.Print("[ERR] ")
	}
}

func printInfoMsg() {
	if printColor {
		cyan("[INFO] ")
	} else {
		fmt.Print("[INFO] ")
	}
}

func printFatalMsg() {
	if printColor {
		bgRed("[FATAL]")
	} else {
		fmt.Print("[FATAL]")
	}
	fmt.Print(" ")
}
