package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// HTTP request metrics
	HTTPRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"path"},
	)

	HTTPRequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"path"},
	)
	// Database connection metrics
	ActiveConnections = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "db_active_connections",
			Help: "Number of active database connections.",
		},
	)

	IdleConnections = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "db_idle_connections",
			Help: "Number of idle database connections.",
		},
	)

	TotalConnections = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "db_total_connections",
			Help: "Total number of database connections.",
		},
	)

	FailedConnections = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "db_failed_connections",
			Help: "Number of failed database connection attempts.",
		},
	)
)
