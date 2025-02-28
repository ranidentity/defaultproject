package middleware

import (
	"defaultproject/model"
	"defaultproject/repository"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// CurrentUser 获取登录用户
func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		_uid := session.Get("user_id")
		if _uid != nil {
			uid := _uid.(uint)
			user, err := repository.GetUserFromId(uid)
			if err == nil {
				c.Set("user", &user)
			}
		}
		c.Next()
	}
}

// AuthRequired 需要登录
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, _ := c.Get("user"); user != nil {
			if _, ok := user.(*model.User); ok {
				c.Next()
				return
			}
		}
		// c.JSON(200, serializer.CheckLogin()) when login fail
		c.Abort()
	}
}
