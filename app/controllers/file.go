package controllers

import (
	Result "classroom-system/app/common"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"strings"
)

var path = "C:\\Users\\黄叶\\Desktop\\test\\"

func UploadFile(c *gin.Context) {

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(200, Result.Fail("文件上传失败"))
		return
	}

	id := uuid.NewString()
	filenames := strings.Split(file.Filename, ".")

	c.SaveUploadedFile(file, path+id+"."+filenames[len(filenames)-1])

	c.JSON(200, Result.Success(id+"."+filenames[len(filenames)-1]))
}

func DownloadFile(c *gin.Context) {
	filename := c.Query("filename")

	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Header("Content-Transfer-Encoding", "binary")
	c.File(path + filename)
}
