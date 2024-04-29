package controllers

import (
	Result "classroom-system/app/common"
	Service "classroom-system/app/services/qrcode"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 将二维码存入redis
func SetQrcode(c *gin.Context) {
	var qrcodeInfo Service.QrcodeInfo

	err := c.ShouldBindJSON(&qrcodeInfo)
	if err != nil {
		c.JSON(200, Result.Fail("参数错误"))
		return
	}
	qrcodeInfoRedis, err := Service.SetRedisQrcode(&qrcodeInfo)
	if err != nil {
		c.JSON(200, Result.Fail(err.Error()))
		return
	}
	c.JSON(200, Result.Success(qrcodeInfoRedis))
}

// 生成二维码图像
func GetQrcode(c *gin.Context) {
	key := c.Query("teacherCard")

	value, err := Service.GenerateQrcode(key)
	if err != nil {
		c.JSON(200, Result.Fail(err.Error()))
		return
	}

	// 返回二维码图像
	c.Data(http.StatusOK, "image/png", value)
}

func RabbitmqQrcode(c *gin.Context) {
	var qrcodeInfo Service.QrcodeInfo

	err := c.ShouldBindJSON(&qrcodeInfo)
	if err != nil {
		c.JSON(200, Result.Fail("参数错误"))
		return
	}
	qrcodeInfoRedis, err := Service.SetRedisQrcode(&qrcodeInfo)
	if err != nil {
		c.JSON(200, Result.Fail(err.Error()))
		return
	}
	c.JSON(200, Result.Success(qrcodeInfoRedis))

}
