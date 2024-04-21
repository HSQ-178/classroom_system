package qrcode

import (
	"classroom-system/database"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/skip2/go-qrcode"
	"time"
)

// 二维码内容
type QrcodeInfo struct {
	TeacherCard     string `json:"teacherCard"`
	Grade           string `json:"grade"`
	College         string `json:"college"`
	Major           string `json:"major"`
	CourseId        int    `json:"courseId"`
	QrcodeDuration  int    `json:"qrcodeDuration"`  //二维码有效时长
	QrcodeFrequency int    `json:"qrcodeFrequency"` //二维码刷新频率
	QrcodeExpireAt  int64  `json:"qrcodeExpireAt"`  //二维码过期时间
}

// 将二维码内容存入redis
func SetRedisQrcode(qrcodeInfo *QrcodeInfo) (string, error) {
	var err error

	key := qrcodeInfo.TeacherCard
	duration := time.Duration(qrcodeInfo.QrcodeDuration) * time.Second * 60

	// 这里将二维码内容设置为 JSON 格式
	redisContent, err := json.Marshal(qrcodeInfo)
	if err != nil {
		fmt.Printf(err.Error())
		return "", errors.New("序列化错误")
	}

	// 获取 Redis 客户端
	redisClient := database.GetRedis()
	if redisClient == nil {
		fmt.Printf(err.Error())
		return "", errors.New("无法连接到 Redis")
	}

	//判断是否已存在key
	value := redisClient.Get(key).Val()
	if value != "" {
		return "", errors.New("该key已存在")
	}

	err = redisClient.Set(key, string(redisContent), duration).Err()
	if err != nil {
		fmt.Printf(err.Error())
		return "", errors.New("存入redis失败")
	}

	return key, nil
}

func GenerateQrcode(key string) ([]byte, error) {
	var qrcodeInfo QrcodeInfo
	var qrcodeTask []byte

	if key == "" {
		return []byte(""), errors.New("key不能为空")
	}

	value := database.GetRedis().Get(key).Val()
	if value == "" {
		return []byte(""), errors.New("二维码签到结束")
	}

	//将存在redis中的二维码内容解析出来
	err := json.Unmarshal([]byte(value), &qrcodeInfo)
	if err != nil {
		return []byte(""), err
	}

	// 生成二维码内容字符串
	qrcodeContent := GenerateNewQRCodeContent(qrcodeInfo)

	//生成二维码
	qrcodeData, err := qrcode.Encode(qrcodeContent, qrcode.Medium, 256)
	if err != nil {
		return []byte(""), errors.New("生成二维码失败")
	}
	qrcodeTask = qrcodeData

	// 启动定时器，定时更新二维码内容并生成新的二维码
	go func() {
		ticker := time.Tick(time.Duration(qrcodeInfo.QrcodeFrequency) * time.Second)

		for range ticker {
			// 生成二维码内容字符串
			qrcodeContent := GenerateNewQRCodeContent(qrcodeInfo)

			//生成二维码
			qrcodeData, err := qrcode.Encode(qrcodeContent, qrcode.Medium, 256)
			if err != nil {
				fmt.Printf(err.Error())
				continue
			}
			qrcodeTask = qrcodeData
		}
	}()

	return qrcodeTask, nil

}

// 生成新的二维码内容字符串
func GenerateNewQRCodeContent(qrcodeInfo QrcodeInfo) string {
	// 更新二维码过期时间
	qrcodeInfo.QrcodeExpireAt = time.Now().Unix() + int64(qrcodeInfo.QrcodeDuration*60)

	// 生成二维码内容字符串
	qrcodeContent := fmt.Sprintf("%v,%v,%v,%v,%v,%v,%v,%v,%v,%v",
		qrcodeInfo.TeacherCard, qrcodeInfo.College, qrcodeInfo.Major,
		qrcodeInfo.Grade, qrcodeInfo.CourseId, qrcodeInfo.QrcodeDuration, qrcodeInfo.QrcodeFrequency,
		qrcodeInfo.QrcodeExpireAt, time.Now())
	qrcodeContentBase64 := base64.StdEncoding.EncodeToString([]byte(qrcodeContent))

	return qrcodeContentBase64
}
