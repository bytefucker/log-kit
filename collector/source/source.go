package source

import "github.com/yihongzhi/log-kit/collector/task"

// LogMessage 日志消息体
type LogMessage struct {
	AppKey string //应用ID
	Msg    string //日志消息
}

type LogSource interface {
	Start()
	GetMessage() *task.LogContent
}
