package controller

import (
	"bluebell/logic"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SignUpHandler 处理注册请求的函数
func SignUpHandler(context *gin.Context) {
	// tip : 1.获取参数和参数校验
	// 2.业务处理
	logic.SignUp()
	// 3.返回结果
	context.JSON(http.StatusOK, gin.H{
		"name":     "Ren Jingo",
		"password": "daohaozhe250",
	})
}
