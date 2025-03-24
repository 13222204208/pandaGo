package login

import "github.com/gogf/gf/v2/frame/g"

type LoginReq struct {
	g.Meta     `path:"/login" method:"post" tags:"登陆" summary:"后台用户登陆"`
	Username   string `p:"username" v:"required" dc:"用户名"`
	Password   string `p:"password" v:"required" dc:"密码"`
	CaptchaId  string `p:"captchaId" v:"required" dc:"验证码ID"`
	VerifyCode string `p:"verifyCode" v:"required" dc:"验证码"`
}

type LoginRes struct {
	UserId       int64    `json:"userId" dc:"用户ID"`
	Username     string   `json:"username" dc:"用户名"`
	Nickname     string   `json:"nickname" dc:"昵称"`
	Avatar       string   `json:"avatar" dc:"头像"`
	Roles        []string `json:"roles" dc:"角色列表"`
	Permissions  []string `json:"permissions" dc:"权限列表"`
	AccessToken  string   `json:"accessToken" dc:"访问令牌"`
	RefreshToken string   `json:"refreshToken" dc:"刷新令牌"`
	Expires      string   `json:"expires" dc:"过期时间"`
}

type CaptchaReq struct {
	g.Meta `path:"/captcha" method:"get" tags:"captcha" summary:"获取验证码"`
}

type CaptchaRes struct {
	CaptchaId  string `json:"captchaId"`
	CaptchaImg string `json:"captchaImg"`
}

// RefreshTokenReq 刷新令牌请求
type RefreshTokenReq struct {
    g.Meta       `path:"/refresh-token" method:"post" tags:"Login" summary:"刷新访问令牌"`
    RefreshToken string `p:"refreshToken" v:"required" dc:"刷新令牌"`
}

// RefreshTokenRes 刷新令牌响应
type RefreshTokenRes struct {
    g.Meta       `mime:"application/json"`
    AccessToken  string `json:"accessToken"  dc:"访问令牌"`
    RefreshToken string `json:"refreshToken" dc:"刷新令牌"`
    Expires      string `json:"expires"      dc:"过期时间"`
}
