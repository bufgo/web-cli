package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"time"
)

// SQLITE3 DB 数据库链接实例
var (
	SQLITE3DB *gorm.DB
)

// SQLITE3Init 初始化SQLITE3链接
func SQLITE3Init(connString string) {
	db, err := gorm.Open("sqlite3", connString)
	db.LogMode(true)
	if err != nil {
		// TODO
		// 连接数据库失败
		//util.Log().Panic("连接数据库不成功", err)
		fmt.Println("Connect sqlite3 error")
		return
	}
	//设置连接池
	//空闲
	db.DB().SetMaxIdleConns(50)
	//打开
	db.DB().SetMaxOpenConns(100)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)

	SQLITE3DB = db

	fmt.Println("Connect sqlite3 success")
}