package server

import (
	"fmt"
	"net/http"
	"time"
)

// Take a list of middleware functions and return a closure that will take a function
// and run the function with the middleware in the stack
func makeStack(chain ...func(http.Handler) http.Handler) func(http.Handler) http.Handler {
	return func(final http.Handler) http.Handler {
		return pipe(final, chain...)
	}
}

func pipe(final http.Handler, chain ...func(http.Handler) http.Handler) http.Handler {
	method := final
	for _, m := range chain {
		method = m(method)
	}
	return method
}

func web(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}

type logResponseWriter struct {
	StatusCode int
	baseWriter http.ResponseWriter
}

func newLogResponseWriter(w http.ResponseWriter) *logResponseWriter {
	return &logResponseWriter{
		StatusCode: 200,
		baseWriter: w,
	}
}

func (l logResponseWriter) Header() http.Header {
	return l.baseWriter.Header()
}

func (l logResponseWriter) Write(w []byte) (int, error) {
	return l.baseWriter.Write(w)
}

func (l *logResponseWriter) WriteHeader(statusCode int) {
	l.StatusCode = statusCode
	l.baseWriter.WriteHeader(statusCode)
}

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now().UnixMilli()
		logHandler := newLogResponseWriter(w)

		next.ServeHTTP(logHandler, r)

		runtime := time.Now().UnixMilli() - start
		fmt.Println(
			fmt.Sprintf(
				"[%s] %s --- %d in %dms",
				r.Method,
				r.RequestURI,
				logHandler.StatusCode,
				runtime,
			),
		)
	})
}
