// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PandaCrontab is the golang structure for table panda_crontab.
type PandaCrontab struct {
	Id        int64       `json:"id"        orm:"id"         description:"ID"`              // ID
	Title     string      `json:"title"     orm:"title"      description:"任务标题"`            // 任务标题
	Name      string      `json:"name"      orm:"name"       description:"任务方法"`            // 任务方法
	Params    string      `json:"params"    orm:"params"     description:"函数参数"`            // 函数参数
	Pattern   string      `json:"pattern"   orm:"pattern"    description:"表达式"`             // 表达式
	Policy    int64       `json:"policy"    orm:"policy"     description:"策略"`              // 策略
	Count     int64       `json:"count"     orm:"count"      description:"执行次数"`            // 执行次数
	Sort      int         `json:"sort"      orm:"sort"       description:"排序"`              // 排序
	Remark    string      `json:"remark"    orm:"remark"     description:"备注"`              // 备注
	Status    int         `json:"status"    orm:"status"     description:"任务状态（1-正常 0-停用）"` // 任务状态（1-正常 0-停用）
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`            // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`            // 更新时间
}
