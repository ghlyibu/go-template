package models

import (
	"crypto/md5"
	"encoding/hex"
	"go-template/dao/mysql"
	"go-template/pkg/e"
)

type User struct {
	UserID   int64  `db:"user_id"`
	Username string `db:"username"`
	Password string `db:"email"`
	Gender   string `db:"gender"`
}

// ParamSignUp 注册请求参数
type ParamSignUp struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Gender   string `json:"gender"`
}

// ParamLogin 登录请求参数
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// CheckUserExist 检查指定用户名的用户是否存在
func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int64
	if err := mysql.Db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return e.ErrorUserExist
	}
	return
}

// InsertUser 想数据库中插入一条新的用户记录
func InsertUser(user *User) (err error) {
	// 对密码进行加密
	user.Password = encryptPassword(user.Password)
	// 执行SQL语句入库
	sqlStr := `insert into user(user_id, username, password) values(?,?,?)`
	_, err = mysql.Db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}

// encryptPassword 密码加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
