package web

import (
	"log/slog"
	"net/http"
	"time"
)

const AuthID = "middleware.auth.userID"

type Middleware func(h http.Handler) http.Handler

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func Use(ms ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(ms) - 1; i >= 0; i-- {
			m := ms[i]
			next = m(next)
		}
		return next
	}
}

func (w *wrappedWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapped := &wrappedWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}
		next.ServeHTTP(wrapped, r)
		slog.Info(
			"Got Request",
			slog.String("M", r.Method),
			slog.String("URL", r.URL.Path),
			slog.Int("S", wrapped.statusCode),
			slog.Duration("T", time.Since(start)),
		)
	})
}
