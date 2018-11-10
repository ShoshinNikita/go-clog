package log

import (
	"fmt"
	"runtime"
	"time"

	"github.com/fatih/color"
)

var (
	// For time
	timePrintf = color.New(color.FgHiGreen).SprintfFunc()

	// For [INFO]
	infoPrint = color.New(color.FgCyan).SprintFunc()

	// For [WARN]
	warnPrint = color.New(color.FgYellow).SprintFunc()

	// For [ERR]
	errorPrint = color.New(color.FgRed).SprintFunc()

	// For [FATAL]
	fatalPrint = color.New(color.BgRed).SprintFunc()
)

func getCaller() string {
	// We need to skip 2 functions (this and log.Error(), log.Errorf() and etc.)
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		return ""
	}
	var shortFile string
	for i := len(file) - 1; i >= 0; i-- {
		if file[i] == '/' {
			shortFile = file[i+1:]
			break
		}
	}
	return fmt.Sprintf("%s:%d ", shortFile, line)
}

func (l Logger) getTime() string {
	if l.PrintColor {
		return timePrintf("%s ", time.Now().Format(timeLayout))
	}
	return fmt.Sprintf("%s ", time.Now().Format(timeLayout))
}

func (l Logger) getInfoMsg() string {
	if l.PrintColor {
		return infoPrint("[INFO] ")
	}
	return "[INFO] "
}

func (l Logger) getWarnMsg() string {
	if l.PrintColor {
		return warnPrint("[WARN] ")
	}
	return "[WARN] "
}

func (l Logger) getErrMsg() string {
	if l.PrintColor {
		return errorPrint("[ERR] ")
	}
	return "[ERR] "
}

func (l Logger) getFatalMsg() (s string) {
	if l.PrintColor {
		return fatalPrint("[FATAL]") + " "
	}
	return "[FATAL] "
}
