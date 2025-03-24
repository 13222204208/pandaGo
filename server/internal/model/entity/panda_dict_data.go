// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PandaDictData is the golang structure for table panda_dict_data.
type PandaDictData struct {
	Id        int64       `json:"id"        orm:"id"         description:"字典数据ID"`                                    // 字典数据ID
	Label     string      `json:"label"     orm:"label"      description:"字典标签"`                                      // 字典标签
	Value     string      `json:"value"     orm:"value"      description:"字典键值"`                                      // 字典键值
	ValueType string      `json:"valueType" orm:"value_type" description:"键值数据类型：string,int,uint,bool,datetime,date"` // 键值数据类型：string,int,uint,bool,datetime,date
	Type      string      `json:"type"      orm:"type"       description:"字典类型"`                                      // 字典类型
	ListClass string      `json:"listClass" orm:"list_class" description:"表格回显样式"`                                    // 表格回显样式
	IsDefault int         `json:"isDefault" orm:"is_default" description:"是否为系统默认"`                                   // 是否为系统默认
	Sort      int         `json:"sort"      orm:"sort"       description:"字典排序"`                                      // 字典排序
	Remark    string      `json:"remark"    orm:"remark"     description:"备注"`                                        // 备注
	Status    int         `json:"status"    orm:"status"     description:"状态"`                                        // 状态
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`                                      // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`                                      // 更新时间
}
