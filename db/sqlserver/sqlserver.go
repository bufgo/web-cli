package db

import (
	"bufgo/logger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"time"
)

// SQLSERVER DB 数据库链接实例
var (
	SQLSERVERClient *gorm.DB
)

// SQLSERVERInit 初始化SQLSERVER链接
func SQLSERVERInit(connString string) {
	db, err := gorm.Open("mssql", connString)
	db.LogMode(true)
	if err != nil {
		// 连接数据库失败
		logger.Panicf("Connect mssql error")
	}
	//设置连接池
	//空闲
	db.DB().SetMaxIdleConns(50)
	//打开
	db.DB().SetMaxOpenConns(100)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)

	SQLSERVERClient = db

	logger.Infof("Connect mssql success")
}
