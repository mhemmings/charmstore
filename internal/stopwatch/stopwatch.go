package stopwatch

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Stopwatch struct {
	start       time.Time
	writer      io.Writer
	writeFormat string
	writeArgs   []interface{}
}

func New(format string, a ...interface{}) *Stopwatch {
	return &Stopwatch{
		start:       time.Now(),
		writer:      os.Stdout,
		writeFormat: format,
		writeArgs:   a,
	}
}

func Fnew(writer io.Writer, format string, a ...interface{}) *Stopwatch {
	return &Stopwatch{
		start:       time.Now(),
		writer:      writer,
		writeFormat: format,
		writeArgs:   a,
	}
}

func (s *Stopwatch) Time(f func(), format string, a ...interface{}) {
	sw := New(format, a...)
	f()
	sw.Done()
}

func (s *Stopwatch) Reset() {
	s.start = time.Now()
}

func (s *Stopwatch) Done() {
	var durstr string
	dur := int64(time.Since(s.start) / time.Millisecond)
	if dur > 0 {
		durstr = fmt.Sprintf("%dms", dur)
	} else {
		durstr = "<1ms"
	}
	fmt.Fprintf(s.writer, "%s: ", durstr)
	fmt.Fprintf(s.writer, s.writeFormat, s.writeArgs...)
	fmt.Fprintln(s.writer)
}
