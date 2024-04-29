package routes

import (
	"classroom-system/app/controllers"
	"classroom-system/app/middlewares/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	//将日志写入文件和控制台
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
			teacher.POST("/login", controllers.Login)       //老师登录
		}

		//学生
		student := root.Group("/student")
		{
			student.POST("/getAllStudents", controllers.GetAllStudents) //获取该老师所有学生
			student.POST("/login", controllers.LoginStudent)            //学生端登录
		}

		//课程
		course := root.Group("/course")
		{
			course.POST("/getAllCourse", controllers.GerAllCourse) //获取该老师的所有课程
		}

		//记录
		record := root.Group("/record")
		{
			record.POST("/getNotSignInList", controllers.GetNotSignInList)               //获取未签到学生列表
			record.POST("/getSignInAndAbsenceList", controllers.GetSignInAndAbsenceList) //获取签到与缺勤
			record.POST("/addRecords", controllers.AddRecords)                           //添加签到或缺勤记录
			record.POST("/deleteRecords", controllers.DeleteRecords)                     //删除签到或缺勤记录
			record.POST("/updateRecords", controllers.UpdateRecords)                     //更新签到或缺勤记录
		}

		//redis
		redis := root.Group("/redis")
		{
			redis.POST("/setRedis", controllers.SetNoticeRedis) //存入redis
			redis.POST("/getRedis", controllers.GetNoticeRedis) //获取redis
		}

		qrcode := root.Group("/qrcode")
		{
			qrcode.POST("/setQrcode", controllers.SetQrcode) //将二维码存入redis
			qrcode.GET("/getQrcode", controllers.GetQrcode)  //获取二维码
			qrcode.POST("/rabbitmqQrcode", controllers.RabbitmqQrcode)
		}
	}

	return r
}
