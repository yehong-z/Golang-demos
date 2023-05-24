package kafka_demo

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func SendMsg() {
	topic := "test-topic"
	partition := 0
	conn, err := kafka.DialLeader(context.Background(), "tcp", "10.20.121.247:9092", topic, partition)
	if err != nil {
		panic(fmt.Sprintf("failed to dial leader: %v", err))
	}
	defer conn.Close()

	msg := kafka.Message{
		Key:   []byte("key"),
		Value: []byte("hello kafka-go"),
	}
	_, err = conn.WriteMessages(msg)
	if err != nil {
		panic(fmt.Sprintf("failed to write message: %v", err))
	}
}
