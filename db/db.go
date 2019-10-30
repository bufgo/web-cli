package db

import (
	mysqlDB "bufgo/db/mysql"
	redisDB "bufgo/db/redis"
	"os"
)

// Init 初始化数据库
func Init(databaseList []string) {
	// 判断需要开启的数据库
	for i := range databaseList {
		//fmt.Printf("%s\n", databaseList[i])
		switch databaseList[i] {
		case "MYSQL":
			mysqlDB.MYSQLInit(os.Getenv("MYSQL_DSN"))
		case "REDIS":
			redisDB.REDISInit(os.Getenv("REDIS_DB"))
		}
	}
}