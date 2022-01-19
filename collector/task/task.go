package task

import (
	"github.com/hpcloud/tail"
	"github.com/yihongzhi/log-kit/config"
	"github.com/yihongzhi/log-kit/logger"
	"github.com/yihongzhi/log-kit/metrics"
	"io"
	"regexp"
	"strings"
	"time"
)

var log = logger.Log

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
		Location:  &tail.SeekInfo{Offset: 0, Whence: io.SeekEnd}, //启动任务从最后行开始读取
		Logger:    log.Logger,
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
	log.Infof("start task: appId=[%s] path=[%s] multiline=[%+v]", t.AppId, t.LogPath, t.Multiline)
	if t.Multiline == nil || !t.Multiline.Enable {
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

func (t *TailTask) sendLog(log string) {
	if log == "" {
		return
	}
	msgObj := &LogContent{
		AppId:   t.AppId,
		Content: log,
	}
	t.msgChan <- msgObj
	metrics.ReadFileLogInc(t.AppId)
}

//单行日志
func singleLineTask(t *TailTask) {
	for true {
		select {
		case lineMsg, ok := <-t.tailFile.Lines:
			if !ok {
				log.Warnf("read obj:[%s] filed continue", t.LogPath)
				continue
			}
			t.sendLog(lineMsg.Text)
		// 任务退出
		case <-t.exitChan:
			log.Infof("task %s exit ", t.AppId)
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
				log.Warnf("read obj:[%s] filed continue", t.LogPath)
				continue
			}
			if regex.MatchString(lineMsg.Text) {
				//检测到新行，先发送缓冲区日志
				t.sendLog(buffer.String())
				buffer.Reset()
			}
			buffer.WriteString(lineMsg.Text)
		// 任务退出
		case <-t.exitChan:
			log.Infof("task %s exit ", t.AppId)
			return
		case <-time.After(5 * time.Second):
			//超过5s无新日志产生，发送缓冲区日志，防止日志最后一行读不到
			t.sendLog(buffer.String())
			buffer.Reset()
		}
	}
}
