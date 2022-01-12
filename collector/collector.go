package collector

import (
	log "github.com/sirupsen/logrus"
	"github.com/yihongzhi/log-kit/collector/sender"
	"github.com/yihongzhi/log-kit/collector/source"
	"github.com/yihongzhi/log-kit/config"
	"github.com/yihongzhi/log-kit/metrics"
	"os"
	"time"
)

type LogCollector struct {
	source  source.LogSource
	sender  sender.LogMessageSender
	metrics *metrics.Client
}

func NewLogCollector(config *config.AppConfig) (*LogCollector, error) {
	source, err := source.NewFileSource(config.Source, config.BufferSize)
	if err != nil {
		log.Errorln("Init LogSource error", err)
		return nil, err
	}
	sender, err := sender.NewKafkaSender(config.Kafka)
	if err != nil {
		log.Errorln("Init LogMessageSender error", err)
		return nil, err
	}
	client := metrics.NewMetricsClient(config.Manager.Port)
	return &LogCollector{
		source:  source,
		sender:  sender,
		metrics: client,
	}, nil
}

// Start 开启日志收集任务
func (c *LogCollector) Start() error {
	if err := c.metrics.Start(); err != nil {
		return err
	}
	if err := c.source.Start(); err != nil {
		return err
	}
	hostname, _ := os.Hostname()
	for true {
		log := c.source.GetMessage()
		message := sender.LogMessage{
			Time:    time.Now(),
			Host:    hostname,
			AppId:   log.AppId,
			Content: log.Content,
		}
		c.sender.SendMessage(&message)
	}
	return nil
}
