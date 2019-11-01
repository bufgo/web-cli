package main

import (
	"bufgo/conf"
	"bufgo/server"
)

func main() {
	// 初始化配置
	conf.Init()

	// 加载路由
	r := server.NewRouter()
	r.Run()
}
