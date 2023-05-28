package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func init() {
	// 创建新的 logger
	logger = logrus.New()
	// 设置日志格式为 JSON 格式
	// logger.SetFormatter(&logrus.JSONFormatter{})
	// 记录时间戳
	logger.SetReportCaller(true)
	// 设置日志级别
	logger.SetLevel(logrus.DebugLevel)
	// 输出到标准输出流
	logger.SetOutput(os.Stdout)
	// 对每个 logger 的日志添加请求信息、错误类型、错误代码、堆栈追踪等字段
	// logger.Hooks.Add(&ContextHook{})
}

func InitLog() {
}

func Info(v ...interface{}) {
	logger.Info(v...)
}
