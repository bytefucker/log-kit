package kafka

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
	"github.com/yihongzhi/log-kit/config"
	"time"
)

type LogContent struct {
	Msg string `json:"msg"` //日志内容
	Ip  string `json:"ip"`  //机器IP
}

// KafkaClient kafka客户端
type KafkaClient struct {
	Client    sarama.SyncProducer
	TopicName string
}

// NewKafkaClient 初始化kafka生产者
func NewKafkaClient(config *config.KafkaConfig) (client *KafkaClient, err error) {
	conf := sarama.NewConfig()
	conf.Producer.RequiredAcks = sarama.WaitForAll
	conf.Producer.Partitioner = sarama.NewRandomPartitioner
	conf.Producer.Return.Successes = true
	conf.Producer.Timeout = 5 * time.Second

	producer, err := sarama.NewSyncProducer(config.BrokerList, conf)
	if err != nil {
		log.Error("SyncProduce create failed !", err)
		return
	}
	client = &KafkaClient{
		Client:    producer,
		TopicName: config.TopicName,
	}
	return
}

func (c *KafkaClient) SendMsg(appId string, msg LogContent) (err error) {
	text, err := json.Marshal(&msg)
	if err != nil {
		log.Debug("serilazid msg failed ", err)
		return
	}
	message := sarama.ProducerMessage{
		Topic: c.TopicName,
		Key:   sarama.StringEncoder(appId),
		Value: sarama.StringEncoder(text),
	}
	_, _, err = c.Client.SendMessage(&message)
	if err != nil {
		log.Error("kafka kafka msg failed", err)
		return
	}
	log.Debugf("send to kafka -->toppic:[%s] appId:[%v],msg:[%v]", c.TopicName, appId, string(text))
	return
}
