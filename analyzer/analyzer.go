package analyzer

import (
	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
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
	/*	esClient, err := elastic.NewESClient(config.Elastic)
		if err != nil {
			return nil, err
		}*/
	kafkaConsumer, err := kafka.NewKafkaConsumer(config.Kafka)
	if err != nil {
		return nil, err
	}
	return &LogAnalyzer{
		TopName:       config.Kafka.TopicName,
		EsClient:      nil,
		KafkaConsumer: kafkaConsumer,
		LogParser:     nil,
	}, nil
}

func (a *LogAnalyzer) Start() error {
	partition, err := a.KafkaConsumer.ConsumePartition(a.TopName, 1, sarama.OffsetOldest)
	if err != nil {
		return err
	}
	for true {
		select {
		case msg := <-partition.Messages():
			log.Infof("partition=[%d],key=[%s],value=[%s]", msg.Partition, msg.Key, msg.Value)
		}
	}
	return nil
}
