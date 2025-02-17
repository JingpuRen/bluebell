package routes

import (
	"bluebell/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetUp() *gin.Engine {
	router := gin.New()
	router.Use(logger.GinLogger(), logger.GinRecovery(true))

	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "ok")
	})
	return router
}
