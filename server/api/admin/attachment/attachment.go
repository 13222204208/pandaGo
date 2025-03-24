package attachment

import (
	"server/api/common"
	"server/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

// UploadReq 上传附件
type UploadReq struct {
	g.Meta `path:"/attachment/upload" method:"post" mime:"multipart/form-data" tags:"Attachment" summary:"上传附件"`
}
type UploadRes struct{}

// DeleteReq 删除附件
type DeleteReq struct {
	g.Meta `path:"/attachment/{id}" method:"delete" tags:"Attachment" summary:"删除附件"`
	Id     int64 `p:"id" v:"required" dc:"附件ID"`
}
type DeleteRes struct{}

// UpdateReq 更新附件名称
type UpdateReq struct {
	g.Meta `path:"/attachment/{id}" method:"put" tags:"Attachment" summary:"更新附件名称"`
	Id     int64  `p:"id" v:"required" dc:"附件ID"`
	Name   string `p:"name"  dc:"文件名称"`
	Status int    `p:"status" d:"-1" v:"in:-1,0,1" dc:"附件状态"`
}
type UpdateRes struct{}

// GetListReq 获取附件列表
type GetListReq struct {
	g.Meta    `path:"/attachment" method:"get" tags:"Attachment" summary:"获取附件列表"`
	Name      string `p:"name" dc:"文件名称"`
	Type      string `p:"type" dc:"文件类型"`
	Drive     int    `p:"drive" dc:"存储驱动"`
	Status    int    `p:"status" d:"-1" v:"in:-1,0,1" dc:"附件状态"`
	StartTime string `p:"startTime" dc:"开始时间"`
	EndTime   string `p:"endTime" dc:"结束时间"`
	common.PageReq
}
type GetListRes struct {
	common.ListRes
	List []*entity.PandaAttachment `json:"list" dc:"附件列表"`
}
