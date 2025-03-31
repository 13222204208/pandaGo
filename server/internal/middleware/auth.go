package middleware

import (
	"fmt"
	"server/internal/service"
	"server/utility"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
)

// Auth 认证中间件
func Auth(r *ghttp.Request) {
	// 获取 token
	token := r.Header.Get("Authorization")
	if token == "" {
		r.Response.WriteJson(gerror.New("未授权访问"))
		r.Exit()
		return
	}

	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	}

	// 添加调试信息
	//fmt.Printf("收到的token: %s\n", token)
	utility.InitJwtConfig()
	// 解析 token 获取用户 ID
	userId, err := utility.GetUserIdFromToken(token)
	if err != nil {
		fmt.Printf("token解析错误: %v\n", err)
		// 解析 token 详情
		claims, parseErr := utility.ParseToken(token)
		if parseErr == nil {
			fmt.Printf("token详情: %+v\n", claims)
			fmt.Printf("过期时间: %v\n", claims.ExpiresAt)
			fmt.Printf("当前时间: %v\n", time.Now())
		}
		r.Response.WriteJson(gerror.New("无效的令牌"))
		r.Exit()
		return
	}

	// 将用户 ID 存入上下文，使用统一的变量名
	r.SetCtxVar("userId", userId)
	r.Middleware.Next()
}

// Permission 权限验证中间件
func Permission(permission string) ghttp.HandlerFunc {
	return func(r *ghttp.Request) {
		// 获取用户 ID
		userId := r.GetCtxVar("userId").Int64()
		if userId == 0 {
			r.Response.WriteJson(gerror.New("未授权访问"))
			r.Exit()
			return
		}

		// 检查权限
		hasPermission, err := service.Permission().CheckPermission(r.GetCtx(), userId, permission)
		if err != nil {
			r.Response.WriteJson(gerror.New("权限检查失败"))
			r.Exit()
			return
		}

		if !hasPermission {
			r.Response.WriteJson(gerror.New("没有操作权限"))
			r.Exit()
			return
		}

		r.Middleware.Next()
	}
}
