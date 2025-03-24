// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PandaDictType is the golang structure of table panda_dict_type for DAO operations like Where/Data.
type PandaDictType struct {
	g.Meta    `orm:"table:panda_dict_type, do:true"`
	Id        interface{} // 字典类型ID
	Pid       interface{} // 父类字典类型ID
	Name      interface{} // 字典类型名称
	Type      interface{} // 字典类型标识（唯一）
	Sort      interface{} // 排序
	Remark    interface{} // 备注
	Status    interface{} // 字典类型状态（1-正常 0-停用）
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
