package source

import (
	log "github.com/sirupsen/logrus"
	"github.com/yihongzhi/log-kit/collector/task"
	"github.com/yihongzhi/log-kit/config"
)

type fileSource struct {
	Tasks   []*task.TailTask
	MsgChan chan *task.LogContent
}

func (c *fileSource) Start() error {
	for _, task := range c.Tasks {
		go task.Start()
	}
	return nil
}

func (c *fileSource) GetMessage() *task.LogContent {
	return <-c.MsgChan
}

func NewFileSource(config *config.SourceConfig) (*fileSource, error) {
	var list []*task.TailTask
	var msgChan = make(chan *task.LogContent, config.BufferSize)
	for _, file := range config.FileSource {
		if tailTask, err := task.NewTailTask(file.AppId, file.Path, msgChan); err == nil {
			list = append(list, tailTask)
			log.Infof("init task %v", tailTask)
		}
	}
	return &fileSource{Tasks: list, MsgChan: msgChan}, nil
}
