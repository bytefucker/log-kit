package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var logTailReadCount = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "log_tail_read_count",
}, []string{"app_id"})

var logKafkaSendCount = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "log_kafka_send_count",
}, []string{"app_id"})

func init() {
	prometheus.MustRegister(logTailReadCount, logKafkaSendCount)
}

func LogReadInc(appId string) {
	logTailReadCount.With(prometheus.Labels{
		"app_id": appId,
	}).Inc()
}

func LogSendInc(appId string) {
	logKafkaSendCount.With(prometheus.Labels{
		"app_id": appId,
	}).Inc()
}
