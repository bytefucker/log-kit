package dest

import "time"

type LogMessage struct {
	Time    time.Duration `json:"time"`
	Host    string        `json:"host"`
	AppId   string        `json:"app_id"`
	content string        `json:"content"`
}

type LogDestination interface {
	// Send 发送日志
	Send(content *LogMessage) error
}
