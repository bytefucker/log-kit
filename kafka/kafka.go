package kafka

import (
	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
	"github.com/yihongzhi/log-kit/config"
	"time"
)

type Consumer struct {
	sarama.ConsumerGroup
}

type Producer struct {
	sarama.SyncProducer
}

func NewKafkaConsumer(config *config.KafkaConfig) (*Consumer, error) {
	cfg := sarama.NewConfig()
	consumer, err := sarama.NewConsumerGroup(config.BrokerList, config.GroupId, cfg)
	if err != nil {
		log.Errorln("init kafka consumer error", err)
		return nil, err
	}
	return &Consumer{consumer}, nil
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
	return &Producer{producer}, nil
}
