package prom

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
    GetRequestCounter = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "api_get_requests_total",
            Help: "Total number of GET requests.",
        },
        []string{"path"},
    )
    PutRequestCounter = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "api_put_requests_total",
            Help: "Total number of PUT requests.",
        },
        []string{"path"},
    )
)

func InitPrometheus() {
    prometheus.MustRegister(GetRequestCounter)
    prometheus.MustRegister(PutRequestCounter)
}