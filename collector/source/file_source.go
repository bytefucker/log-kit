package source

import (
	"github.com/yihongzhi/log-kit/collector/task"
	"github.com/yihongzhi/log-kit/config"
)

type fileSource struct {
	tasks   []*task.TailTask
	msgChan chan *task.LogContent
}

func NewFileSource(config *config.SourceConfig, bufferSize int32) (*fileSource, error) {
	var list []*task.TailTask
	var msgChan = make(chan *task.LogContent, bufferSize)
	for _, file := range config.FileSource {
		if tailTask, err := task.NewTailTask(file, msgChan); err == nil {
			list = append(list, tailTask)
			log.Infof("Init TailTask: app_id=%s path=%s", tailTask.AppId, tailTask.LogPath)
		}
	}
	return &fileSource{tasks: list, msgChan: msgChan}, nil
}

func (c *fileSource) Start() error {
	for _, task := range c.tasks {
		go task.Start()
	}
	return nil
}

func (c *fileSource) GetMessage() *task.LogContent {
	return <-c.msgChan
}
