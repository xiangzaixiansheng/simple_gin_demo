package middleware

import (
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Cors 跨域配置
func Cors() gin.HandlerFunc {
	config := cors.Config{
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders: []string{
			"Origin",
			"Content-Length",
			"Content-Type",
			"Authorization",
			"Accept",
			"X-Requested-With",
			"Cookie",
		},
		ExposeHeaders: []string{
			"Content-Length",
			"Access-Control-Allow-Origin",
			"Access-Control-Allow-Headers",
		},
		MaxAge: 12 * time.Hour,
	}

	if gin.Mode() == gin.ReleaseMode {
		// 生产环境从环境变量读取允许的域名
		allowedOrigins := os.Getenv("ALLOWED_ORIGINS")
		if allowedOrigins != "" {
			config.AllowOrigins = strings.Split(allowedOrigins, ",")
		} else {
			// 默认生产环境域名
			config.AllowOrigins = []string{"https://www.example.com"}
		}
	} else {
		// 开发环境允许本地域名
		config.AllowOriginFunc = func(origin string) bool {
			// 允许本地开发域名
			if regexp.MustCompile(`^https?://127\.0\.0\.1(:\d+)?$`).MatchString(origin) {
				return true
			}
			if regexp.MustCompile(`^https?://localhost(:\d+)?$`).MatchString(origin) {
				return true
			}
			// 允许开发环境域名
			if regexp.MustCompile(`^https?://[a-zA-Z0-9-]+\.local(:\d+)?$`).MatchString(origin) {
				return true
			}
			return false
		}
	}

	// 允许携带认证信息
	config.AllowCredentials = true

	return cors.New(config)
}
