package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"fmt"
	"net/http"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// SignUpHandler 处理注册请求的函数
func SignUpHandler(ctx *gin.Context) {
	// tip : 1.获取参数和参数校验
	var p models.ParamSignUp
	err := ctx.ShouldBindJSON(&p)
	if err != nil {
		zap.L().Error("Controller\\user.go SignUpHandler failed", zap.Error(err))
		// 请求参数有误
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "请求参数有误",
		})
		return
	}
	// 对请求参数进行更为详细的业务规则校验
	if len(p.Username) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0 || p.RePassword != p.Password {
		zap.L().Error("Controller\\user.go SignUpHandler failed")
		// 请求参数有误
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "参数为空或者两次密码的输入不一致",
		})
		return
	}
	fmt.Println(p)
	// tip : 2.业务处理
	err = logic.SignUp(&p)

	// tip : 3.返回结果
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "注册失败",
		})
		return // 防止返回下面的结果，因此在这里直接返回
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}

// SignInHandler 处理登录请求的函数
func SignInHandler(ctx *gin.Context) {
	// tip : 1. 获取请求参数及参数校验
	var p models.ParamSignIn
	if err := ctx.ShouldBindJSON(&p); err != nil {
		zap.L().Error("Controller\\user.go SignInHandler failed", zap.Error(err))
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "参数不正确",
		})
		return
	}
	if len(p.Username) == 0 || len(p.Password) == 0 {
		zap.L().Error("Controller\\user.go SignInHandler failed")
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "参数中有空值",
		})
		return
	}
	// tip : 2. 业务处理
	token, err := logic.SignIn(&p)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}
	// tip : 3. 返回结果
	ctx.JSON(http.StatusOK, gin.H{
		// 像这种有多个返回值的时候，一般会按照首字母的顺序进行排列
		"msg":  "用户登录成功",
		"data": token,
	})
}
