package sender

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
	"github.com/yihongzhi/log-kit/config"
	"time"
)

type KafkaProducer struct {
	Client    sarama.SyncProducer
	TopicName string
}

func NewKafkaSender(config *config.KafkaConfig) (*KafkaProducer, error) {
	conf := sarama.NewConfig()
	conf.Producer.RequiredAcks = sarama.WaitForLocal
	conf.Producer.Partitioner = sarama.NewHashPartitioner
	conf.Producer.Return.Successes = true
	conf.Producer.Timeout = 5 * time.Second
	producer, err := sarama.NewSyncProducer(config.BrokerList, conf)
	if err != nil {
		log.Error("SyncProduce create failed !", err)
		return nil, err
	}
	return &KafkaProducer{
		Client:    producer,
		TopicName: config.TopicName,
	}, nil
}

// SendMessage 发送日志消息
func (d *KafkaProducer) SendMessage(message *LogMessage) error {
	text, err := json.Marshal(&message)
	if err != nil {
		log.Errorln("serialization msg failed ", err)
		return err
	}
	msg := sarama.ProducerMessage{
		Topic: d.TopicName,
		Key:   sarama.StringEncoder(message.AppId),
		Value: sarama.StringEncoder(text),
	}
	partition, offset, err := d.Client.SendMessage(&msg)
	if err != nil {
		log.Errorln("send kafka msg failed", err)
		return err
	}
	log.Debugf("send to kafka appId:[%s],toppic:[%s],partition:[%d],offset:[%d]", message.AppId, d.TopicName, partition, offset)
	return nil
}
