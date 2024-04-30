package handler

import (
	"classroom-system/app/models"
	"classroom-system/database"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"time"
)

type SignByQrcodeInfo struct {
	UUId        string `json:"uuid"`
	TeacherCard string `json:"teacherCard"`
	Grade       string `json:"grade"`
	College     string `json:"college"`
	Major       string `json:"major"`
	Class       string `json:"class"`
	CourseId    int    `json:"courseId"`
}

func GenerateOrUpdateRedis(messageBody []byte) {
	var qrcodeInfo models.QrcodeInfo
	var signByQrcodeInfo SignByQrcodeInfo

	err := json.Unmarshal(messageBody, &qrcodeInfo)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = json.Unmarshal([]byte(qrcodeInfo.Content), &signByQrcodeInfo)
	if err != nil {
		fmt.Println(err.Error())
	}

	signByQrcodeInfo.UUId = uuid.NewString()
	signByQrcodeInfoBytes, _ := json.Marshal(signByQrcodeInfo)
	qrcodeInfo.Content = string(signByQrcodeInfoBytes)
	qrcodeInfoBytes, _ := json.Marshal(qrcodeInfo)

	result := database.GetRedis().Set(context.Background(), signByQrcodeInfo.TeacherCard, qrcodeInfoBytes, time.Duration(qrcodeInfo.QrcodeDuration)*time.Minute)
	if result.Err() != nil {
		fmt.Println(result.Err().Error())
	}

	for {
		time.Sleep(time.Duration(qrcodeInfo.QrcodeFrequency) * time.Second)

		//判断是否已存在key
		exist, err := database.GetRedis().Exists(context.Background(), signByQrcodeInfo.TeacherCard).Result()
		if err != nil {
			fmt.Println(err.Error())
		}

		// key存在
		if exist == 1 {
			signByQrcodeInfo.UUId = uuid.NewString()
			signByQrcodeInfoBytes, _ := json.Marshal(signByQrcodeInfo)
			qrcodeInfo.Content = string(signByQrcodeInfoBytes)
			qrcodeInfoBytes, _ := json.Marshal(qrcodeInfo)

			//只更新内容，保持过期时间
			database.GetRedis().Set(context.Background(), signByQrcodeInfo.TeacherCard, qrcodeInfoBytes, redis.KeepTTL)
			fmt.Println(qrcodeInfoBytes)
		} else {
			break
		}
	}

}
