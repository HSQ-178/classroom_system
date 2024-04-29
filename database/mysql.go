package database

import (
	"fmt"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB 数据库连接池
var db *gorm.DB

func InitMysql() {
	var err error

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,
		},
	)

	dsn := "root:123456@tcp(localhost:3306)/classroom_system?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
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
