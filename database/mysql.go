package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB 数据库连接池
var db *gorm.DB

func InitMysql() {
	var err error

	dsn := "root:08181029hsq@tcp(localhost:3306)/classroom_system?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("数据库连接失败,err: %v\n", err)
	}
	if db.Error != nil {
		fmt.Printf("数据库连接失败,err: %v\n", db.Error)
	}

	sqlDB, _ := db.DB()
	//设置最大连接数
	sqlDB.SetMaxOpenConns(100)
	//设置上数据库最大闲置连接数
	sqlDB.SetMaxIdleConns(20)

	fmt.Println("连接数据库成功！")
}

func GetMysql() *gorm.DB {
	return db
}
