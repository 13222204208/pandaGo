// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PandaUserRole is the golang structure for table panda_user_role.
type PandaUserRole struct {
	Id        int64       `json:"id"        orm:"id"         description:"用户角色关联ID"` // 用户角色关联ID
	UserId    int64       `json:"userId"    orm:"user_id"    description:"用户ID"`     // 用户ID
	RoleId    int64       `json:"roleId"    orm:"role_id"    description:"角色ID"`     // 角色ID
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`     // 创建时间
}
