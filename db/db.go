package db

import (
	mysqlDB "bufgo/db/mysql"
	postgresqlDB "bufgo/db/postgresql"
	redisDB "bufgo/db/redis"
	//sqlite3DB "bufgo/db/sqlite3"
	sqlserverDB "bufgo/db/sqlserver"
	"os"
)

// Init 初始化数据库
func Init(databaseList []string) {
	// 判断需要开启的数据库
	for i := range databaseList {
		switch databaseList[i] {
		case "MYSQL":
			mysqlDB.MYSQLInit(os.Getenv("MYSQL_DSN"))
		case "REDIS":
			redisDB.REDISInit(os.Getenv("REDIS_DB"))
		case "POSTGRESQL":
			postgresqlDB.POSTGRESQLInit(os.Getenv("POSTGRESQL_DSN"))
		//case "SQLITE3":
		//	sqlite3DB.SQLITE3Init(os.Getenv("SQLITE3_DSN"))
		case "SQLSERVER":
			sqlserverDB.SQLSERVERInit(os.Getenv(os.Getenv("SQLSERVER_DSN")))
		}
	}
}