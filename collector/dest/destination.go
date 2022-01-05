package dest

import "time"

type LogMessage struct {
	Time    time.Time `json:"time"`
	Host    string    `json:"host"`
	AppId   string    `json:"app_id"`
	Content string    `json:"content"`
}

type LogDestination interface {
	// Send 发送日志
	Send(message *LogMessage) error
}
