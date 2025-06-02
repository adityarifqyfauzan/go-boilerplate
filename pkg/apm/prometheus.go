package apm

import (
	"runtime"
	"time"

	"github.com/adityarifqyfauzan/go-boilerplate/config"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	httpRequestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"path", "method"},
	)

	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"path", "method"},
	)

	dbOpenConnections = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "db_open_connections",
			Help: "Number of open connections to the database",
		},
	)

	dbIdleConnections = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "db_idle_connections",
			Help: "Number of idle connections in the pool",
		},
	)

	goRoutines = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "go_routines",
			Help: "Number of active goroutines",
		},
	)
)

func init() {
	prometheus.MustRegister(
		httpRequestCount,
		httpRequestDuration,
		dbOpenConnections,
		dbIdleConnections,
		goRoutines,
	)
}

func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer for the request
		start := time.Now()
		c.Next()
		// Stop timer for the request
		elapsed := time.Since(start)
		path := c.FullPath()
		if path == "" {
			path = c.Request.URL.Path
		}

		httpRequestCount.WithLabelValues(path, c.Request.Method).Inc()
		httpRequestDuration.WithLabelValues(path, c.Request.Method).Observe(elapsed.Seconds())
	}
}

func CollectRuntimeMetrics() {
	go func() {
		for {
			goRoutines.Set(float64(runtime.NumGoroutine()))

			if config.SqlDB == nil {
				dbOpenConnections.Set(float64(config.SqlDB.Stats().OpenConnections))
				dbIdleConnections.Set(float64(config.SqlDB.Stats().Idle))
			}

			time.Sleep(5 * time.Second)
		}
	}()
}
