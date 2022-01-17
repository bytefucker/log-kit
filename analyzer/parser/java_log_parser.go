package parser

import (
	"fmt"
	"github.com/yihongzhi/log-kit/collector/sender"
	"regexp"
	"strings"
	"time"
)

type JavaLogParser struct {
	regx     *regexp.Regexp
	fieldMap map[string]string
}

func NewJavaLogParser() *JavaLogParser {
	re := regexp.MustCompile(`(?ms)(.+)\s-\s(\w+)\s\[TxId\s:(.+),\sSpanId\s:(.+)].+\[(.+)]\s(\S+)\s\s:\s(.+)`)
	return &JavaLogParser{
		regx: re,
		fieldMap: map[string]string{
			"": "",
		},
	}
}

func (p *JavaLogParser) Parse(log *sender.LogMessage) (*LogContent, error) {
	matches := p.regx.FindStringSubmatch(log.Content)
	if matches == nil {
		return nil, nil
	}
	fmt.Println("------------------------------")
	for i, match := range matches {
		fmt.Println(i, "->", strings.TrimSpace(match))
	}
	return &LogContent{
		Time:      time.Now(),
		Level:     "",
		TxId:      "",
		SpanId:    "",
		AppId:     "",
		Host:      "",
		ParseTime: time.Now(),
		Field:     nil,
		Content:   "",
	}, nil
}
