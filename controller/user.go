package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-template/logic"
	"go-template/models"
	"go-template/pkg/e"
	"go.uber.org/zap"
)

type _ResponseSign struct {
	Code e.ResCode
	Msg  string
	Data any
}

// SignUpHandler 注册
// @Summary 用户注册接口
// @Description 注册
// @Tags 注册
// @Accept json
// @Produce json
// @Param data body models.ParamSignUp true "用户名，密码，确认密码"
// @Success 200 {object} _ResponseSign
// @Router /api/v1/signup [post]
func SignUpHandler(c *gin.Context) {
	// 1. 获取参数和参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, e.CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, e.CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	// 2. 业务处理
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("logic.SignUp failed", zap.Error(err))
		if errors.Is(err, e.ErrorUserExist) {
			ResponseError(c, e.CodeUserExist)
			return
		}
		ResponseError(c, e.CodeServerBusy)
		return
	}
	// 3. 返回响应
	ResponseSuccess(c, nil)
}
