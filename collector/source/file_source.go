package source

import (
	log "github.com/sirupsen/logrus"
	"github.com/yihongzhi/log-kit/collector/task"
	"github.com/yihongzhi/log-kit/config"
)

type fileSource struct {
	Tasks []*task.TailTask
}

func (c *fileSource) GetMessage() *task.LogContent {
	return nil
}

func (c *fileSource) Start() {
	for _, task := range c.Tasks {
		go task.Start()
	}
}

func NewFileSource(config *config.SourceConfig) (*fileSource, error) {
	var list []*task.TailTask
	for _, file := range config.FileSource {
		if tailTask, err := task.NewTailTask(file.AppId, file.Path, config.BufferSize); err == nil {
			list = append(list, tailTask)
			log.Infof("init task %v", tailTask)
		}
	}
	return &fileSource{Tasks: list}, nil
}
