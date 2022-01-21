package parser

import (
	"github.com/yihongzhi/log-kit/collector/sender"
	"github.com/yihongzhi/log-kit/config"
)

type JsonLogParser struct {
}

func (p *JsonLogParser) Parse(log *sender.LogMessage) (*LogContent, error) {
	panic("implement me")
}

func NewJsonLogParser(config *config.LogParserConfig) *JsonLogParser {
	return &JsonLogParser{}
}
