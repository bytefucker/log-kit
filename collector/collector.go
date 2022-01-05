package collector

import (
	"github.com/yihongzhi/log-kit/collector/dest"
	"github.com/yihongzhi/log-kit/collector/source"
	"github.com/yihongzhi/log-kit/config"
	"time"
)

type LogCollector struct {
	source source.LogSource
	dest   dest.LogDestination
}

func NewLogCollector(config *config.CollectorConfig) (*LogCollector, error) {
	s, err := source.NewFileSource(config.Source)
	if err != nil {
		return nil, err
	}
	d, err := dest.NewKafkaDestination(config.Destination.Kafka)
	if err != nil {
		return nil, err
	}
	return &LogCollector{
		source: s,
		dest:   d,
	}, nil
}

func (c *LogCollector) Start() error {
	if err := c.source.Start(); err != nil {
		return err
	}
	for true {
		log := c.source.GetMessage()
		message := dest.LogMessage{
			Time:    time.Now(),
			Host:    "",
			AppId:   log.AppId,
			Content: log.Content,
		}
		c.dest.Send(&message)
	}
	return nil
}
