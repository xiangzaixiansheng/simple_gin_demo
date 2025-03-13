package api

import (
	"net/http"
	"simple_gin_demo/serializer"
	"simple_gin_demo/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserRegister 用户注册接口
func UserRegister(c *gin.Context) {
	var service service.UserRegisterService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Register()
		if res.Code != 0 {
			c.JSON(http.StatusBadRequest, res)
			return
		}
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

// UserLogin 用户登录接口
func UserLogin(c *gin.Context) {
	var service service.UserLoginService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Login(c)
		if res.Code != 0 {
			c.JSON(http.StatusUnauthorized, res)
			return
		}
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}

// UserMe 用户详情
func UserMe(c *gin.Context) {
	user := CurrentUser(c)
	if user == nil {
		c.JSON(http.StatusUnauthorized, serializer.Response{
			Code: 40001,
			Msg:  "用户未登录",
		})
		return
	}
	res := serializer.BuildUserResponse(*user)
	c.JSON(http.StatusOK, res)
}

// UserLogout 用户登出
func UserLogout(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	if err := s.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, serializer.Response{
			Code: 50001,
			Msg:  "登出失败",
		})
		return
	}
	c.JSON(http.StatusOK, serializer.Response{
		Code: 0,
		Msg:  "登出成功",
	})
}
