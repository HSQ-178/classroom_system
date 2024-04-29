package main

import (
	"classroom-system/app/middlewares/rabbitmq"
	"classroom-system/database"
	"classroom-system/routes"
)

func main() {
	database.InitMysql() //初始化MYSQL数据库

	database.InitClient() // 初始化Redis数据库

	rabbitmq := rabbitmq.RabbitMQConn("", "gdxz003", "")
	go rabbitmq.Consume()

	router := routes.InitRouter() // 初始化路由

	_ = router.Run(":8082")

}
