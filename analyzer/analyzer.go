package analyzer

import (
	"context"
	"github.com/yihongzhi/log-kit/analyzer/parser"
	"github.com/yihongzhi/log-kit/config"
	"github.com/yihongzhi/log-kit/elastic"
	"github.com/yihongzhi/log-kit/kafka"
)

// LogAnalyzer 日志解析服务
type LogAnalyzer struct {
	TopName       string
	EsClient      *elastic.ESClient
	KafkaConsumer *kafka.Consumer
	LogParser     parser.LogParser
}

func NewLogAnalyzer(config *config.AppConfig) (*LogAnalyzer, error) {
	esClient, err := elastic.NewESClient(config.Elastic)
	if err != nil {
		return nil, err
	}
	kafkaConsumer, err := kafka.NewKafkaConsumer(config.Kafka)
	if err != nil {
		return nil, err
	}
	return &LogAnalyzer{
		TopName:       config.Kafka.TopicName,
		EsClient:      esClient,
		KafkaConsumer: kafkaConsumer,
		LogParser:     nil,
	}, nil
}

func (a *LogAnalyzer) Start() error {
	err := a.KafkaConsumer.Consume(context.Background(), []string{a.TopName}, nil)
	if err != nil {
		return err
	}
	return nil
}
