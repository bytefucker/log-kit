package elastic

import (
	"crypto/tls"
	es "github.com/elastic/go-elasticsearch/v7"
	log "github.com/sirupsen/logrus"
	"github.com/yihongzhi/log-kit/config"
	"net/http"
)

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
	_, err = client.Ping()
	if err != nil {
		log.Errorln("Ping es client error:", err)
		return nil, err
	}
	return &ESClient{client}, nil
}
