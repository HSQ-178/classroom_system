package controllers

import (
	Result "classroom-system/app/common"
	"classroom-system/app/models"
	teacherService "classroom-system/app/services/teacher"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var registerReq models.TeachersRegisterReq //存储用户注册请求中的数据

	err := c.ShouldBindJSON(&registerReq) //将请求中的 JSON 数据绑定到 registerReq 结构体中
	if err != nil {
		c.JSON(200, Result.Fail("参数错误"))
		return
	}

	err = teacherService.TeacherRegister(&registerReq)
	if err != nil {
		c.JSON(200, Result.Fail(err.Error()))
		return
	}

	c.JSON(200, Result.Success("注册成功"))
}
