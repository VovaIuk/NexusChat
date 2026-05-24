package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	RequestsTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "nexus_http_requests_total",
		Help: "Общее количество HTTP-запросов",
	}, []string{"method", "path", "status"})

	RequestDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "nexus_http_request_duration_seconds",
		Help:    "Длительность HTTP-запросов в секундах",
		Buckets: prometheus.DefBuckets,
	}, []string{"method", "path"})

	MessagesCreatedTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "nexus_messages_created_total",
		Help: "Количество сообщений, сохранённых в БД",
	})

	WsConnections = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "nexus_ws_connections",
		Help: "Число активных WebSocket-соединений (после авторизации)",
	})

	WsBroadcastTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "nexus_ws_broadcast_total",
		Help: "Количество broadcast-сообщений по WebSocket",
	})

	ChatHistoryRequestsTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "nexus_chat_history_requests_total",
		Help: "Запросы истории сообщений чата",
	}, []string{"paginated"})

	ChatMessagesReturned = promauto.NewHistogram(prometheus.HistogramOpts{
		Name:    "nexus_chat_messages_returned",
		Help:    "Число сообщений, возвращённых за один запрос истории",
		Buckets: []float64{5, 10, 25, 50, 100, 250, 500},
	})

	ChatMessagesLimit = promauto.NewHistogram(prometheus.HistogramOpts{
		Name:    "nexus_chat_messages_limit",
		Help:    "Параметр limit в запросах истории сообщений",
		Buckets: []float64{5, 10, 25, 50, 100, 250, 500},
	})
)

func RecordChatHistory(paginated bool, limit, returned int) {
	paginatedLabel := "false"
	if paginated {
		paginatedLabel = "true"
	}
	ChatHistoryRequestsTotal.WithLabelValues(paginatedLabel).Inc()
	ChatMessagesLimit.Observe(float64(limit))
	ChatMessagesReturned.Observe(float64(returned))
}
