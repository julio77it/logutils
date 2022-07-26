package logutils

import (
	"io/ioutil"
	"testing"
)

var gMessages [][]byte

func init() {
	gMessages = [][]byte{
		[]byte("[GID] foo"),
		[]byte("[GID] bar"),
	}
}

func BenchmarkGIDFilter(b *testing.B) {
	filter := &GIDFilter{
		Writer: ioutil.Discard,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		filter.Write(gMessages[i%len(gMessages)])
	}
}

func BenchmarkGIDFilterWithoutGID(b *testing.B) {
	filter := &GIDFilter{
		Writer:    ioutil.Discard,
		GIDString: "[NOGID]",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		filter.Write(gMessages[i%len(gMessages)])
	}
}

func BenchmarkGetGID(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		getGID()
	}
}
