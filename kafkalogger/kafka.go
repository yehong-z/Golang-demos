package kafkalogger

import (
	"fmt"

	"github.com/Shopify/sarama"
)

type KafkaWriter struct {
	p     sarama.SyncProducer
	topic string
}

func NewKafkaWriter(brokers []string, topic string) (*KafkaWriter, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Version = sarama.V2_4_0_0 // 设置 Kafka 协议版本

	// 连接 Kafka 集群
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		fmt.Println("-----------------------------------")
		panic(err)
	}

	return &KafkaWriter{
		p:     producer,
		topic: topic,
	}, nil
}

// 实现 io.Writer 接口
func (kw *KafkaWriter) Write(p []byte) (int, error) {
	n := len(p)
	if n == 0 {
		return 0, nil
	}

	// 构建消息
	message := &sarama.ProducerMessage{Topic: kw.topic, Value: sarama.StringEncoder(p)}

	// 发送消息
	partition, offset, err := kw.p.SendMessage(message)
	if err != nil {
		fmt.Printf("Send message to kafka failed: err=%v\n", err)
		return 0, err
	}
	fmt.Printf("Send message to kafka success, partition=%d, offset=%d\n", partition, offset)

	return n, nil
}

// 关闭 Kafka Writer
func (kw *KafkaWriter) Close() error {
	return kw.p.Close()
}
