package logic

import (
	"go-template/models"
	"go-template/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error) {
	// 1.判断用户存不存在
	if err := models.CheckUserExist(p.Username); err != nil {
		return err
	}
	// 2.生成UID
	userID := snowflake.GenID()
	// 构造一个User实例
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	// 3.保存进数据库
	return models.InsertUser(user)
}
