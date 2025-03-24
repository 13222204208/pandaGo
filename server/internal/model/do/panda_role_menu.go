// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PandaRoleMenu is the golang structure of table panda_role_menu for DAO operations like Where/Data.
type PandaRoleMenu struct {
	g.Meta    `orm:"table:panda_role_menu, do:true"`
	Id        interface{} // 角色菜单关联ID
	RoleId    interface{} // 角色ID
	MenuId    interface{} // 菜单ID
	CreatedAt *gtime.Time // 创建时间
}
