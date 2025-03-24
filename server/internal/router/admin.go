package router

import (
	"server/internal/controller/admin"
	"server/internal/middleware"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func InitAdminRouter() {
	s := g.Server()
	//不验证token
	s.Group("/api/admin", func(group *ghttp.RouterGroup) {
		group.Middleware(ghttp.MiddlewareHandlerResponse)
		group.Bind(
			admin.NewLogin(), // 后台登录
		)
	})

	s.Group("/api/admin", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Auth)
		group.Middleware(ghttp.MiddlewareHandlerResponse)
		group.Bind(
			admin.NewDicttype(),   // 字典类型
			admin.NewDictdata(),   // 字典数据
			admin.NewUser(),       // 后台用户数据
			admin.NewRole(),       // 后台角色
			admin.NewAttachment(), // 附件管理
			admin.NewCrontab(),    // 定时任务
			admin.NewMenu(),       // 菜单管理
		)
	})
}
