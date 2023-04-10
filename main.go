package main

import (
	"simple_gin_demo/conf"
	"simple_gin_demo/server"
	"simple_gin_demo/util"
)

func main() {
	// 从配置文件读取配置
	conf.Init()

	r := server.NewRouter()
	util.Log().Info("服务启动%v", 3000)
	r.Run(":3000")
}
