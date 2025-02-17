package logic

import (
	"bluebell/dao/mysql"
	"bluebell/pkg/snowflake"
)

func SignUp() {
	// 判断用户是否已经注册
	mysql.QueryUserByUsername()
	// 生成UID
	snowflake.GenID()
	// 密码加密操作

	// 保存UID进数据库
	mysql.InsertUser()
}
