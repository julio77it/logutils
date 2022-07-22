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

	// The Printf format for logging :
	//    %d for the Goroutine ID
	// that will
	Format string

	once sync.Once
}

// get Gorouting ID from go runtime
func getGID() uint64 {
	// Scott Mansfield
	// Goroutine IDs
	// https://blog.sgmansfield.com/2015/12/goroutine-ids/
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

func (f *GIDFilter) init() {
	if len(f.Format) == 0 {
		f.Format = "[%d]"
	}
}

func (f *GIDFilter) Write(p []byte) (n int, err error) {
	// Note in general that io.Writer can receive any byte sequence
	// to write, but the "log" package always guarantees that we only
	// get a single line. We use that as a slight optimization within
	// this method, assuming we're dealing with a single, complete line
	// of log data.

	f.once.Do(f.init)

	gid := fmt.Sprintf(f.Format, getGID())

	p = append([]byte(gid), p...)

	return f.Writer.Write(p)
}
