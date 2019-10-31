package conf

import (
	"bufgo/db"
	"fmt"
	"github.com/joho/godotenv"
)

var (
	// 环境变量列表
	envList map[string]string
	// env 错误
	envErr error
	// 数据库列表
	databaseList []string
)

// Init 初始化
func Init() {
	// 从本地读取环境变量
	godotenv.Load()

	// 获取需要启动的数据库
	envList, envErr = godotenv.Read()
	if envErr != nil {
		fmt.Printf("Error: %s", envErr)
	}
	for k,v := range envList {
		if v == "true" {
			databaseList = append(databaseList, k)
		}
	}

	// 数据库
	db.Init(databaseList)
}