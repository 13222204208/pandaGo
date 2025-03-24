package user

import (
	"server/api/common"
	"server/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

// User 用户管理基础信息
type User struct {
	Username string `p:"username" v:"required" dc:"用户名"`
	Nickname string `p:"nickname" v:"required" dc:"昵称"`
	Email    string `p:"email" v:"email" dc:"邮箱"`
	Phone    string `p:"phone" v:"phone" dc:"手机号"`
	Status   int    `p:"status" v:"required|in:0,1" dc:"状态（0停用 1正常）"`
}

// CreateReq 创建用户
type CreateReq struct {
	g.Meta   `path:"/user" method:"post" tags:"User" summary:"创建用户"`
	Password string `p:"password" v:"required" dc:"密码"`
	User
}
type CreateRes struct {
	g.Meta `mime:"application/json"`
}

// DeleteReq 删除用户
type DeleteReq struct {
	g.Meta `path:"/user/{id}" method:"delete" tags:"User" summary:"删除用户"`
	Id     int64 `p:"id" v:"required" dc:"用户ID"`
}
type DeleteRes struct{}

// UpdateReq 更新用户
type UpdateReq struct {
	g.Meta `path:"/user/{id}" method:"put" tags:"User" summary:"更新用户"`
	Id     int64 `p:"id" v:"required" dc:"用户ID"`
	User
}
type UpdateRes struct{}

// ResetReq 重置用户的状态,头像,密码,角色分配
type ResetReq struct {
	g.Meta   `path:"/user/{id}/reset" method:"put" tags:"User" summary:"重置用户状态,头像,密码,角色分配"`
	Id       int64   `p:"id" v:"required" dc:"用户ID"`
	Status   *int    `p:"status" v:"in:0,1" dc:"状态（0停用 1正常）"`
	Avatar   string  `p:"avatar" dc:"头像"`
	Password string  `p:"password" dc:"密码"`
	RoleId   []int64 `p:"roleId" dc:"角色ID"`
}
type ResetRes struct{}

// GetListReq 获取用户列表
type GetListReq struct {
	g.Meta   `path:"/user" method:"get" tags:"User" summary:"获取用户列表"`
	Username string `p:"username" dc:"用户名"`
	Phone    string `p:"phone" dc:"手机号"`
	Status   int    `p:"status" v:"in:0,1" dc:"状态（0停用 1正常）"`
	common.PageReq
}
type GetListRes struct {
	List  []*entity.PandaUser `json:"list" dc:"用户列表"`
	Total int                 `json:"total" dc:"数据总数"`
}

// GetUserRoleIdsReq 获取用户的角色id
type GetUserRoleIdsReq struct {
	g.Meta `path:"/user/{id}/role" method:"get" tags:"User" summary:"获取用户的角色id"`
	Id     int64 `p:"id" v:"required" dc:"用户ID"`
}
type GetUserRoleIdsRes struct {
	RoleIds []int64 `json:"roleIds" dc:"角色ID列表"`
}

type MenuTreeItem struct {
	Path      string          `json:"path"`               // 路由路径
	Component string          `json:"component"`          // 组件路径
	Name      string          `json:"name"`               // 菜单名称
	Meta      MenuMeta        `json:"meta"`               // 元数据
	Children  []*MenuTreeItem `json:"children,omitempty"` // 子菜单
}

// MenuMeta 菜单元数据
type MenuMeta struct {
	Icon     string   `json:"icon"`               // 图标
	Title    string   `json:"title"`              // 标题
	Rank     int      `json:"rank,omitempty"`     // 排序
	Roles    []string `json:"roles,omitempty"`    // 角色
	HideMenu bool     `json:"hideMenu,omitempty"` // 是否隐藏菜单
}

// GetUserMenuTreeReq GetMenuTreeReq 获取用户菜单树
type GetUserMenuTreeReq struct {
	g.Meta `path:"/user/menu" method:"get" tags:"User" summary:"获取用户菜单树"`
}
type GetUserMenuTreeRes struct {
	g.Meta `mime:"application/json"`
	Menu   []*MenuTreeItem `json:"menu" dc:"菜单树"`
}
