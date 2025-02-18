package middleware

import (
	"bluebell/pkg/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	ContextUserIDKey string = "userID"
	ContextJwtPrefix string = "Bearer"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// tip : 这里的具体实现方式要依据我们的实际业务情况决定，并不是一成不变的！！！！
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 2003,
				"msg":  "请求头中auth为空",
			})
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == ContextJwtPrefix) {
			c.JSON(http.StatusOK, gin.H{
				"code": 2004,
				"msg":  "请求头中auth格式有误",
			})
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 2005,
				"msg":  "无效的Token",
			})
			c.Abort()
			return
		}
		// 将当前请求的userID信息保存到请求的上下文c上
		// tip : 我们不应该在代码中出现一些莫名其妙的字符串，而是应当使用常量来代替这个字符串
		c.Set(ContextUserIDKey, mc.UserID)

		// tip : 后续的处理函数可以用过c.Get(ContextUserIDKey)来获取当前请求的用户信息，并且每个用户的userID都是独一无二的
		// tip : 因此我们这样做就可以不用混淆了！！！
		c.Next()
	}
}
