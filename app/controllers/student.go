package controllers

import (
	Result "classroom-system/app/common"
	"classroom-system/app/models"
	StudentService "classroom-system/app/services/student"
	"github.com/gin-gonic/gin"
)

func GetAllStudents(c *gin.Context) {
	var Students models.StudentReq

	// 将接口值转换为字符串类型
	err := c.ShouldBindJSON(&Students)
	if err != nil {
		c.JSON(200, Result.Fail("参数错误"))
		return
	}

	studentList, err := StudentService.GetStudent(&Students)
	if err != nil {
		c.JSON(200, Result.Fail(err.Error()))
		return
	}

	c.JSON(200, Result.Success(studentList))
}
