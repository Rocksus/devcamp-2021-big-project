package monitoring

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{w, http.StatusOK}
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		route := mux.CurrentRoute(r)
		path, _ := route.GetPathTemplate()

		timer := prometheus.NewTimer(latency.WithLabelValues(path))

		rw := newResponseWriter(w)
		next.ServeHTTP(rw, r)

		responseStatus.WithLabelValues(strconv.Itoa(rw.statusCode)).Inc()
		totalRequests.WithLabelValues(path).Inc()
		timer.ObserveDuration()
	})
}
