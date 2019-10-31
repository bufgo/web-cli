package db

import (
	"bufgo/logger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"time"
)

// SQLITE3 DB 数据库链接实例
var (
	SQLITE3Client *gorm.DB
)

// SQLITE3Init 初始化SQLITE3链接
func SQLITE3Init(connString string) {
	db, err := gorm.Open("sqlite3", connString)
	db.LogMode(true)
	if err != nil {
		// 连接数据库失败
		logger.Panicf("Connect sqlite3 error")
	}
	//设置连接池
	//空闲
	db.DB().SetMaxIdleConns(50)
	//打开
	db.DB().SetMaxOpenConns(100)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)

	SQLITE3Client = db

	logger.Infof("Connect sqlite3 success")
}