package db

import (
	"bufgo/logger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

// POSTGRES DB 数据库链接实例
var (
	POSTGRESClient *gorm.DB
)

// POSTGRESQLInit 初始化POSTGRES链接
func POSTGRESQLInit(connString string) {
	db, err := gorm.Open("postgres", connString)
	db.LogMode(true)
	if err != nil {
		// 连接数据库失败
		logger.Panicf("Connect postgres error")
	}
	//设置连接池
	//空闲
	db.DB().SetMaxIdleConns(50)
	//打开
	db.DB().SetMaxOpenConns(100)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)

	POSTGRESClient = db

	logger.Infof("Connect postgres success")
}