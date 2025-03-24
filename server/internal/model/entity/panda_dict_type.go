// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PandaDictType is the golang structure for table panda_dict_type.
type PandaDictType struct {
	Id        int64       `json:"id"        orm:"id"         description:"字典类型ID"`            // 字典类型ID
	Pid       int64       `json:"pid"       orm:"pid"        description:"父类字典类型ID"`          // 父类字典类型ID
	Name      string      `json:"name"      orm:"name"       description:"字典类型名称"`            // 字典类型名称
	Type      string      `json:"type"      orm:"type"       description:"字典类型标识（唯一）"`        // 字典类型标识（唯一）
	Sort      int         `json:"sort"      orm:"sort"       description:"排序"`                // 排序
	Remark    string      `json:"remark"    orm:"remark"     description:"备注"`                // 备注
	Status    int         `json:"status"    orm:"status"     description:"字典类型状态（1-正常 0-停用）"` // 字典类型状态（1-正常 0-停用）
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`              // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`              // 更新时间
}
