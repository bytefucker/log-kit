package sender

import (
	"github.com/yihongzhi/log-kit/logger"
	"time"
)

var log = logger.Log

type LogMessage struct {
	Time    time.Time `json:"time"`
	Host    string    `json:"host"`
	AppId   string    `json:"app_id"`
	Content string    `json:"content"`
}

type LogMessageSender interface {
	// SendMessage  发送日志
	SendMessage(message *LogMessage) error
}
