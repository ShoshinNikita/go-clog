package log_test

import (
	"fmt"
	defLog "log"
	"testing"

	"github.com/ShoshinNikita/log"
)

func BenchmarkInfoln(b *testing.B) {
	for i := 0; i < b.N; i++ {
		log.Infoln("Hello, dear world")
	}
}

func BenchmarkFmtPrintln(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Println("Hello, dear world")
	}
}

func BenchmarkLogPrintln(b *testing.B) {
	for i := 0; i < b.N; i++ {
		defLog.Println("Hello, dear world")
	}
}
