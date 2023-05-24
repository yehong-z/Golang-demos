package kafka_demo

import (
	"fmt"
	"testing"

	"github.com/Shopify/sarama"
)

func TestSendAndGetMsg(t *testing.T) {
	//SendMsg()
	GetMsg()
}

func TestTopic(t *testing.T) {
	config := sarama.NewConfig()
	config.Version = sarama.V2_1_0_0 // 设置 Kafka 协议版本
	client, err := sarama.NewClient([]string{"10.20.121.247:9092"}, config)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// 2. 获取 Kafka 中所有的主题名称
	topics, err := client.Topics()
	if err != nil {
		panic(err)
	}
	for _, t := range topics {
		fmt.Println(t)
	}
}

func TestCreateTopic(t *testing.T) {
	brokers := []string{"10.20.121.247:9092"}
	config := sarama.NewConfig()
	admin, err := sarama.NewClusterAdmin(brokers, config)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := admin.Close(); err != nil {
			fmt.Printf("Close Kafka admin failed: err=%v\n", err)
		}
	}()

	// 创建主题
	topicDetails := &sarama.TopicDetail{
		NumPartitions:     1,
		ReplicationFactor: 1,
	}
	topics := "log"
	if err := admin.CreateTopic(topics, topicDetails, false); err != nil {
		panic(err)
	}

	fmt.Println("Create Kafka topic success")

}
