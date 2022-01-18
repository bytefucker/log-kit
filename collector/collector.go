package collector

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/yihongzhi/log-kit/collector/sender"
	"github.com/yihongzhi/log-kit/collector/source"
	"github.com/yihongzhi/log-kit/config"
	"github.com/yihongzhi/log-kit/logger"
	"net/http"
	"os"
	"time"
)

var log = logger.Log

type LogCollector struct {
	port   string
	source source.LogSource
	sender sender.LogMessageSender
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
	return &LogCollector{
		port:   config.Port,
		source: source,
		sender: sender,
	}, nil
}

// Start 开启日志收集任务
func (c *LogCollector) Start() error {
	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(":"+c.port, nil)
	if err := c.source.Start(); err != nil {
		return err
	}
	hostname, _ := os.Hostname()
	for {
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
