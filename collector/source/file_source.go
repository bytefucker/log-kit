package source

import (
	"github.com/yihongzhi/log-kit/collector/task"
	"github.com/yihongzhi/log-kit/config"
)

const chanSize = 10000

type FileSourceClient struct {
	MsgChan chan *LogMessage
	Tasks   []*task.TailTask
}

func (c *FileSourceClient) GetMessage() *LogMessage {
	return <-c.MsgChan
}

func NewSourceClient(config *config.SourceConfig) (*LogSourceClient, error) {
	return nil, nil
}
