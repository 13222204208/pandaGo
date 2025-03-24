// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PandaRole is the golang structure of table panda_role for DAO operations like Where/Data.
type PandaRole struct {
	g.Meta    `orm:"table:panda_role, do:true"`
	Id        interface{} // 角色ID
	Name      interface{} // 角色名
	Code      interface{} // 角色标识
	Remark    interface{} // 备注
	Status    interface{} // 用户状态（1-正常 0-停用）
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
