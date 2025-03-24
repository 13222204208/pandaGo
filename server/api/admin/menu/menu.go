package menu

import (
	"github.com/gogf/gf/v2/frame/g"
)

// MenuBase 菜单基础字段
type MenuBase struct {
	ParentId        int64  `p:"parentId" dc:"父级ID"`
	Title           string `p:"title" v:"required" dc:"菜单名称"`
	Name            string `p:"name" dc:"路由名称"`
	Path            string `p:"path" dc:"路由路径"`
	Component       string `p:"component" dc:"组件路径"`
	Rank            int    `p:"rank" d:"0" dc:"排序"`
	Redirect        string `p:"redirect" dc:"重定向路径"`
	Icon            string `p:"icon" dc:"菜单图标"`
	ExtraIcon       string `p:"extraIcon" dc:"额外图标"`
	EnterTransition string `p:"enterTransition" dc:"进入动画"`
	LeaveTransition string `p:"leaveTransition" dc:"离开动画"`
	ActivePath      string `p:"activePath" dc:"激活路径"`
	Auths           string `p:"auths" dc:"权限标识"`
	FrameSrc        string `p:"frameSrc" dc:"iframe地址"`
	MenuType        int    `p:"menuType" v:"required|in:0,1,2,3" dc:"菜单类型(0:菜单 1:iframe 2:外链 3:按钮)"`
	FrameLoading    int    `p:"frameLoading" dc:"iframe加载状态"`
	KeepAlive       int    `p:"keepAlive" dc:"是否缓存"`
	HiddenTag       int    `p:"hiddenTag" dc:"是否隐藏标签"`
	FixedTag        int    `p:"fixedTag" dc:"是否固定标签"`
	ShowLink        int    `p:"showLink" d:"true" dc:"是否显示"`
	ShowParent      int    `p:"showParent" d:"true" dc:"是否显示父级"`
}

// CreateReq 创建菜单请求
type CreateReq struct {
	g.Meta `path:"/menu" method:"post" tags:"Menu" summary:"创建菜单"`
	MenuBase
}

// CreateRes 创建菜单响应
type CreateRes struct{}

// DeleteReq 删除菜单请求
type DeleteReq struct {
	g.Meta `path:"/menu/{id}" method:"delete" tags:"Menu" summary:"删除菜单"`
	Id     int64 `p:"id" v:"required" dc:"菜单ID"`
}

// DeleteRes 删除菜单响应
type DeleteRes struct{}

// UpdateReq 更新菜单请求
type UpdateReq struct {
	g.Meta `path:"/menu/{id}" method:"put" tags:"Menu" summary:"更新菜单"`
	Id     int64 `p:"id" v:"required" dc:"菜单ID"`
	MenuBase
}

// UpdateRes 更新菜单响应
type UpdateRes struct{}

// GetListReq 获取菜单列表请求
type GetListReq struct {
	g.Meta   `path:"/menu" method:"get" tags:"Menu" summary:"获取菜单列表"`
	Title    string `p:"title" dc:"菜单名称"`
	MenuType int    `p:"menuType" d:"-1" dc:"菜单类型(0:菜单 1:iframe 2:外链 3:按钮 -1:全部)"`
}

// GetListRes 获取菜单列表响应
type GetListRes struct {
	List interface{} `json:"list" dc:"菜单列表"`
}

// GetInfoReq 获取菜单详情请求
type GetInfoReq struct {
	g.Meta `path:"/menu/{id}" method:"get" tags:"Menu" summary:"获取菜单详情"`
	Id     int64 `p:"id" v:"required" dc:"菜单ID"`
}

// GetInfoRes 获取菜单详情响应
type GetInfoRes struct {
	Info interface{} `json:"info" dc:"菜单信息"`
}

// GetSimpleListReq 获取简化菜单列表请求
type GetSimpleListReq struct {
	g.Meta `path:"/menu/simple" method:"get" tags:"Menu" summary:"获取简化菜单列表"`
}

// SimpleMenu 简化菜单结构
type SimpleMenu struct {
	Id       int64  `json:"id"`
	ParentId int64  `json:"parentId"`
	MenuType int    `json:"menuType"`
	Title    string `json:"title"`
}

// GetSimpleListRes 获取简化菜单列表响应
type GetSimpleListRes struct {
	List []*SimpleMenu `json:"list" dc:"菜单列表"`
}
