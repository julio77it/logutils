// Package logutils augments the standard log package with goroutine id.
package logutils

import (
	"bytes"
	"fmt"
	"io"
	"runtime"
	"strconv"
	"sync"
)

// LevelFilter is an io.Writer that can be used with a logger that
// will filter out log messages that aren't at least a certain level.
//
// Once the filter is in use somewhere, it is not safe to modify
// the structure.
type GIDFilter struct {
	// The underlying io.Writer where log messages that pass the filter
	// will be set.
	Writer io.Writer

	// the string to be substuited in the logline by gid
	GIDString string
	gIDbytes  []byte
	Format    string

	once sync.Once
}

// get Gorouting ID from go runtime
func getGID() uint64 {
	b := make([]byte, 32)

	runtime.Stack(b[:], false)

	b = b[10:20]
	b = b[:bytes.IndexByte(b, ' ')]
	id, _ := strconv.ParseUint(string(b), 10, 64)
	return id
}

func (f *GIDFilter) init() {
	if len(f.GIDString) == 0 {
		f.GIDString = "[GID]"
	}
	if len(f.Format) == 0 {
		f.Format = "[%d]"
	}
	f.gIDbytes = []byte(f.GIDString)
}

func (f *GIDFilter) Write(p []byte) (n int, err error) {
	// Note in general that io.Writer can receive any byte sequence
	// to write, but the "log" package always guarantees that we only
	// get a single line. We use that as a slight optimization within
	// this method, assuming we're dealing with a single, complete line
	// of log data.
	f.once.Do(f.init)

	if bytes.Contains(p, f.gIDbytes) {
		gid := fmt.Sprintf(f.Format, getGID())
		p = bytes.Replace(p, f.gIDbytes, []byte(gid), 1)
	}
	return f.Writer.Write(p)
}
