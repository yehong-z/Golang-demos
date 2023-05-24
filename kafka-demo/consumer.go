package kafka_demo

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

// 这里可能有主机名无法解析的问题，修改本机的hosts文件暂时解决这个问题

func GetMsg() {
	topic := "test-topic"
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"10.20.121.247:9092"},
		Topic:   topic,
	})
	defer r.Close()

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		fmt.Printf("received message: %s\n", string(m.Value))
	}
}
