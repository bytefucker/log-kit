package kafka

import (
	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
	"github.com/yihongzhi/log-kit/config"
	"time"
)

type Consumer struct {
	TopicName string
	sarama.ConsumerGroup
}

type Producer struct {
	TopicName string
	GroupId   string
	sarama.SyncProducer
}

func NewKafkaConsumer(config *config.KafkaConfig) (*Consumer, error) {
	cfg := sarama.NewConfig()
	cfg.Consumer.Offsets.AutoCommit.Interval = 1 * time.Second
	cfg.Consumer.Offsets.AutoCommit.Enable = true
	cfg.Consumer.Offsets.Initial = sarama.OffsetOldest
	cfg.Consumer.Offsets.Retry.Max = 3
	consumer, err := sarama.NewConsumerGroup(config.BrokerList, config.GroupId, cfg)
	if err != nil {
		log.Errorln("init kafka consumer error", err)
		return nil, err
	}
	return &Consumer{
		TopicName:     config.TopicName,
		ConsumerGroup: consumer,
	}, nil
}

func NewKafkaProducer(config *config.KafkaConfig) (*Producer, error) {
	cfg := sarama.NewConfig()
	cfg.Producer.RequiredAcks = sarama.WaitForLocal
	cfg.Producer.Partitioner = sarama.NewHashPartitioner
	cfg.Producer.Return.Successes = true
	cfg.Producer.Timeout = 5 * time.Second
	producer, err := sarama.NewSyncProducer(config.BrokerList, cfg)
	if err != nil {
		log.Errorln("init kafka producer producer", err)
		return nil, err
	}
	return &Producer{
		TopicName:    config.TopicName,
		GroupId:      config.GroupId,
		SyncProducer: producer,
	}, nil
}
