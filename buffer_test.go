package log

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestBufferWriterSize(t *testing.T) {
	w := &BufferWriter{
		BufferSize:    1000,
		FlushDuration: 1000 * time.Millisecond,
		Writer:        os.Stderr,
	}
	for i := 0; i < 100; i++ {
		fmt.Fprintf(w, "%s, %d during buffer writer 1k buff size\n", timeNow(), i)
	}
	time.Sleep(time.Second)
}

func TestBufferWriterSizeZero(t *testing.T) {
	w := &BufferWriter{
		BufferSize: 0,
		Writer:     os.Stderr,
	}
	fmt.Fprintf(w, "%s, before buffer writer zero size\n", timeNow())
	time.Sleep(1100 * time.Millisecond)
	fmt.Fprintf(os.Stderr, "%s, after buffer writer zero size\n", timeNow())
}

func TestBufferWriterDuration(t *testing.T) {
	w := &BufferWriter{
		BufferSize:    1000,
		FlushDuration: 10 * time.Millisecond,
		Writer:        os.Stderr,
	}
	fmt.Fprintf(w, "%s, during buffer writer tiny duration\n", timeNow())
	time.Sleep(200 * time.Millisecond)
}

func TestBufferWriterFlushAuto(t *testing.T) {
	w := &BufferWriter{
		BufferSize:    8192,
		FlushDuration: 1000 * time.Millisecond,
		Writer:        os.Stderr,
	}
	fmt.Fprintf(w, "%s, before buffer writer auto flush\n", timeNow())
	time.Sleep(1100 * time.Millisecond)
	fmt.Fprintf(os.Stderr, "%s, after buffer writer auto flush\n", timeNow())
}

func TestBufferWriterFlushCall(t *testing.T) {
	w := &BufferWriter{
		BufferSize:    8192,
		FlushDuration: 1000 * time.Millisecond,
		Writer:        os.Stderr,
	}
	fmt.Fprintf(w, "%s, before buffer writer flush\n", timeNow())
	w.Flush()
	fmt.Fprintf(os.Stderr, "%s, after buffer writer flush\n", timeNow())
}

func TestBufferWriterFlusher(t *testing.T) {
	w := &BufferWriter{
		BufferSize:    8192,
		FlushDuration: 1000 * time.Millisecond,
		Writer:        os.Stderr,
	}
	fmt.Fprintf(w, "%s, before buffer writer flush\n", timeNow())

	Flush(w)
}

func TestBufferWriterClose(t *testing.T) {
	w := &BufferWriter{
		BufferSize:    8192,
		FlushDuration: 1000 * time.Millisecond,
		Writer:        func() *os.File { f, _ := os.Open(os.DevNull); return f }(),
	}
	fmt.Fprintf(w, "%s, before buffer writer flush\n", timeNow())
	w.Close()
	fmt.Fprintf(os.Stderr, "%s, after buffer writer flush\n", timeNow())
}
