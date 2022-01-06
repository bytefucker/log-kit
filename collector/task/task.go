package task

import (
	"github.com/hpcloud/tail"
	logs "github.com/sirupsen/logrus"
)

// TailTask Tail 任务
type TailTask struct {
	AppId    string
	LogPath  string
	tailFile *tail.Tail
	msgChan  chan<- *LogContent
	exitChan chan int
}

// LogContent 日志消息体
type LogContent struct {
	AppId   string
	Content string
}

func NewTailTask(appId string, path string, msgChan chan<- *LogContent) (*TailTask, error) {
	tailFile, err := tail.TailFile(path, tail.Config{
		ReOpen:    true,
		Follow:    true,
		MustExist: false,
		Poll:      true,
	})
	if err != nil {
		return nil, err
	}
	return &TailTask{
		AppId:    appId,
		LogPath:  path,
		tailFile: tailFile,
		msgChan:  msgChan,
		exitChan: make(chan int),
	}, nil
}

// Start 开始读取日志
func (t *TailTask) Start() {
	logs.Infof("Start Task: appId=%s path=%s", t.AppId, t.LogPath)
	for true {
		select {
		case lineMsg, ok := <-t.tailFile.Lines:
			if !ok {
				logs.Warnf("read obj:[%s] filed continue", t.LogPath)
				continue
			}
			// 消息为空跳过
			if lineMsg.Text == "" {
				continue
			}
			msgObj := &LogContent{
				AppId:   t.AppId,
				Content: lineMsg.Text,
			}
			t.msgChan <- msgObj
		// 任务退出
		case <-t.exitChan:
			logs.Infof("task %s exit ", t.AppId)
			return
		}
	}
}

// Stop 停止读取日志
func (t *TailTask) Stop() {
	t.exitChan <- 1
	close(t.exitChan)
}
