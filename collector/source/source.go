package source

import (
	"github.com/yihongzhi/log-kit/collector/task"
	"github.com/yihongzhi/log-kit/logger"
)

var log = logger.Log

// LogMessage 日志消息体
type LogMessage struct {
	AppKey string //应用ID
	Msg    string //日志消息
}

type LogSource interface {
	Start() error
	GetMessage() *task.LogContent
}
