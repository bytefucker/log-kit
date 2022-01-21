package analyzer

import (
	"context"
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/yihongzhi/log-kit/analyzer/parser"
	"github.com/yihongzhi/log-kit/collector/sender"
	"github.com/yihongzhi/log-kit/config"
	"github.com/yihongzhi/log-kit/elastic"
	"github.com/yihongzhi/log-kit/kafka"
	"github.com/yihongzhi/log-kit/logger"
	"sync"
)

var log = logger.Log

// LogAnalyzer 日志解析服务
type LogAnalyzer struct {
	EsClient      *elastic.ESClient
	KafkaConsumer *kafka.Consumer
	LogParsers    map[string]parser.LogParser
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
	parsers := make(map[string]parser.LogParser)
	for _, p := range config.Analyzer.Parser {
		if p.Type == "regex" {
			parsers[p.AppId] = parser.NewRegexLogParser(p)
		} else if p.Type == "json" {
			parsers[p.AppId] = parser.NewJsonLogParser(p)
		}
	}
	return &LogAnalyzer{
		EsClient:      esClient,
		KafkaConsumer: kafkaConsumer,
		LogParsers:    parsers,
	}, nil
}

func (a *LogAnalyzer) Start() error {
	h := &logMessageHandler{
		parsers: a.LogParsers,
		client:  a.EsClient,
	}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if err := a.KafkaConsumer.Consume(context.Background(), []string{a.KafkaConsumer.TopicName}, h); err != nil {
				// 当setup失败的时候，error会返回到这里
				log.Errorf("Error from consumer: %v", err)
				return
			}
		}
	}()
	log.Infoln("consumer up and running!...")
	wg.Wait()
	return nil
}

type logMessageHandler struct {
	parsers map[string]parser.LogParser
	client  *elastic.ESClient
}

func (h *logMessageHandler) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (h *logMessageHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (h *logMessageHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		message := &sender.LogMessage{}
		err := json.Unmarshal(msg.Value, message)
		if err != nil {
			return err
		}
		logParser := h.parsers[message.AppId]
		if logParser != nil {
			content, err := logParser.Parse(message)
			if err == nil {
				h.client.InsertDoc("alias_log_kit", content)
			}
		}
		session.MarkMessage(msg, "")
	}
	claim.Messages()
	return nil
}
