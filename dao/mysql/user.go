package mysql

import (
	"bluebell/models"
	"crypto/md5"
	"encoding/hex"
	"errors"
)

const salty = "JingpuRen"

// CheckUserExist 根据用户名查询指定的用户是否存在
func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var cnt int
	if err = db.Get(&cnt, sqlStr, username); err != nil {
		return
	}

	if cnt > 0 {
		return errors.New("user has been created")
	}

	return

}

// InsertUser 向user表插入新的用户
func InsertUser(user *models.User) (err error) {
	// 对密码进行加密存储
	oPassword := encryptPassword(user.Password)
	// 执行SQL语句入库
	sqlStr := `insert into user(user_id,username,password) value(?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, oPassword)
	return
}

// encryptPassword 对传入的密码进行加密，返回的是加密后的密码
func encryptPassword(oPassword string) string {
	h := md5.New()
	// 带盐的加密，就像计算机面试刷题平台那样
	h.Write([]byte(salty))
	// tip : h.Sum([]byte(oPassword))返回的是字节，通过EncodeToString()方法将其转化为十六进制的字符串
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
