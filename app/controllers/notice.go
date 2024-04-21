package controllers

import (
	Result "classroom-system/app/common"
	Notices "classroom-system/app/services/notice"
	"github.com/gin-gonic/gin"
)

func SetNoticeRedis(c *gin.Context) {
	var notice Notices.Notice

	err := c.ShouldBindJSON(&notice)
	if err != nil {
		c.JSON(200, Result.Fail("参数错误"))
		return
	}
	noticeRedis, err := Notices.SetNoticeRedis(&notice)
	if err != nil {
		c.JSON(200, Result.Fail(err.Error()))
		return
	}
	c.JSON(200, Result.Success(noticeRedis))
}

func GetNoticeRedis(c *gin.Context) {
	// 从请求体中绑定 JSON 数据到一个映射，期望参数名为 "teacherCard"
	var requestData struct {
		TeacherCard string `json:"teacherCard"`
	}

	// 使用 ShouldBindJSON 将请求体中的 JSON 数据绑定到 requestData 结构体
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(200, Result.Fail("参数错误"))
		return
	}

	// 获取 requestData 中的 teacherCard 作为 key
	key := requestData.TeacherCard

	// 调用 GetNoticeRedis 方法从 Redis 中获取数据
	value, err := Notices.GetNoticeRedis(key)
	if err != nil {
		c.JSON(200, Result.Fail(err.Error()))
		return
	}

	// 返回获取到的 value 数据
	c.JSON(200, Result.Success(value))
}
