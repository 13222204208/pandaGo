// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PandaAttachment is the golang structure for table panda_attachment.
type PandaAttachment struct {
	Id        int64       `json:"id"        orm:"id"         description:"ID"`                                    // ID
	Type      string      `json:"type"      orm:"type"       description:"文件类型（image doc audio video zip other）"` // 文件类型（image doc audio video zip other）
	Name      string      `json:"name"      orm:"name"       description:"文件原始名"`                                 // 文件原始名
	Path      string      `json:"path"      orm:"path"       description:"文件路径"`                                  // 文件路径
	Size      string      `json:"size"      orm:"size"       description:"文件大小"`                                  // 文件大小
	Ext       string      `json:"ext"       orm:"ext"        description:"扩展名"`                                   // 扩展名
	Drive     int         `json:"drive"     orm:"drive"      description:"上传驱动类型 1 本地 2 阿里云 3 腾讯云 4 七牛云"`         // 上传驱动类型 1 本地 2 阿里云 3 腾讯云 4 七牛云
	Status    int         `json:"status"    orm:"status"     description:"文件状态（1-正常 0-停用）"`                       // 文件状态（1-正常 0-停用）
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`                                  // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`                                  // 更新时间
}
