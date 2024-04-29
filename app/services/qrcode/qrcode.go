package qrcode

import (
	Rabbitmq "classroom-system/app/middlewares/rabbitmq"
	"classroom-system/app/models"
	"classroom-system/database"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/skip2/go-qrcode"
)

// 将二维码内容存入redis
func SetRedisQrcode(qrcodeInfo *models.QrcodeInfo) error {
	var err error

	//将消息加入消息队列
	mq := Rabbitmq.RabbitMQConn("", qrcodeInfo.Id, "")

	qrcodeInfoBytes, _ := json.Marshal(qrcodeInfo)
	mq.Publish(Rabbitmq.Message{
		Type: "qrcode",
		Body: qrcodeInfoBytes,
	})

	return err
}

func GenerateQrcodeByKey(key string) ([]byte, error) {
	var qrcodeInfo models.QrcodeInfo

	if key == "" {
		return []byte(""), errors.New("key不能为空")
	}

	result := database.GetRedis().Get(context.Background(), key)
	if result.Err() != nil {
		fmt.Println(result.Err())
		return []byte(""), errors.New("二维码签到结束")
	}

	//将存在redis中的二维码内容解析出来
	err := json.Unmarshal([]byte(result.Val()), &qrcodeInfo)
	if err != nil {
		return []byte(""), err
	}

	// 生成二维码内容字符串
	qrcodeBytes, err := json.Marshal(qrcodeInfo)
	if err != nil {
		return []byte(""), err
	}

	//生成二维码
	qrcodeData, err := qrcode.Encode(string(qrcodeBytes), qrcode.Medium, 256)
	if err != nil {
		return []byte(""), errors.New("生成二维码失败")
	}

	return qrcodeData, nil
}
