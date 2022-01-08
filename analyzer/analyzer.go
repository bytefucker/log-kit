package analyzer

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	logs "github.com/sirupsen/logrus"
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
	handler := &logMessageHandler{}
	if err := a.KafkaConsumer.Consume(context.Background(), []string{a.TopName}, handler); err != nil {
		logs.Panicf("Error from consumer: %v", err)
	}
	logs.Println("Sarama consumer up and running!...")
	return nil
}

type logMessageHandler struct {
}

func (h *logMessageHandler) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (h *logMessageHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (h *logMessageHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("Message topic:%q partition:%d offset:%d\n", msg.Topic, msg.Partition, msg.Offset)
		// 标记消息已处理，sarama会自动提交
		session.MarkMessage(msg, "")
	}
	claim.Messages()
	return nil
}
