package middleware

import (
	"net/http"
	"simple_gin_demo/model"
	"simple_gin_demo/serializer"
	"simple_gin_demo/util"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// CurrentUser 获取登录用户
func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		uid := session.Get("user_id")
		if uid != nil {
			user, err := model.GetUser(uid)
			if err == nil {
				c.Set("user", &user)
			} else {
				util.LogrusObj.Warnf("Failed to get user: %v", err)
				// Clear invalid session
				session.Clear()
				session.Save()
			}
		}
		c.Next()
	}
}

// AuthRequired 需要登录
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, exists := c.Get("user"); exists && user != nil {
			if _, ok := user.(*model.User); ok {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusUnauthorized, serializer.Response{
			Code: 40001,
			Msg:  "需要登录",
		})
		c.Abort()
	}
}

// AdminRequired 需要管理员权限
func AdminRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, exists := c.Get("user"); exists && user != nil {
			if u, ok := user.(*model.User); ok && u.Role == "admin" {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, serializer.Response{
			Code: 40003,
			Msg:  "需要管理员权限",
		})
		c.Abort()
	}
}
