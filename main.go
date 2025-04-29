package main

import (
	"net/http"
	"time"

	"github.com/LetsFocus/account-service/empdep/handlers"
	"github.com/LetsFocus/account-service/empdep/metrics"
	"github.com/LetsFocus/account-service/empdep/services"
	"github.com/LetsFocus/account-service/empdep/store"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// responseWriter is a wrapper around http.ResponseWriter to capture the status code
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader captures the status code
func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func main() {
	store := store.New()
	app := services.New(store)
	handler := handlers.New(app)

	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/getDep", metricsMiddleware(handler.GetDepatments))
	http.HandleFunc("/createDep", metricsMiddleware(handler.CreateDepartment))
	http.HandleFunc("/updateDep", metricsMiddleware(handler.UpdateDepartment))
	http.HandleFunc("/deleteDep", metricsMiddleware(handler.DeleteDepartment))
	http.HandleFunc("/createEmp", metricsMiddleware(handler.CreateEmployee))
	http.HandleFunc("/getAll", metricsMiddleware(handler.GetEmployee))
	http.HandleFunc("/updateEmp", metricsMiddleware(handler.UpdateEmployee))
	http.HandleFunc("/deleteEmp", metricsMiddleware(handler.DeleteEmployee))

	server := &http.Server{
		Addr: ":8888",
	}
	server.ListenAndServe()
}

func metricsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		start := time.Now()

		// Wrap the original ResponseWriter
		rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		next(rw, r)

		duration := time.Since(start).Seconds()

		metrics.HTTPRequestsTotal.WithLabelValues(path).Inc()
		metrics.HTTPRequestDuration.WithLabelValues(path).Observe(duration)
	}
}
