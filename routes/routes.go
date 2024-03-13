package routes

import (
	"classroom-system/app/controllers"
	"classroom-system/app/middlewares/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// 初始化一个http服务对象
	r := gin.Default()

	//静态文件服务
	r.Static("/static", "./static")

	//启用跨域中间件
	r.Use(cors.Cors())

	root := r.Group("/api")
	{
		//老师
		teacher := root.Group("/teacher")
		{
			teacher.POST("/register", controllers.Register) //老师注册
		}
	}

	return r
}
