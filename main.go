package main

import (
	"simple_gin_demo/conf"
	"simple_gin_demo/server"
)

func main() {
	// 从配置文件读取配置
	conf.Init()

	r := server.NewRouter()
	r.Run(":3000")
}
