// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PandaAttachment is the golang structure of table panda_attachment for DAO operations like Where/Data.
type PandaAttachment struct {
	g.Meta    `orm:"table:panda_attachment, do:true"`
	Id        interface{} // ID
	Type      interface{} // 文件类型（image doc audio video zip other）
	Name      interface{} // 文件原始名
	Path      interface{} // 文件路径
	Size      interface{} // 文件大小
	Ext       interface{} // 扩展名
	Drive     interface{} // 上传驱动类型 1 本地 2 阿里云 3 腾讯云 4 七牛云
	Status    interface{} // 文件状态（1-正常 0-停用）
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
