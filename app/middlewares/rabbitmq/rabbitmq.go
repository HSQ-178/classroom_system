package rabbitmq

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	conn         *amqp.Connection // 连接
	channel      *amqp.Channel    // 通道
	ExchangeName string           // 交换机名称
	QueueName    string           // 队列名称
	RouteKey     string           // 路由键
	Mqurl        string           // 连接信息
}

type Message struct {
	Type string `json:"type"`
	Body []byte `json:"body"`
}

func RabbitMQConn(exchangeName, queueName, routeKey string) *RabbitMQ {
	var err error

	rabbitmq := &RabbitMQ{
		ExchangeName: exchangeName,
		QueueName:    queueName,
		RouteKey:     routeKey,
		Mqurl:        "amqp://root:08181029hsq@localhost:5672/",
	}

	// 1. 建立连接
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	if err != nil {
		fmt.Println("建立连接失败", err)
		panic(err)
	}
	// 2. 获取通道
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	if err != nil {
		fmt.Println("获取通道失败", err)
		panic(err)
	}
	return rabbitmq
}
