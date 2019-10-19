package clog

import (
	"bytes"
	stdlog "log"
	"os"
	"testing"
)

// -----------------------------------------------------------------------------
// Tests
// -----------------------------------------------------------------------------

func TestLoggerLevels(t *testing.T) {
	printFunction := func(log *Logger) {
		log.Debug("debug")
		log.Debugf("debugf %s\n", "arg")

		log.Info("info")
		log.Infof("infof %s\n", "arg")

		log.Warn("warn")
		log.Warnf("warnf %s %d\n", "arg", 15)

		log.Error("error")
		log.Errorf("errorf %s\n", "arg")

		// log.Fatal("fatal")
		// log.Fatalf("fatalf %s", "arg")

		log.Print("print")
		log.Printf("printf %s\n", "arg")

		log.Write([]byte("bytes"))
		log.WriteString("string")
	}

	tests := []struct {
		description string
		config      *Config
		output      []byte
	}{
		{
			description: "debug level",
			config: &Config{
				level:          LevelDebug,
				printColor:     false,
				printErrorLine: false,
				printTime:      false,
				timeLayout:     DefaultTimeLayout,
			},
			output: []byte(
				"[DBG] debug\n" +
					"[DBG] debugf arg\n" +
					"[INF] info\n" +
					"[INF] infof arg\n" +
					"[WRN] warn\n" +
					"[WRN] warnf arg 15\n" +
					"[ERR] error\n" +
					"[ERR] errorf arg\n" +
					"print\n" +
					"printf arg\n" +
					"bytesstring"),
		},
		{
			description: "info level",
			config: &Config{
				level:          LevelInfo,
				printColor:     false,
				printErrorLine: false,
				printTime:      false,
				timeLayout:     DefaultTimeLayout,
			},
			output: []byte(
				"[INF] info\n" +
					"[INF] infof arg\n" +
					"[WRN] warn\n" +
					"[WRN] warnf arg 15\n" +
					"[ERR] error\n" +
					"[ERR] errorf arg\n" +
					"print\n" +
					"printf arg\n" +
					"bytesstring"),
		},
		{
			description: "warn level",
			config: &Config{
				level:          LevelWarn,
				printColor:     false,
				printErrorLine: false,
				printTime:      false,
				timeLayout:     DefaultTimeLayout,
			},
			output: []byte(
				"[WRN] warn\n" +
					"[WRN] warnf arg 15\n" +
					"[ERR] error\n" +
					"[ERR] errorf arg\n" +
					"print\n" +
					"printf arg\n" +
					"bytesstring"),
		},
		{
			description: "error level",
			config: &Config{
				level:          LevelError,
				printColor:     false,
				printErrorLine: false,
				printTime:      false,
				timeLayout:     DefaultTimeLayout,
			},
			output: []byte(
				"[ERR] error\n" +
					"[ERR] errorf arg\n" +
					"print\n" +
					"printf arg\n" +
					"bytesstring"),
		},
		{
			description: "fatal level",
			config: &Config{
				level:          LevelFatal,
				printColor:     false,
				printErrorLine: false,
				printTime:      false,
				timeLayout:     DefaultTimeLayout,
			},
			output: []byte(
				"print\n" +
					"printf arg\n" +
					"bytesstring"),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.description, func(t *testing.T) {
			buff := &bytes.Buffer{}
			tt.config.SetOutput(buff)

			log := tt.config.Build()

			printFunction(log)

			res := buff.Bytes()
			if !bytes.Equal(res, tt.output) {
				t.Errorf("different output")
				t.Log(string(res))
				t.Log(string(tt.output))
			}
		})
	}
}

func TestWithPrefix(t *testing.T) {
	printFunction := func(log *Logger) {
		log.Debug("debug")
		log.Debugf("debugf %s\n", "arg")

		log.Info("info")
		log.Infof("infof %s\n", "arg")

		log.Warn("warn")
		log.Warnf("warnf %s %d\n", "arg", 15)

		log.Error("error")
		log.Errorf("errorf %s\n", "arg")

		// log.Fatal("fatal")
		// log.Fatalf("fatalf %s", "arg")

		log.Print("print")
		log.Printf("printf %s\n", "arg")

		log.Write([]byte("bytes"))
		log.WriteString("string")
	}

	tests := []struct {
		log    *Logger
		output []byte
	}{
		{
			log: NewDevConfig().
				PrintColor(false).
				PrintErrorLine(false).
				PrintTime(false).
				SetPrefix("prefix").
				Build(),
			output: []byte(
				"[DBG] prefixdebug\n" +
					"[DBG] prefixdebugf arg\n" +
					"[INF] prefixinfo\n" +
					"[INF] prefixinfof arg\n" +
					"[WRN] prefixwarn\n" +
					"[WRN] prefixwarnf arg 15\n" +
					"[ERR] prefixerror\n" +
					"[ERR] prefixerrorf arg\n" +
					"prefixprint\n" +
					"prefixprintf arg\n" +
					"bytesstring"),
		},
		{
			log: NewDevConfig().
				PrintColor(false).
				PrintErrorLine(false).
				PrintTime(false).
				Build().WithPrefix("prefix"),
			output: []byte(
				"[DBG] prefix: debug\n" +
					"[DBG] prefix: debugf arg\n" +
					"[INF] prefix: info\n" +
					"[INF] prefix: infof arg\n" +
					"[WRN] prefix: warn\n" +
					"[WRN] prefix: warnf arg 15\n" +
					"[ERR] prefix: error\n" +
					"[ERR] prefix: errorf arg\n" +
					"prefix: print\n" +
					"prefix: printf arg\n" +
					"bytesstring"),
		},
		{
			log: NewDevConfig().
				PrintColor(false).
				PrintErrorLine(false).
				PrintTime(false).
				Build().WithPrefix("first prefix").WithPrefix("second prefix"),
			output: []byte(
				"[DBG] second prefix: first prefix: debug\n" +
					"[DBG] second prefix: first prefix: debugf arg\n" +
					"[INF] second prefix: first prefix: info\n" +
					"[INF] second prefix: first prefix: infof arg\n" +
					"[WRN] second prefix: first prefix: warn\n" +
					"[WRN] second prefix: first prefix: warnf arg 15\n" +
					"[ERR] second prefix: first prefix: error\n" +
					"[ERR] second prefix: first prefix: errorf arg\n" +
					"second prefix: first prefix: print\n" +
					"second prefix: first prefix: printf arg\n" +
					"bytesstring"),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run("", func(t *testing.T) {
			buff := &bytes.Buffer{}
			tt.log.output = buff

			printFunction(tt.log)

			res := buff.Bytes()
			if !bytes.Equal(res, tt.output) {
				t.Errorf("different output")
				t.Log(string(res))
				t.Log(string(tt.output))
			}
		})
	}
}

// -----------------------------------------------------------------------------
// Benchmarks
// -----------------------------------------------------------------------------

const (
	file = "test.txt"
	msg  = "Hello, dear world!!!"
)

func BenchmarkStdLogPrintlnWithPrefixes(b *testing.B) {
	f, err := os.Create(file)
	if err != nil {
		stdlog.Fatalln(err)
	}
	defer f.Close()

	l := &stdlog.Logger{}

	// github.com/ShoshinNikita/log prints it by default
	l.SetFlags(stdlog.Lshortfile | stdlog.Ltime | stdlog.Ldate)
	l.SetOutput(f)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.Println(msg)
	}
}

func BenchmarkDevLogPrintln(b *testing.B) {
	f, err := os.Create(file)
	if err != nil {
		stdlog.Fatalln(err)
	}
	defer f.Close()

	l := NewDevConfig().SetOutput(f).Build()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.Print(msg)
	}
}

func BenchmarkDevLogErrorln(b *testing.B) {
	f, err := os.Create(file)
	if err != nil {
		stdlog.Fatalln(err)
	}
	defer f.Close()

	l := NewDevConfig().SetOutput(f).Build()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.Error(msg)
	}
}

func BenchmarkProdLogPrintln(b *testing.B) {
	f, err := os.Create(file)
	if err != nil {
		stdlog.Fatalln(err)
	}
	defer f.Close()

	l := NewProdConfig().SetOutput(f).Build()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.Print(msg)
	}
}

func BenchmarkProdLogErrorln(b *testing.B) {
	f, err := os.Create(file)
	if err != nil {
		stdlog.Fatalln(err)
	}
	defer f.Close()

	l := NewProdConfig().SetOutput(f).Build()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.Error(msg)
	}
}
