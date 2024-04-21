package notice

import (
	"classroom-system/database"
	"errors"
	"fmt"
	"time"
)

type Notice struct {
	Grade       string        `json:"grade"`       // 年级
	Major       string        `json:"major"`       // 专业
	College     string        `json:"college"`     // 学院
	CourseId    int           `json:"courseId"`    // 课程id
	Duration    time.Duration `json:"duration"`    // 时长
	TeacherCard string        `json:"teacherCard"` // 教师编号
}

func check(n *Notice) error {
	if n.TeacherCard == "" {
		return errors.New("教师编号不能为空")
	}
	if n.Grade == "" {
		return errors.New("年级不能为空")
	}
	if n.Major == "" {
		return errors.New("专业不能为空")
	}
	if n.College == "" {
		return errors.New("学院不能为空")
	}
	if n.CourseId == 0 {
		return errors.New("课程编号不能为空")
	}
	if n.Duration == 0 {
		return errors.New("时长不能为空")
	}
	return nil

}

func SetNoticeRedis(n *Notice) (string, error) {

	err := check(n)
	if err != nil {
		return "", err
	}

	key := n.TeacherCard
	duration := 60 * n.Duration * time.Second

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

	err = redisClient.Set(key, "127.0.0.1", duration).Err()
	if err != nil {
		fmt.Printf(err.Error())
		return "", errors.New("存入redis失败")
	}
	return key, nil
}

func GetNoticeRedis(key string) (string, error) {
	value := database.GetRedis().Get(key).Val()
	if value == "" {
		return "", errors.New("该key不存在")
	}

	return value, nil
}
