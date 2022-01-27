package parser

import (
	"github.com/yihongzhi/log-kit/collector/sender"
	"github.com/yihongzhi/log-kit/logger"
	"time"
)

var log = logger.Log

type LogContent struct {
	Time      time.Time         `json:"time"`
	Level     string            `json:"level"`
	AppId     string            `json:"app_id"`
	Host      string            `json:"host"`
	ParseTime time.Time         `json:"parse_time"`
	Field     map[string]string `json:"field"`
	Content   string            `json:"content"`
}

// LogParser 日志解析器
type LogParser interface {
	Parse(log *sender.LogMessage) (*LogContent, error)
}
