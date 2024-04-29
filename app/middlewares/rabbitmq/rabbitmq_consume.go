package rabbitmq

import (
	"classroom-system/app/middlewares/rabbitmq/handler"
	"encoding/json"
	"fmt"
)

func (r *RabbitMQ) Consume() {
	// 1. 申请队列，如果队列不存在会自动创建，如何存在则跳过创建
	// 保证队列存在，消息能发送到队列中
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		false, // 是否持久化
		false, // 是否自动删除
		false, // 是否具有排他性
		false, // 是否阻塞
		nil,   // 额外属性
	)
	if err != nil {
		fmt.Println(err)
	}

	//接受消息
	msgs, err := r.channel.Consume(
		r.QueueName,
		"",    // 用来区分多个消费者
		false, // 是否自动应答
		false, // 是否具有排他性
		false, // 如果设置为true,表示不能将同一个connection中发送的消息传递给这个connection中的消费者
		false, // 队列是否阻塞
		nil,   // 排他性
	)
	if err != nil {
		fmt.Println(err)
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			var message Message
			err := json.Unmarshal(d.Body, &message)
			if err != nil {
				fmt.Println(err)
			}

			switch message.Type {
			case "qrcode":
				go handler.GenerateOrUpdateRedis(message.Body)
			}

			d.Ack(false)
		}
	}()
	<-forever
}
