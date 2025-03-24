// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PandaCrontab is the golang structure of table panda_crontab for DAO operations like Where/Data.
type PandaCrontab struct {
	g.Meta    `orm:"table:panda_crontab, do:true"`
	Id        interface{} // ID
	Title     interface{} // 任务标题
	Name      interface{} // 任务方法
	Params    interface{} // 函数参数
	Pattern   interface{} // 表达式
	Policy    interface{} // 策略
	Count     interface{} // 执行次数
	Sort      interface{} // 排序
	Remark    interface{} // 备注
	Status    interface{} // 任务状态（1-正常 0-停用）
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
