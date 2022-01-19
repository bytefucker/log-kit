package elastic

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	es "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/yihongzhi/log-kit/config"
	"github.com/yihongzhi/log-kit/logger"
	"net/http"
)

var log = logger.Log

// ESClient elastic服务
type ESClient struct {
	*es.Client
}

func NewESClient(config *config.ElasticConfig) (*ESClient, error) {
	cfg := es.Config{
		Addresses: []string{config.Url},
		Username:  config.Username,
		Password:  config.Password,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	client, err := es.NewClient(cfg)
	if err != nil {
		log.Errorln("init es client error:", err)
		return nil, err
	}
	res, err := client.Ping()
	if err != nil {
		log.Errorln("ping es error:", err)
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Errorln("connect es error:", res.String())
		return nil, errors.New(res.String())
	}
	log.Infoln("connect es success:", res.String())
	return &ESClient{client}, nil
}

//插入文档
func (c *ESClient) InsertDoc(index string, data interface{}) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}
	request := esapi.IndexRequest{
		Index: index,
		Body:  bytes.NewReader(body),
	}
	res, err := request.Do(context.Background(), c)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	log.Debugln("insert one document :", res.String())
	if http.StatusCreated != res.StatusCode {
		return errors.New(res.String())
	}
	return nil
}
