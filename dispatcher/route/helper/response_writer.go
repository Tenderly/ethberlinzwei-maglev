package helper

import "net/http"

type WrappedResponseWriter struct {
	w                http.ResponseWriter
	StatusCode       int
	ResponseBodySize int
}

func (w *WrappedResponseWriter) Header() http.Header {
	return w.w.Header()
}

func (w *WrappedResponseWriter) Write(bytes []byte) (int, error) {
	result, err := w.w.Write(bytes)
	if err == nil {
		w.ResponseBodySize += len(bytes)
	}

	return result, err
}

func (w *WrappedResponseWriter) WriteHeader(statusCode int) {
	w.StatusCode = statusCode
	w.w.WriteHeader(statusCode)
}

func NewWrappedResponseWriter(w http.ResponseWriter) http.ResponseWriter {
	return &WrappedResponseWriter{
		w: w,
	}
}
