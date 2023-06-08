package kafkalogger

import (
	"io"
	"log"
	"os"
	"testing"
	"time"
)

func TestLog(t *testing.T) {
	brokers := []string{"121.36.89.81:9092"}
	kafkaWriter, err := NewKafkaWriter(brokers, "log")
	if err != nil {
		log.Fatalf("failed to create kafka writer: %v", err)
	}
	defer kafkaWriter.Close()

	// 设置output,默认为stderr,可以为任何io.Writer，比如文件*os.File
	file, err := os.OpenFile("logfile.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
	writers := []io.Writer{
		file,
		os.Stdout,
		kafkaWriter,
	}
	multiWriter := io.MultiWriter(writers...)
	InitLogger(multiWriter)
	for i := 0; i < 1000; i++ {
		Logger.Info("test log", time.Now())
		time.Sleep(time.Millisecond * 100)
	}
}
