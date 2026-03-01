package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	RequestsTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "nexus_http_requests_total",
		Help: "Общее количество HTTP-запросов",
	})
)
