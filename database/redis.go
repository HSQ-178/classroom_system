package database

import "github.com/redis/go-redis/v9"

var redisDB *redis.Client

func InitClient() {
	redisDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func GetRedis() *redis.Client {
	return redisDB
}
