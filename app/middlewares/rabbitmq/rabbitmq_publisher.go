package rabbitmq

import (
	"classroom-system/app/services/qrcode"
	"context"
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func (r *RabbitMQ) Publish(qrcodeInfo *qrcode.QrcodeInfo) {
	qrcodeInfoBytes, err := json.Marshal(qrcodeInfo)
	if err != nil {
		log.Println("Error marshalling QrcodeInfo to bytes:", err)
		return
	}

	// 1. 申请队列，如果队列不存在会自动创建，如何存在则跳过创建
	// 保证队列存在，消息能发送到队列中
	_, err = r.channel.QueueDeclare(
		r.QueueName, // 队列名
		false,       // 是否持久化
		false,       // 是否自动删除
		false,       // 是否具有排他性
		false,       // 是否阻塞
		nil,         // 额外属性
	)
	if err != nil {
		fmt.Println(err)
	}

	//发送消息到队列中
	r.channel.PublishWithContext(
		context.Background(),
		r.ExchangeName,
		r.QueueName,
		false, // 如果为true, 会根据exchange类型和routkey规则，如果无法找到符合条件的队列那么会把发送的消息返回给发送者
		false, // 如果为true, 当exchange发送消息到队列后发现队列上没有绑定消费者，则会把消息发还给发送者
		amqp.Publishing{
			//ContentType: "text/plain",
			ContentType: "application/json",
			Body:        qrcodeInfoBytes,
		},
	)
}
