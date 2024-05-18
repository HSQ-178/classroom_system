package controllers

import (
	Result "classroom-system/app/common"
	"classroom-system/app/models"
	AuthcodeService "classroom-system/app/services/authcode"
	"fmt"
	"github.com/gin-gonic/gin"
)

func AddAuthcode(c *gin.Context) {
	var authcodeReq models.AuthcodeReq

	err := c.ShouldBindJSON(&authcodeReq)
	if err != nil {
		c.JSON(200, Result.Fail("参数错误"))
		fmt.Printf("数据: ", authcodeReq)
	}

	err = AuthcodeService.AddAuthcode(&authcodeReq)
	if err != nil {
		c.JSON(200, Result.Fail(err.Error()))
	}

	c.JSON(200, Result.Success("添加签到码成功"))
}

func GetAuthcode(c *gin.Context) {
	var authcodeReq models.AuthcodeReq

	err := c.ShouldBindJSON(&authcodeReq)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(200, Result.Fail("参数错误"))
	}

	authcodeResp, err := AuthcodeService.GetAuthcode(&authcodeReq)
	if err != nil {
		c.JSON(200, Result.Fail(err.Error()))
	}

	c.JSON(200, Result.Success(authcodeResp))
}
