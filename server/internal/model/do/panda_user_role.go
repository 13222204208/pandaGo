// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PandaUserRole is the golang structure of table panda_user_role for DAO operations like Where/Data.
type PandaUserRole struct {
	g.Meta    `orm:"table:panda_user_role, do:true"`
	Id        interface{} // 用户角色关联ID
	UserId    interface{} // 用户ID
	RoleId    interface{} // 角色ID
	CreatedAt *gtime.Time // 创建时间
}
