// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PandaMenu is the golang structure of table panda_menu for DAO operations like Where/Data.
type PandaMenu struct {
	g.Meta          `orm:"table:panda_menu, do:true"`
	Id              interface{} // 菜单ID
	ParentId        interface{} // 父级ID
	Title           interface{} // 菜单标题
	Name            interface{} // 路由名称
	Path            interface{} // 路由地址
	Component       interface{} // 组件路径
	Rank            interface{} // 排序号
	Redirect        interface{} // 重定向地址
	Icon            interface{} // 图标
	ExtraIcon       interface{} // 额外图标
	EnterTransition interface{} // 进场动画
	LeaveTransition interface{} // 离场动画
	ActivePath      interface{} // 激活路径
	Auths           interface{} // 权限标识
	MenuType        interface{} // 菜单类型（0菜单 1iframe 2外链 3按钮）
	FrameSrc        interface{} // 内嵌iframe地址
	FrameLoading    interface{} // iframe加载状态(0关闭 1开启)
	KeepAlive       interface{} // 缓存路由(0关闭 1开启)
	HiddenTag       interface{} // 隐藏标签(0显示 1隐藏)
	FixedTag        interface{} // 固定标签(0不固定 1固定)
	ShowLink        interface{} // 显示链接(0隐藏 1显示)
	ShowParent      interface{} // 显示父级(0隐藏 1显示)
	Status          interface{} // 状态(0停用 1正常)
	CreatedAt       *gtime.Time // 创建时间
	UpdatedAt       *gtime.Time // 更新时间
}
