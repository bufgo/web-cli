package db

import (
	"bufgo/logger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

// MYSQL DB 数据库链接实例
var (
	MYSQLClient *gorm.DB
)

// MYSQLInit 初始化MYSQL链接
func MYSQLInit(connString string) {
	db, err := gorm.Open("mysql", connString)
	db.LogMode(true)
	if err != nil {
		// 连接数据库失败
		logger.Panicf("Connect mysql error")
	}
	//设置连接池
	//空闲
	db.DB().SetMaxIdleConns(50)
	//打开
	db.DB().SetMaxOpenConns(100)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)

	MYSQLClient = db

	logger.Infof("Connect mysql success")
}