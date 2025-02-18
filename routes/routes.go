package routes

import (
	"bluebell/controller"
	"bluebell/logger"
	"bluebell/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetUp() *gin.Engine {
	router := gin.New()
	router.Use(logger.GinLogger(), logger.GinRecovery(true))

	v1 := router.Group("/api/v1")

	// 注册业务路由
	v1.POST("./signup", controller.SignUpHandler)

	// 登录业务路由
	v1.POST("./login", controller.SignInHandler)

	/**
	tip : middleware.JWTAuthMiddleware() 中间件会应
	tip : 用在 v1.Use(middleware.JWTAuthMiddleware())
	tip : 语句之后注册的所有 v1 路由分组下的路由上，
	tip : 而在该语句之前注册的路由不会应用这个中间件。
	*/
	v1.Use(middleware.JWTAuthMiddleware())

	{
		// tip : 应用中间件函数后，中间件函数和处理函数的执行顺序是：
		// tip : 中间件前置逻辑 -> 处理函数 -> 中间件后置逻辑

		// 查询所有社区的名称
		v1.POST("/community", controller.CommunityHandler)
		v1.POST("./community/:id", controller.CommunityDetailHandler)
	}

	// 一个路由都没有匹配上的话！！！！
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"code": "404",
			"msg":  "Not Found",
		})
	})

	return router
}
