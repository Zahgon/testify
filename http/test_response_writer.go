package http

import (
	"net/http"
)

type TestResponseWriter struct {
	StatusCode int

	Output string

	header http.Header
}

func (rw *TestResponseWriter) Header() http.Header {
	_ = "STUB: not implemented"
	return *new(http.Header)
}

func (rw *TestResponseWriter) Write(bytes []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (rw *TestResponseWriter) WriteHeader(i int) { _ = "STUB: not implemented"; return }
