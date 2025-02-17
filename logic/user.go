package logic

import "bluebell/dao/mysql"

func SignUp() {
	// 生成UID

	// 保存UID进数据库
	mysql.SignUp()
}
