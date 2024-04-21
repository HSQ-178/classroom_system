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

// AddRecords 添加记录
func AddRecords(c *gin.Context) {
	var record models.Record

	err := c.ShouldBindJSON(&record)
	if err != nil {
		c.JSON(200, Result.Fail(err.Error()))
		return
	}

	err = RecordService.AddRecords(&record)
	if err != nil {
		c.JSON(200, Result.Fail(err.Error()))
		return
	}
	c.JSON(200, Result.Success("添加记录成功"))
}

// DeleteRecords 删除记录
func DeleteRecords(c *gin.Context) {
	var records models.Record

	err := c.ShouldBindJSON(&records)
	if err != nil {
		c.JSON(200, Result.Fail(err.Error()))
		return
	}
	err = RecordService.DeleteRecords(&records)
	if err != nil {
		c.JSON(200, Result.Fail(err.Error()))
		return
	}
	c.JSON(200, Result.Success("删除记录成功"))
}

// UpdateRecords 更新记录
// 签到 => 缺勤，缺勤 => 签到
func UpdateRecords(c *gin.Context) {
	var record models.Record

	err := c.ShouldBindJSON(&record)
	if err != nil {
		c.JSON(200, Result.Fail(err.Error()))
		return
	}

	err = RecordService.UpdateRecords(&record)
	if err != nil {
		c.JSON(200, Result.Fail(err.Error()))
		return
	}
	c.JSON(200, Result.Success("更新记录成功"))
}
