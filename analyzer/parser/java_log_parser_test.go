package parser

import (
	"fmt"
	"regexp"
	"testing"
)

func TestJavaLogParser_Parse(t *testing.T) {
	var re = regexp.MustCompile(`(?ms)(.+)\s-\s(\w+)\s\[(.+)]\s(.*)`)
	var str = `2022-01-17 20:39:04.886 - INFO [TxId : Ignored_Trace , SpanId : ] 6 --- [ntainer#2-0-C-1] c.m.g.a.w.s.impl.AlarmHandleServiceImpl  : 处理报警开始coreId:01FSM0RHV200000000000000002899022-0_0, videoId:6ec61adc-b8e0-40e7-a46b-70c52e3420dd, 和抓拍时间差值:2644`

	for i, match := range re.FindStringSubmatch(str) {
		fmt.Println(i, "-", match)
	}
}
