// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PandaMenu is the golang structure for table panda_menu.
type PandaMenu struct {
	Id              int64       `json:"id"              orm:"id"               description:"菜单ID"`                      // 菜单ID
	ParentId        int64       `json:"parentId"        orm:"parent_id"        description:"父级ID"`                      // 父级ID
	Title           string      `json:"title"           orm:"title"            description:"菜单标题"`                      // 菜单标题
	Name            string      `json:"name"            orm:"name"             description:"路由名称"`                      // 路由名称
	Path            string      `json:"path"            orm:"path"             description:"路由地址"`                      // 路由地址
	Component       string      `json:"component"       orm:"component"        description:"组件路径"`                      // 组件路径
	Rank            int         `json:"rank"            orm:"rank"             description:"排序号"`                       // 排序号
	Redirect        string      `json:"redirect"        orm:"redirect"         description:"重定向地址"`                     // 重定向地址
	Icon            string      `json:"icon"            orm:"icon"             description:"图标"`                        // 图标
	ExtraIcon       string      `json:"extraIcon"       orm:"extra_icon"       description:"额外图标"`                      // 额外图标
	EnterTransition string      `json:"enterTransition" orm:"enter_transition" description:"进场动画"`                      // 进场动画
	LeaveTransition string      `json:"leaveTransition" orm:"leave_transition" description:"离场动画"`                      // 离场动画
	ActivePath      string      `json:"activePath"      orm:"active_path"      description:"激活路径"`                      // 激活路径
	Auths           string      `json:"auths"           orm:"auths"            description:"权限标识"`                      // 权限标识
	MenuType        int         `json:"menuType"        orm:"menu_type"        description:"菜单类型（0菜单 1iframe 2外链 3按钮）"` // 菜单类型（0菜单 1iframe 2外链 3按钮）
	FrameSrc        string      `json:"frameSrc"        orm:"frame_src"        description:"内嵌iframe地址"`                // 内嵌iframe地址
	FrameLoading    int         `json:"frameLoading"    orm:"frame_loading"    description:"iframe加载状态(0关闭 1开启)"`       // iframe加载状态(0关闭 1开启)
	KeepAlive       int         `json:"keepAlive"       orm:"keep_alive"       description:"缓存路由(0关闭 1开启)"`             // 缓存路由(0关闭 1开启)
	HiddenTag       int         `json:"hiddenTag"       orm:"hidden_tag"       description:"隐藏标签(0显示 1隐藏)"`             // 隐藏标签(0显示 1隐藏)
	FixedTag        int         `json:"fixedTag"        orm:"fixed_tag"        description:"固定标签(0不固定 1固定)"`            // 固定标签(0不固定 1固定)
	ShowLink        int         `json:"showLink"        orm:"show_link"        description:"显示链接(0隐藏 1显示)"`             // 显示链接(0隐藏 1显示)
	ShowParent      int         `json:"showParent"      orm:"show_parent"      description:"显示父级(0隐藏 1显示)"`             // 显示父级(0隐藏 1显示)
	Status          int         `json:"status"          orm:"status"           description:"状态(0停用 1正常)"`               // 状态(0停用 1正常)
	CreatedAt       *gtime.Time `json:"createdAt"       orm:"created_at"       description:"创建时间"`                      // 创建时间
	UpdatedAt       *gtime.Time `json:"updatedAt"       orm:"updated_at"       description:"更新时间"`                      // 更新时间
}
