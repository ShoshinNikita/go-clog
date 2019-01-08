package log

import (
	"fmt"
	"runtime"
	"strings"
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

// getTime returns "file:line" if l.printErrorLine == true, else it returns empty string
func (l Logger) getCaller() string {
	if !l.printErrorLine {
		return ""
	}

	var (
		file string
		line int
		ok   bool
	)

	if l.global {
		_, file, line, ok = runtime.Caller(5)
	} else {
		_, file, line, ok = runtime.Caller(4)
	}
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

// getTime returns time if l.printTime == true, else it returns empty string
func (l Logger) getTime() string {
	if !l.printTime {
		return ""
	}

	if l.printColor {
		return timePrintf("%s ", time.Now().Format(timeLayout))
	}
	return fmt.Sprintf("%s ", time.Now().Format(timeLayout))
}

func (l Logger) getInfoMsg() string {
	if l.printColor {
		return infoPrint("[INFO] ")
	}
	return "[INFO] "
}

func (l Logger) getWarnMsg() string {
	if l.printColor {
		return warnPrint("[WARN] ")
	}
	return "[WARN] "
}

func (l Logger) getErrMsg() string {
	if l.printColor {
		return errorPrint("[ERR] ")
	}
	return "[ERR] "
}

func (l Logger) getFatalMsg() (s string) {
	if l.printColor {
		return fatalPrint("[FATAL]") + " "
	}
	return "[FATAL] "
}

type prefixFunc func() string

// addPrefixes adds prefixes. It uses strings.Builder
func addPrefixes(str string, prefixes ...prefixFunc) string {
	b := strings.Builder{}

	for _, f := range prefixes {
		b.WriteString(f())
	}
	b.WriteString(str)

	return b.String()
}
