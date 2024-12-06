package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	CommandDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "command_duration_seconds",
			Help:    "Duration of commands in seconds",
			Buckets: []float64{.005, .01, .025, .05, .1, .25, .5, 1},
		},
		[]string{"command_type", "status"},
	)

	QueryDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "query_duration_seconds",
			Help:    "Duration of queries in seconds",
			Buckets: []float64{.005, .01, .025, .05, .1, .25, .5, 1},
		},
		[]string{"query_type", "status"},
	)

	CommandCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "commands_total",
			Help: "Total number of commands processed",
		},
		[]string{"command_type", "status"},
	)

	QueryCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "queries_total",
			Help: "Total number of queries processed",
		},
		[]string{"query_type", "status"},
	)
)
