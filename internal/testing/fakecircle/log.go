package fakecircle

import (
	"log/slog"
	"net/http"
	"time"
)

func loggingHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		path := r.URL.Path
		query := r.URL.RawQuery

		hr := &httpWriter{
			ResponseWriter: w,
		}
		next.ServeHTTP(hr, r)

		slog.Info(r.Method,
			"path", path,
			"status", hr.statusCode,
			"query", query,
			"duration", time.Since(start),
		)
	})
}

type httpWriter struct {
	http.ResponseWriter
	statusCode int
}

func (h *httpWriter) WriteHeader(statusCode int) {
	h.statusCode = statusCode
	h.ResponseWriter.WriteHeader(statusCode)
}
