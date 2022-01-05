package task

import (
	"github.com/hpcloud/tail"
	logs "github.com/sirupsen/logrus"
	"github.com/yihongzhi/log-kit/collector/source"
)

//Tail 任务
type TailTask struct {
	TaskInfo *TailTaskInfo
	TailObj  *tail.Tail
	MsgChan  chan *source.LogMessage
	ExitChan chan int
}

//任务详情
type TailTaskInfo struct {
	AppKey  string `json:"appKey"`  //应用id
	LogPath string `json:"logPath"` //日志路径
}

func NewTailTask(task *TailTaskInfo, msgChan chan *source.LogMessage) *TailTask {
	tailObj, err := tail.TailFile(task.LogPath, tail.Config{
		ReOpen:    true,
		Follow:    true,
		MustExist: false,
		Poll:      true,
	})
	if err != nil {
		logs.Warnf("task [%v] create failed, %v", task, err)
		return nil
	}
	tailTask := &TailTask{
		TailObj:  tailObj,
		TaskInfo: task,
		MsgChan:  msgChan,
		ExitChan: make(chan int, 1),
	}
	return tailTask
}

// Start 开始读取日子文件
func (t *TailTask) Start() {
	for true {
		select {
		case lineMsg, ok := <-t.TailObj.Lines:
			if !ok {
				logs.Warnf("read obj:[%v] topic:[%v] filed continue", t, t.TaskInfo.AppKey)
				continue
			}
			// 消息为空跳过
			if lineMsg.Text == "" {
				continue
			}
			msgObj := &source.LogMessage{
				AppKey: t.TaskInfo.AppKey,
				Msg:    lineMsg.Text,
			}
			t.MsgChan <- msgObj
		// 任务退出
		case <-t.ExitChan:
			logs.Infof("task [%v] exit ", t.TaskInfo)
			return
		}
	}
}
