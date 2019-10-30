package db

import (
	"fmt"
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
		// TODO
		//util.Log().Panic("连接Redis不成功", err)
		fmt.Println("Connect redis error")
		return
	}

	RedisClient = client

	fmt.Println("Connect redis success")
}