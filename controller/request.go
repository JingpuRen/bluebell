package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// todo : 用来获取登录用户的id

// GetCurrentUser tip : 从上下文中获取用户的登录信息，如果获取不到，说明用户根本就没有登录
func GetCurrentUser(ctx *gin.Context) (userId int64, err error) {
	// tip : 下次可以看看每个方法的返回值的类型到底是什么！！！

	uid, isOk := ctx.Get("userID")
	if isOk != true {
		err = errors.New("用户未登录")
		return
	}
	userId = uid.(int64)

	// 返回用户的登录信息
	return
}
