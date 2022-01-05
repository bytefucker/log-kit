package task

import (
	"github.com/hpcloud/tail"
	logs "github.com/sirupsen/logrus"
)

// TailTask Tail 任务
type TailTask struct {
	AppId    string
	LogPath  string
	MsgChan  chan *LogContent
	tailFile *tail.Tail
	exitChan chan int
}

// LogContent 日志消息体
type LogContent struct {
	AppId   string
	Content string
}

func NewTailTask(appId string, path string, bufferSize int32) (*TailTask, error) {
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
		MsgChan:  make(chan *LogContent, bufferSize),
		exitChan: make(chan int, 1),
	}, nil
}

// Start 开始读取日志
func (t *TailTask) Start() {
	for true {
		select {
		case lineMsg, ok := <-t.tailFile.Lines:
			if !ok {
				logs.Warnf("read obj:[%v] topic:[%v] filed continue", t, t.AppId)
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
			t.MsgChan <- msgObj
		// 任务退出
		case <-t.exitChan:
			logs.Infof("task [%v] exit ", t)
			return
		}
	}
}

// Stop 停止读取日志
func (t *TailTask) Stop() {
	t.exitChan <- 1
}
