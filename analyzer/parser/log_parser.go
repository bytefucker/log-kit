package parser

import "time"

type LogContent struct {
	Time      time.Time
	Level     string
	TxId      string
	SpanId    string
	AppId     string
	Host      string
	ParseTime time.Time
	Field     map[string]string
	Content   string
}

// LogParser 日志解析器
type LogParser interface {
	Parse() (*LogContent, error)
}
