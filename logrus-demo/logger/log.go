package logger

import (
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
)

var logger *logrus.Logger

func init() {
	logger = logrus.New()

}

func Info() {
	logger := logrus.New()
	_, file, _, _ := runtime.Caller(1)
	println(file)
	entry := logger.WithField("file", file)
	entry.Infoln("aaa")
}

func NewLogger() *logrus.Logger {
	// 创建新的 logger
	logger := logrus.New()
	// 设置日志格式为 JSON 格式
	logger.SetFormatter(&logrus.JSONFormatter{})
	// 记录时间戳
	logger.SetReportCaller(true)
	// 设置日志级别
	logger.SetLevel(logrus.DebugLevel)
	// 输出到标准输出流
	logger.SetOutput(os.Stdout)
	// 对每个 logger 的日志添加请求信息、错误类型、错误代码、堆栈追踪等字段

	return logger
}

// NewMiddlewareLogger
// JSON 格式
// 记录时间戳,
// 设置日志级别,
// 输出到标准输出流
func NewMiddlewareLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.DebugLevel)
	logger.SetOutput(os.Stdout)
	return logger
}
