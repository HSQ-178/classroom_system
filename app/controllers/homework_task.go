package controllers

import (
	Result "classroom-system/app/common"
	"classroom-system/app/models"
	homeworkTaskService "classroom-system/app/services/homework_task"
	"github.com/gin-gonic/gin"
)

func CreateHomeworkTask(c *gin.Context) {
	var homeworkTask models.HomeworkTask

	err := c.ShouldBindJSON(&homeworkTask)
	if err != nil {
		c.JSON(200, Result.Fail("参数错误"))
		return
	}

	id, err := homeworkTaskService.CreateHomeworkTask(&homeworkTask)
	if err != nil {
		c.JSON(200, Result.Fail(err.Error()))
		return
	}

	c.JSON(200, Result.Success(id))
}

func ListHomeworkTask(c *gin.Context) {
	var req models.HomeworkTaskReq

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(200, Result.Fail("参数错误"))
		return
	}

	homeworkTaskList, err := homeworkTaskService.ListHomeworkTask(&req)
	if err != nil {
		c.JSON(200, Result.Fail(err.Error()))
		return
	}

	c.JSON(200, Result.Success(homeworkTaskList))
}
