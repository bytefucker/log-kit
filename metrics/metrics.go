package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

var logCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "task_counter",
}, []string{"app_id"})

type Client struct {
	Port string
}

func NewMetricsClient(port string) *Client {
	prometheus.MustRegister(logCounter)
	return &Client{
		Port: port,
	}
}

func (c *Client) Start() error {
	http.Handle("/metrics", promhttp.Handler())
	return http.ListenAndServe(":"+c.Port, nil)
}
