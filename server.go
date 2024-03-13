package main

import (
	"classroom-system/database"
	"classroom-system/routes"
)

func main() {
	database.InitMysql() //初始化MYSQL数据库

	router := routes.InitRouter() // 初始化路由

	_ = router.Run(":8081")
}
