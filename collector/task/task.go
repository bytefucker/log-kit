package task

import (
	"github.com/hpcloud/tail"
	logs "github.com/sirupsen/logrus"
	"github.com/yihongzhi/log-kit/config"
	"regexp"
	"strings"
)

// TailTask Tail 任务
type TailTask struct {
	AppId     string
	LogPath   string
	Multiline *config.Multiline
	tailFile  *tail.Tail
	msgChan   chan<- *LogContent
	exitChan  chan int
}

// LogContent 日志消息体
type LogContent struct {
	AppId   string
	Content string
}

func NewTailTask(source *config.LogFileSource, msgChan chan<- *LogContent) (*TailTask, error) {
	tailFile, err := tail.TailFile(source.Path, tail.Config{
		ReOpen:    true,
		Follow:    true,
		MustExist: false,
		Poll:      true,
	})
	if err != nil {
		return nil, err
	}
	return &TailTask{
		AppId:     source.AppId,
		LogPath:   source.Path,
		Multiline: source.Multiline,
		tailFile:  tailFile,
		msgChan:   msgChan,
		exitChan:  make(chan int),
	}, nil
}

// Start 开始读取日志
func (t *TailTask) Start() {
	logs.Infof("Start Task: appId=%s path=%s", t.AppId, t.LogPath)
	if t.Multiline == nil || t.Multiline.Pattern == "" {
		singleLineTask(t)
	} else {
		multilineTask(t)
	}
}

// Stop 停止读取日志
func (t *TailTask) Stop() {
	t.exitChan <- 1
	close(t.exitChan)
}

//单行日志
func singleLineTask(t *TailTask) {
	for true {
		select {
		case lineMsg, ok := <-t.tailFile.Lines:
			if !ok {
				logs.Warnf("read obj:[%s] filed continue", t.LogPath)
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

//多行日志任务
func multilineTask(t *TailTask) {
	regex, _ := regexp.Compile(t.Multiline.Pattern)
	var buffer strings.Builder
	for true {
		select {
		case lineMsg, ok := <-t.tailFile.Lines:
			if !ok {
				logs.Warnf("read obj:[%s] filed continue", t.LogPath)
				continue
			}
			if regex.MatchString(lineMsg.Text) {
				//如果新的日志从日期开始，先发送缓冲区的数据
				if buffer.Len() > 0 {
					msgObj := &LogContent{
						AppId:   t.AppId,
						Content: buffer.String(),
					}
					t.msgChan <- msgObj
				}
				buffer.Reset()
			}
			buffer.WriteString(lineMsg.Text)
		// 任务退出
		case <-t.exitChan:
			logs.Infof("task %s exit ", t.AppId)
			return
		}
	}
}
