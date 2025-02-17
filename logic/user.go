package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) error {
	// 判断用户是否已经注册
	err := mysql.CheckUserExist(p.Username)
	// tip : 像这种直接返回来的这种err我们就直接返回就可以
	if err != nil {
		return err
	}
	// 生成UID
	userID := snowflake.GenID()

	// 创建实例
	user := models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}

	// 保存注册后的数据进入数据库
	err = mysql.InsertUser(&user)
	return err
}
