package sender

import "time"

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
