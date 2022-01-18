package parser

import (
	"errors"
	"github.com/yihongzhi/log-kit/collector/sender"
	"regexp"
	"strings"
	"time"
)

type JavaLogParser struct {
	regx   *regexp.Regexp
	fields []string
}

func NewJavaLogParser() *JavaLogParser {
	var re = regexp.MustCompile(`(?ms)(.+)\s-\s(\w+)\s\[TxId\s:(.+),\sSpanId\s:(.+)].+\[(.+)]\s(\S+)\s+:\s(.+)`)
	return &JavaLogParser{
		regx:   re,
		fields: []string{"time", "level", "tx_id", "span_id", "thread", "method", "content"},
	}
}

func (p *JavaLogParser) Parse(logMessage *sender.LogMessage) (*LogContent, error) {
	matches := p.regx.FindStringSubmatch(logMessage.Content)
	if matches == nil || len(matches) != len(p.fields)+1 {
		return nil, errors.New("matches failed")
	}
	fieldMap := make(map[string]string, len(p.fields))
	log.Debugln("------------------------------")
	for i, match := range matches {
		log.Debugln(i, "->", strings.TrimSpace(match))
		if i > 0 {
			fieldMap[p.fields[i-1]] = match
		}
	}
	parse, _ := time.Parse("2006-01-02 15:04:05.000", fieldMap["time"])
	return &LogContent{
		Time:      parse,
		Level:     fieldMap["level"],
		TxId:      fieldMap["tx_id"],
		SpanId:    fieldMap["span_id"],
		AppId:     logMessage.AppId,
		Host:      logMessage.Host,
		ParseTime: time.Now(),
		Field: map[string]string{
			"thread": fieldMap["thread"],
			"method": fieldMap["method"],
		},
		Content: fieldMap["content"],
	}, nil
}
