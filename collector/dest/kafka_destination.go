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

func (d *kafkaDestination) Send(content *LogMessage) error {
	text, err := json.Marshal(&content)
	if err != nil {
		log.Debug("serilazid msg failed ", err)
		return err
	}
	message := sarama.ProducerMessage{
		Topic: d.TopicName,
		Key:   sarama.StringEncoder(content.AppId),
		Value: sarama.StringEncoder(text),
	}
	_, _, err = d.Client.SendMessage(&message)
	if err != nil {
		log.Error("kafka kafka msg failed", err)
		return err
	}
	log.Debugf("send to kafka -->toppic:[%s],msg:[%v]", d.TopicName, string(text))
	return nil
}

func NewKafkaDestination(config *config.KafkaConfig) (*kafkaDestination, error) {
	conf := sarama.NewConfig()
	conf.Producer.RequiredAcks = sarama.WaitForAll
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
