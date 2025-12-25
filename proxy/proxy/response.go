package proxy

import (
	"bytes"
	"net/http"
)

// ResponseWrapper 包装 ResponseWriter 以捕获状态码、响应头和响应体
type ResponseWrapper struct {
	http.ResponseWriter
	StatusCode   int
	Headers      http.Header
	ResponseBody *bytes.Buffer
}

func (w *ResponseWrapper) WriteHeader(statusCode int) {
	w.StatusCode = statusCode
	// 复制响应头
	if w.Headers == nil {
		w.Headers = make(http.Header)
	}
	for k, v := range w.ResponseWriter.Header() {
		w.Headers[k] = v
	}
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *ResponseWrapper) Write(b []byte) (int, error) {
	// 累积响应体
	if w.ResponseBody == nil {
		w.ResponseBody = &bytes.Buffer{}
	}
	w.ResponseBody.Write(b)
	return w.ResponseWriter.Write(b)
}

// Hijack 支持 WebSocket
func (w *ResponseWrapper) Hijack() (interface{}, interface{}, error) {
	if hijacker, ok := w.ResponseWriter.(http.Hijacker); ok {
		return hijacker.Hijack()
	}
	return nil, nil, http.ErrNotSupported
}
