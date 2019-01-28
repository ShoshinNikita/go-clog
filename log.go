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
	"os"
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

func NewDevLogger() *Logger {
	c := &Config{
		output:         color.Output,
		printTime:      true,
		printColor:     true,
		printErrorLine: true,
		timeLayout:     DefaultTimeLayout,
	}

	return c.Build()
}

func NewProdLogger() *Logger {
	c := &Config{
		output:         os.Stdout,
		printTime:      true,
		printColor:     false,
		printErrorLine: true,
		timeLayout:     DefaultTimeLayout,
	}

	return c.Build()
}

type Config struct {
	output         io.Writer
	printTime      bool
	printColor     bool
	printErrorLine bool
	timeLayout     string
}

func (c *Config) Build() *Logger {
	l := new(Logger)
	l.mutex = new(sync.Mutex)

	l.output = color.Output
	if c.output != nil {
		l.output = c.output
	}
	l.printTime = c.printTime
	l.printColor = c.printColor
	l.printErrorLine = c.printErrorLine
	l.timeLayout = c.timeLayout

	return l
}

// PrintTime sets Logger.printTime to b
func (c *Config) PrintTime(b bool) *Config {
	c.printTime = b
	return c
}

// PrintColor sets Logger.printColor to b
func (c *Config) PrintColor(b bool) *Config {
	c.printColor = b
	return c
}

// PrintErrorLine sets Logger.printErrorLine to b
func (c *Config) PrintErrorLine(b bool) *Config {
	c.printErrorLine = b
	return c
}

// ChangeOutput changes Logger.output writer.
// Default Logger.output is github.com/fatih/color.Output
func (c *Config) ChangeOutput(w io.Writer) *Config {
	c.output = w
	return c
}

// ChangeTimeLayout changes Logger.timeLayout
// Default Logger.timeLayout is DefaultTimeLayout
func (c *Config) ChangeTimeLayout(layout string) *Config {
	c.timeLayout = layout
	return c
}
