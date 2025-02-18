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

	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "ok")
	})

	// 注册业务路由
	router.GET("./signup", controller.SignUpHandler)

	// 登录业务路由
	router.GET("./signin", controller.SignInHandler)

	router.GET("./ping", middleware.JWTAuthMiddleware(), func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})

	return router
}
