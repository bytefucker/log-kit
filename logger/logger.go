package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

var Log *Logger

func init() {
	Log = NewLogger()
}

type Logger struct {
	*logrus.Logger
}

func NewLogger() *Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	//log.SetReportCaller(true)
	return &Logger{log}
}

func (l *Logger) SetLevel(level string) {
	Level, err := logrus.ParseLevel(level)
	if err != nil {
		fmt.Errorf("set logger level error %v", err)
		return
	}
	l.Logger.SetLevel(Level)
}
