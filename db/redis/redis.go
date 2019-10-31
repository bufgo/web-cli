package db

import (
	"bufgo/logger"
	"github.com/go-redis/redis"
	"os"
	"strconv"
)


var (
	// RedisClient Redis缓存客户端单例
	RedisClient *redis.Client
)

// REDISInit 初始化redis链接
func REDISInit(connString string) {
	db, _ := strconv.ParseUint(connString, 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PW"),
		DB:       int(db),
	})

	_, err := client.Ping().Result()

	if err != nil {
		// 连接数据库失败
		logger.Panicf("Connect redis error")
	}

	RedisClient = client

	logger.Infof("Connect redis success")
}