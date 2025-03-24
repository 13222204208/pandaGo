package login

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mojocn/base64Captcha"
	"server/api/admin/login"
	"server/internal/dao"
	"server/internal/service"
	"server/utility"
)

type sLogin struct {
	captcha *base64Captcha.Captcha
}

func init() {
	service.RegisterLogin(New())
}

func New() *sLogin {
	// 初始化验证码驱动
	driver := base64Captcha.NewDriverDigit(80, 240, 4, 0.7, 80)
	store := base64Captcha.DefaultMemStore
	return &sLogin{
		captcha: base64Captcha.NewCaptcha(driver, store),
	}
}

// GetCaptcha 获取验证码
func (s *sLogin) GetCaptcha(ctx context.Context, req *login.CaptchaReq) (res *login.CaptchaRes, err error) {
	// 生成验证码
	driver := s.captcha.Driver.(*base64Captcha.DriverDigit)
	driver.Height = 80
	driver.Width = 240
	driver.Length = 4
	driver.MaxSkew = 0.7
	driver.DotCount = 80

	c := base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore)
	id, content, _, err := c.Generate()
	if err != nil {
		return nil, err
	}

	return &login.CaptchaRes{
		CaptchaId:  id,
		CaptchaImg: content,
	}, nil
}

// Login 用户登录
func (s *sLogin) Login(ctx context.Context, req *login.LoginReq) (res *login.LoginRes, err error) {
	// 验证码校验
	if !s.captcha.Verify(req.CaptchaId, req.VerifyCode, true) {
		return nil, gerror.New("验证码错误")
	}

	// 查询用户
	user, err := dao.PandaUser.Ctx(ctx).Where(dao.PandaUser.Columns().Username, req.Username).One()
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, gerror.New("用户名或密码错误")
	}

	// 验证密码
	if err = utility.ComparePassword(user["password"].String(), req.Password); err != nil {
		return nil, gerror.New("用户名或密码错误")
	}

	// 检查用户状态
	if user["status"].Int() != 1 {
		return nil, gerror.New("账号已被禁用")
	}

	// 更新登录信息
	_, err = dao.PandaUser.Ctx(ctx).
		Data(g.Map{
			dao.PandaUser.Columns().LastLoginTime: utility.GetCurrentTime(),
			dao.PandaUser.Columns().LastLoginIp:   g.RequestFromCtx(ctx).GetClientIp(),
		}).
		Where(dao.PandaUser.Columns().Id, user["id"]).
		Update()
	if err != nil {
		return nil, err
	}

	// 生成令牌
	userId := user["id"].Int64()
	tokenInfo, err := s.generateToken(userId)
	if err != nil {
		return nil, err
	}

	// 获取用户角色
	roles, err := s.getUserRoles(ctx, userId)
	if err != nil {
		return nil, err
	}

	// 获取用户权限
	permissions, err := s.getUserPermissions(ctx, userId)
	if err != nil {
		return nil, err
	}

	// 返回用户信息和令牌
	return &login.LoginRes{
		UserId:       userId,
		Username:     user["username"].String(),
		Nickname:     user["nickname"].String(),
		Avatar:       user["avatar"].String(),
		Roles:        roles,
		Permissions:  permissions,
		AccessToken:  tokenInfo.AccessToken,
		RefreshToken: tokenInfo.RefreshToken,
		Expires:      tokenInfo.ExpiresStr,
	}, nil
}

// TokenInfo 令牌信息
type TokenInfo struct {
	AccessToken    string
	RefreshToken   string
	AccessExpires  int64
	RefreshExpires int64
	ExpiresStr     string
}

// generateToken 生成令牌
func (s *sLogin) generateToken(userId int64) (*TokenInfo, error) {
	// 生成访问令牌
	accessToken, err := utility.GenerateToken(userId, "access")
	if err != nil {
		return nil, err
	}

	// 生成刷新令牌
	refreshToken, err := utility.GenerateToken(userId, "refresh")
	if err != nil {
		return nil, err
	}

	// 获取过期时间
	accessExpires := utility.GetTokenExpires("access")
	refreshExpires := utility.GetTokenExpires("refresh")

	// 格式化过期时间为字符串
	expiresStr := utility.FormatTime(accessExpires)

	return &TokenInfo{
		AccessToken:    accessToken,
		RefreshToken:   refreshToken,
		AccessExpires:  accessExpires,
		RefreshExpires: refreshExpires,
		ExpiresStr:     expiresStr,
	}, nil
}

// getUserRoles 获取用户角色
func (s *sLogin) getUserRoles(ctx context.Context, userId int64) ([]string, error) {
	// 查询用户角色
	var roleNames []string

	// 通过用户ID查询角色
	roleRecords, err := dao.PandaUserRole.Ctx(ctx).
		Where(dao.PandaUserRole.Columns().UserId, userId).
		All()
	if err != nil {
		return nil, err
	}

	if len(roleRecords) > 0 {
		var roleIds []int64
		for _, record := range roleRecords {
			roleIds = append(roleIds, record["role_id"].Int64())
		}

		// 查询角色信息
		roles, err := dao.PandaRole.Ctx(ctx).
			Fields(dao.PandaRole.Columns().Code).
			WhereIn(dao.PandaRole.Columns().Id, roleIds).
			All()
		if err != nil {
			return nil, err
		}

		for _, role := range roles {
			roleNames = append(roleNames, role["code"].String())
		}
	}

	// 如果没有角色，默认给一个普通用户角色
	if len(roleNames) == 0 {
		roleNames = append(roleNames, "user")
	}

	return roleNames, nil
}

// getUserPermissions 获取用户权限
func (s *sLogin) getUserPermissions(ctx context.Context, userId int64) ([]string, error) {
	// 使用 Casbin 获取用户权限

	// 如果是管理员角色，给予所有权限

	// 格式化权限
	var permissions []string

	return permissions, nil
}

// RefreshToken 刷新令牌
func (s *sLogin) RefreshToken(ctx context.Context, req *login.RefreshTokenReq) (res *login.RefreshTokenRes, err error) {
    // 刷新访问令牌
    accessToken, expires, err := utility.RefreshToken(req.RefreshToken)
    if err != nil {
        return nil, gerror.New("刷新令牌失败")
    }

    // 格式化过期时间
    expiresStr := utility.FormatTime(expires)

    return &login.RefreshTokenRes{
        AccessToken:  accessToken,
        RefreshToken: req.RefreshToken, // 返回原有的刷新令牌
        Expires:      expiresStr,
    }, nil
}
