package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	// 连接 RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@10.20.121.247:5672/")
	if err != nil {
		log.Fatalf("Failed to connect RabbitMQ: %v", err)
	}

	// 创建 channel
	channel, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}

	// 声明 queue
	queueName := "hello"
	queue, err := channel.QueueDeclare(
		queueName, // 名称
		false,     // durable，是否持久化
		false,     // delete when unused，当队列没有消费者时是否自动删除
		false,     // exclusive，队列是否为私有队列
		false,     // no-wait，是否等待服务器确认
		nil,       // arguments，队列参数
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	// 发送消息
	body := "Hello, RabbitMQ!"
	err = channel.Publish(
		"",         // exchange，交换机名称
		queue.Name, // routing key，路由键
		false,      // mandatary，是否要求服务器确认
		false,      // immediate，是否立即路由消息
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	if err != nil {
		log.Fatalf("Failed to publish a message: %v", err)
	}

	fmt.Printf("Sent message: %s\n", body)

	// 消费消息
	messages, err := channel.Consume(
		queue.Name, // queue，队列名称
		"",         // consumer，消费者名称
		true,       // auto-ack，是否自动确认
		false,      // exclusive，队列是否为私有队列
		false,      // no-local，是否不接收自己发送的消息
		false,      // no-wait，是否等待服务器确认
		nil,        // arguments，消费者参数
	)
	if err != nil {
		log.Fatalf("Failed to consume messages: %v", err)
	}

	for message := range messages {
		fmt.Println("Received message:", string(message.Body))
	}

	// 关闭 channel 和 conn
	err = channel.Close()
	if err != nil {
		return
	}
	err = conn.Close()
	if err != nil {
		return
	}
}
