package controllers

import (
	Result "classroom-system/app/common"
	"classroom-system/app/models"
	"classroom-system/app/services/homework_file"
	"github.com/gin-gonic/gin"
)

func CreateHomeworkFile(c *gin.Context) {
	var homeworkFile models.HomeworkFiles

	err := c.ShouldBindJSON(&homeworkFile)
	if err != nil {
		c.JSON(200, Result.Fail("参数错误"))
		return
	}

	err = homework_file.CreateHomeworkFile(&homeworkFile)
	if err != nil {
		c.JSON(200, Result.Fail(err.Error()))
		return
	}

	c.JSON(200, Result.Success("创建作业文件成功"))
}
