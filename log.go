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
	red   = color.New(color.FgRed).PrintFunc()
	redf  = color.New(color.FgRed).PrintfFunc()
	redln = color.New(color.FgRed).PrintlnFunc()

	// For [INFO]
	cyan   = color.New(color.FgCyan).PrintFunc()
	cyanf  = color.New(color.FgCyan).PrintfFunc()
	cyanln = color.New(color.FgCyan).PrintlnFunc()

	// For time
	yellow   = color.New(color.FgYellow).PrintFunc()
	yellowf  = color.New(color.FgYellow).PrintfFunc()
	yellowln = color.New(color.FgYellow).PrintlnFunc()
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
