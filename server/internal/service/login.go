// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"server/api/admin/login"
)

type (
	ILogin interface {
		// GetCaptcha 获取验证码
		GetCaptcha(ctx context.Context, req *login.CaptchaReq) (res *login.CaptchaRes, err error)
		// Login 用户登录
		Login(ctx context.Context, req *login.LoginReq) (res *login.LoginRes, err error)
		// RefreshToken 刷新令牌
		RefreshToken(ctx context.Context, req *login.RefreshTokenReq) (res *login.RefreshTokenRes, err error)
	}
)

var (
	localLogin ILogin
)

func Login() ILogin {
	if localLogin == nil {
		panic("implement not found for interface ILogin, forgot register?")
	}
	return localLogin
}

func RegisterLogin(i ILogin) {
	localLogin = i
}
