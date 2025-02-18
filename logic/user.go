package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/jwt"
	"bluebell/pkg/snowflake"
	"errors"

	"go.uber.org/zap"
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

func SignIn(p *models.ParamSignIn) (string, error) {
	// 判断用户名是否存在
	if err := mysql.CheckUserIsExistForLogin(p.Username); err != nil {
		zap.L().Error("Logic\\user.go SignIn failed", zap.Error(err))
		return "", errors.New("用户名不存在")
	}
	// 判断该用户输入的密码是否正确
	if err := mysql.CheckPasswordByUsername(p); err != nil {
		zap.L().Error("Logic\\user.go SignIn failed", zap.Error(err))
		return "", errors.New("用户输入的密码不正确")
	}
	// 获取user_id
	userId, err := mysql.GetUserIDByUsername(p.Username)
	if err != nil {
		return "", err
	}

	token, err := jwt.GenToken(userId, p.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}
