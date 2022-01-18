package elastic

import (
	"fmt"
	"github.com/yihongzhi/log-kit/analyzer/parser"
	"github.com/yihongzhi/log-kit/config"
	"testing"
	"time"
)

func TestNewESClient(t *testing.T) {
	var config = &config.ElasticConfig{
		Url:      "https://10.122.94.94:9200",
		Username: "elastic",
		Password: "xoESLDqdYh5",
	}
	client, err := NewESClient(config)
	if err != nil {
		return
	}
	data := parser.LogContent{
		Time:      time.Now(),
		Level:     "INFO",
		TxId:      "",
		SpanId:    "",
		AppId:     "demo",
		Host:      "127.0.0.1",
		ParseTime: time.Now(),
		Field: map[string]string{
			"thread": "trap-executor-0",
			"method": "c.n.d.s.r.aws.ConfigClusterResolver",
		},
		Content: "Resolving eureka endpoints via configuration",
	}
	err = client.InsertDoc("alias_log_kit", data)
	if err != nil {
		fmt.Println(err)
	}
}
