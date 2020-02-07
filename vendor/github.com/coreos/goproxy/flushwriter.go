package goproxy

import (
	"io"
	"net/http"
)

type FlushWriter struct {
	w io.Writer
}

func NewFlushWriter(w io.Writer) *FlushWriter {
	return &FlushWriter{w: w}
}

func (w *FlushWriter) Write(b []byte) (int, error) {
	bytes, err := w.w.Write(b)
	if f, ok := w.w.(http.Flusher); ok {
		f.Flush()
	}
	return bytes, err
}
