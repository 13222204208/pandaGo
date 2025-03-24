// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PandaRoleMenu is the golang structure for table panda_role_menu.
type PandaRoleMenu struct {
	Id        int64       `json:"id"        orm:"id"         description:"角色菜单关联ID"` // 角色菜单关联ID
	RoleId    int64       `json:"roleId"    orm:"role_id"    description:"角色ID"`     // 角色ID
	MenuId    int64       `json:"menuId"    orm:"menu_id"    description:"菜单ID"`     // 菜单ID
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`     // 创建时间
}
