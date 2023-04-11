# simple_gin_demo



一个简单的gin例子，也是方便以后其他项目直接拿来用吧



##### 一、使用到的一些库

"github.com/joho/godotenv" 加载本地.env文件的
"github.com/sirupsen/logrus" log处理的

gorm在线生成model文档
https://www.printlove.cn/tools/sql2gorm



##### 二、遇到的问题

问题一：

\"created_at\": unsupported Scan, storing driver.Value type []uint8 into type *time.Time"

在链接数据库的时候
&parseTime=true

