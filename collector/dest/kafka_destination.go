package dest

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
	"github.com/yihongzhi/log-kit/config"
	"time"
)

type kafkaDestination struct {
	Client    sarama.SyncProducer
	TopicName string
}

func (d *kafkaDestination) Send(message *LogMessage) error {
	text, err := json.Marshal(&message)
	if err != nil {
		log.Debug("serilazid msg failed ", err)
		return err
	}
	msg := sarama.ProducerMessage{
		Topic: d.TopicName,
		Key:   sarama.StringEncoder(message.AppId),
		Value: sarama.StringEncoder(text),
	}
	partition, offset, err := d.Client.SendMessage(&msg)
	if err != nil {
		log.Error("kafka kafka msg failed", err)
		return err
	}
	log.Debugf("send to kafka -->toppic:[%s],partition:[%d],offset:[%d],msg:[%s]",
		d.TopicName, partition, offset, string(text))
	return nil
}

func NewKafkaDestination(config *config.KafkaConfig) (*kafkaDestination, error) {
	conf := sarama.NewConfig()
	conf.Producer.RequiredAcks = sarama.WaitForLocal
	conf.Producer.Partitioner = sarama.NewRandomPartitioner
	conf.Producer.Return.Successes = true
	conf.Producer.Timeout = 5 * time.Second
	producer, err := sarama.NewSyncProducer(config.BrokerList, conf)
	if err != nil {
		log.Error("SyncProduce create failed !", err)
		return nil, err
	}
	return &kafkaDestination{
		Client:    producer,
		TopicName: config.TopicName,
	}, nil
}
