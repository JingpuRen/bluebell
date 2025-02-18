package controller

import (
	"bluebell/logic"
	"net/http"
	"strconv"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// todo 社区相关的处理

// CommunityHandler tip : 查询所有的社区名称
func CommunityHandler(ctx *gin.Context) {
	// 1. 查询所有的社区名称
	communityNameList, err := logic.GetCommunityList()

	if err != nil {
		zap.L().Error("Controller\\community.go CommunityHandler failed", zap.Error(err))
		ctx.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "All communities selected",
		"data": communityNameList,
	})
}

// CommunityDetailHandler tip : 根据路径参数中的id查询社区详情，也就是introduction字段
func CommunityDetailHandler(ctx *gin.Context) {
	// 1. 获取路径参数中的id tip : key就是路径参数中:后面的东西
	cid := ctx.Param("id")
	community_id, err := strconv.Atoi(cid)
	if err != nil {
		zap.L().Error("Controller\\community.go CommunityDetailHandler failed", zap.Error(err))
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "the param id is invalid",
		})
		return
	}
	// 2. 根据id获取社区的详情
	detail, err := logic.GetCommunityDetail(community_id)
	if err != nil {
		zap.L().Error("Controller\\community.go CommunityDetailHandler failed", zap.Error(err))
		ctx.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}
	// 3. 返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "All community details selected",
		"data": detail,
	})
}
