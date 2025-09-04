package monitoring

import (
	"net/http"
	"strconv"
	"time"
)

func MetricsMiddleware(m *Metrics) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			rw := &statusRecorder{ResponseWriter: w, statusCode: 200}

			next.ServeHTTP(rw, r)

			m.ObserveRequest(
				r.Method,
				r.URL.Path,
				strconv.Itoa(rw.statusCode),
				time.Since(start),
			)
		})
	}
}

type statusRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (rw *statusRecorder) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
