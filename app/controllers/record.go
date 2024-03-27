package controllers

import (
	Result "classroom-system/app/common"
	"classroom-system/app/models"
	RecordService "classroom-system/app/services/record"
	"github.com/gin-gonic/gin"
)

// GetNotSignInList 获取未签到学生
func GetNotSignInList(c *gin.Context) {
	var notSignInReq models.RecordReq

	err := c.ShouldBindJSON(&notSignInReq)
	if err != nil {
		c.JSON(400, Result.Fail(err.Error()))
		return
	}

	notSignInList, err := RecordService.GetNotSignInStudentList(&notSignInReq)
	if err != nil {
		c.JSON(400, Result.Fail(err.Error()))
		return
	}

	c.JSON(200, Result.Success(notSignInList))

}
