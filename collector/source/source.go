package source

// LogMessage 日志消息体
type LogMessage struct {
	AppKey string //应用ID
	Msg    string //日志消息
}

type LogSourceClient interface {
	// GetMessage 获取日志
	GetMessage() *LogMessage
}
