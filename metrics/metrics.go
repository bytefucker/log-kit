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

var logKafkaReadCount = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "log_kafka_read_count",
}, []string{"app_id"})

func init() {
	prometheus.MustRegister(logTailReadCount, logKafkaSendCount, logKafkaReadCount)
}

func ReadFileLogInc(appId string) {
	logTailReadCount.With(prometheus.Labels{
		"app_id": appId,
	}).Inc()
}

func SendKafkaLogInc(appId string) {
	logKafkaSendCount.With(prometheus.Labels{
		"app_id": appId,
	}).Inc()
}

func ReadKafkaLogInc(appId string) {
	logKafkaReadCount.With(prometheus.Labels{
		"app_id": appId,
	}).Inc()
}
