package middleware

import (
	"fmt"
	"math"
	"os"
	"simple_gin_demo/util"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// LoggerMiddleware 日志中间件
func LoggerMiddleware() gin.HandlerFunc {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()
		latency := endTime.Sub(startTime)
		spendTime := fmt.Sprintf("%d ms", int(math.Ceil(float64(latency.Nanoseconds()/1000000))))

		statusCode := c.Writer.Status()
		dataSize := c.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}

		// 构建日志字段
		fields := logrus.Fields{
			"hostname":  hostname,
			"status":    statusCode,
			"latency":   spendTime,
			"client_ip": c.ClientIP(),
			"method":    c.Request.Method,
			"path":      path,
			"size":      dataSize,
			"user_agent": c.Request.UserAgent(),
		}

		if query != "" {
			fields["query"] = query
		}

		// 获取当前用户（如果已登录）
		if user, exists := c.Get("user"); exists && user != nil {
			fields["user_id"] = user
		}

		// 添加请求头信息
		if referer := c.Request.Referer(); referer != "" {
			fields["referer"] = referer
		}

		log := util.LogrusObj.WithFields(fields)

		// 根据状态码记录日志
		if len(c.Errors) > 0 {
			log.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		} else {
			switch {
			case statusCode >= 500:
				log.Error("Server error")
			case statusCode >= 400:
				log.Warn("Client error")
			case statusCode >= 300:
				log.Info("Redirection")
			default:
				log.Info("Success")
			}
		}
	}
}
