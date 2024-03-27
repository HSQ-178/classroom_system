package controllers

import (
	Result "classroom-system/app/common"
	"classroom-system/app/models"
	CourseService "classroom-system/app/services/course"
	"github.com/gin-gonic/gin"
)

// GerAllCourse 获取该老师的所有课程
func GerAllCourse(c *gin.Context) {
	var courseReq models.CourseReq

	err := c.ShouldBindJSON(&courseReq)
	if err != nil {
		c.JSON(400, Result.Fail("参数错误"))
		return
	}

	courseList, err := CourseService.GetAllCourse(&courseReq)
	if err != nil {
		c.JSON(400, Result.Fail(err.Error()))
		return
	}

	c.JSON(200, Result.Success(courseList))
}
