package util

import (
	"io"
	"log"
	"os"
	"path"
	"time"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

// BuildLogger 构建logger
func BuildLogger(logLevel logrus.Level) {
	if logger != nil {
		src, _ := setOutputFile()
		//设置输出
		logger.Out = src
		return
	}
	//实例化
	l := logrus.New()
	writer1_file, _ := setOutputFile() //文件
	writer2_console := os.Stdout       //终端

	//同时写入终端和文件中
	l.SetOutput(io.MultiWriter(writer1_file, writer2_console))
	//设置日志级别
	l.SetLevel(logLevel)
	//设置日志格式
	l.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	//增加行号和文件名
	l.SetReportCaller(true)

	/*
		加个hook形成ELK体系
		但是考虑到一些同学一下子接受不了那么多技术栈，
		所以这里的ELK体系加了注释，如果想引入可以直接注释去掉，
		如果不想引入这样注释掉也是没问题的。
	*/
	//hook := model.EsHookLog()
	//l.AddHook(hook)
	logger = l

}

func setOutputFile() (*os.File, error) {
	now := time.Now()
	logFilePath := ""
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/logs/"
	}
	_, err := os.Stat(logFilePath)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(logFilePath, 0777); err != nil {
			log.Println(err.Error())
			return nil, err
		}
	}
	logFileName := now.Format("2006-01-02") + ".log"
	//日志文件
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			log.Println(err.Error())
			return nil, err
		}
	}
	//写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return src, nil
}

// Log 返回日志对象
func Log() *logrus.Logger {
	if logger == nil {
		BuildLogger(logrus.DebugLevel)
	}
	return logger
}
