package logutils

import (
	"io/ioutil"
	"testing"
)

var mMessages [][]byte

func init() {
	mMessages = [][]byte{
		[]byte("[GID][TRACE] foo"),
		[]byte("[GID][DEBUG] foo"),
		[]byte("[GID][INFO] foo"),
		[]byte("[GID][WARN] foo"),
		[]byte("[GID][ERROR] foo"),
	}
}

func BenchmarkMixedFilter(b *testing.B) {
	filter := GIDFilter{
		Writer: &LevelFilter{
			Levels:   []LogLevel{"TRACE", "DEBUG", "INFO", "WARN", "ERROR"},
			MinLevel: "WARN",
			Writer:   ioutil.Discard,
		},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		filter.Write(mMessages[i%len(mMessages)])
	}
}
