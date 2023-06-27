package conf

import (
	"os"
	"simple_gin_demo/cache"
	"simple_gin_demo/model"
	"simple_gin_demo/util"

	"github.com/joho/godotenv"
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	godotenv.Load()

	// 读取翻译文件
	if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
		util.LogrusObj.Panic("翻译文件加载失败", err)
	}

	// 连接数据库
	model.Database(os.Getenv("MYSQL_DSN"))
	cache.Redis()
}
