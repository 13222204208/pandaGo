package admin

import (
	"context"
	"server/internal/service"

	"server/api/admin/login"
)

func (c *ControllerLogin) Login(ctx context.Context, req *login.LoginReq) (res *login.LoginRes, err error) {
	res, err = service.Login().Login(ctx, req)
	return
}
func (c *ControllerLogin) Captcha(ctx context.Context, req *login.CaptchaReq) (res *login.CaptchaRes, err error) {
	res, err = service.Login().GetCaptcha(ctx, req)
	return
}
func (c *ControllerLogin) RefreshToken(ctx context.Context, req *login.RefreshTokenReq) (res *login.RefreshTokenRes, err error) {
	res, err = service.Login().RefreshToken(ctx, req)
	return
}
