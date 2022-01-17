package elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7/esapi"
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
	content := parser.LogContent{
		Time:      time.Now(),
		Level:     "INFO",
		TxId:      "",
		SpanId:    "",
		AppId:     "demo",
		Host:      "127.0.0.1",
		ParseTime: time.Now(),
		Field: map[string]string{
			"Thread": "trap-executor-0",
			"class":  "c.n.d.s.r.aws.ConfigClusterResolver",
		},
		Content: "Resolving eureka endpoints via configuration",
	}
	body, _ := json.Marshal(content)
	request := esapi.IndexRequest{
		Index: "log_kit",
		//DocumentID: "1",
		Body: bytes.NewReader(body),
	}
	res, err := request.Do(context.Background(), client)
	if err != nil {
		return
	}
	defer res.Body.Close()
	fmt.Println(res.String())
}
