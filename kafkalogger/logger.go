package kafkalogger

import (
	"github.com/sirupsen/logrus"
	"io"
)

var Logger *logrus.Logger

func InitLogger(w io.Writer) {
	Logger = logrus.New()
	// 设置日志格式为 JSON 格式
	Logger.SetFormatter(&logrus.JSONFormatter{})
	Logger.SetReportCaller(true)
	Logger.SetLevel(logrus.DebugLevel)
	Logger.SetOutput(w)

}
