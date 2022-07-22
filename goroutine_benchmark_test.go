package logutils

import (
	"io/ioutil"
	"testing"
)

func BenchmarkGIDFilter(b *testing.B) {
	filter := &LevelFilter{
		Writer: ioutil.Discard,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		filter.Write(messages[i%len(messages)])
	}
}
