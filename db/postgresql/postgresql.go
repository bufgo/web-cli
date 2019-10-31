package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

// POSTGRESQL DB 数据库链接实例
var (
	POSTGRESQLDB *gorm.DB
)

// POSTGRESQLInit 初始化POSTGRESQL链接
func POSTGRESQLInit(connString string) {
	// "host=myhost port=myport user=gorm dbname=gorm password=mypassword"
	db, err := gorm.Open("postgres", connString)
	db.LogMode(true)
	if err != nil {
		// TODO
		// 连接数据库失败
		//util.Log().Panic("连接数据库不成功", err)
		fmt.Println("Connect postgresql error")
		return
	}
	//设置连接池
	//空闲
	db.DB().SetMaxIdleConns(50)
	//打开
	db.DB().SetMaxOpenConns(100)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)

	POSTGRESQLDB = db

	fmt.Println("Connect postgresql success")
}