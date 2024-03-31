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
		c.JSON(200, Result.Fail(err.Error()))
		return
	}

	notSignInList, err := RecordService.GetNotSignInStudentList(&notSignInReq)
	if err != nil {
		c.JSON(200, Result.Fail(err.Error()))
		return
	}

	c.JSON(200, Result.Success(notSignInList))

}

// GetSignInAndAbsenceList 获取签到与缺勤
func GetSignInAndAbsenceList(c *gin.Context) {
	var signInAndAbsenceReq models.RecordReq

	err := c.ShouldBindJSON(&signInAndAbsenceReq)
	if err != nil {
		c.JSON(200, Result.Fail(err.Error()))
		return
	}

	signInAndAbsenceList, err := RecordService.GetSignInAndAbsence(&signInAndAbsenceReq)
	if err != nil {
		c.JSON(200, Result.Fail(err.Error()))
		return
	}

	c.JSON(200, Result.Success(signInAndAbsenceList))
}
