package collector

import "github.com/yihongzhi/log-kit/internal/kafka"

type LogCollector struct {
	kafkaClient *kafka.KafkaClient
}

func (c *LogCollector) Start() {

}
