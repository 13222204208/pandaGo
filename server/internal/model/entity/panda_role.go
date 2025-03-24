// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PandaRole is the golang structure for table panda_role.
type PandaRole struct {
	Id        int64       `json:"id"        orm:"id"         description:"角色ID"`            // 角色ID
	Name      string      `json:"name"      orm:"name"       description:"角色名"`             // 角色名
	Code      string      `json:"code"      orm:"code"       description:"角色标识"`            // 角色标识
	Remark    string      `json:"remark"    orm:"remark"     description:"备注"`              // 备注
	Status    int         `json:"status"    orm:"status"     description:"用户状态（1-正常 0-停用）"` // 用户状态（1-正常 0-停用）
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`            // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`            // 更新时间
}
