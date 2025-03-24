// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"server/api/admin/attachment"

	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IAttachment interface {
		// Upload 上传附件
		Upload(ctx context.Context, file []*ghttp.UploadFile) (res *attachment.UploadRes, err error)
		// Delete 删除附件
		Delete(ctx context.Context, req *attachment.DeleteReq) error
		// Update 更新附件名称
		Update(ctx context.Context, req *attachment.UpdateReq) error
		// GetList 获取附件列表
		GetList(ctx context.Context, req *attachment.GetListReq) (res *attachment.GetListRes, err error)
	}
)

var (
	localAttachment IAttachment
)

func Attachment() IAttachment {
	if localAttachment == nil {
		panic("implement not found for interface IAttachment, forgot register?")
	}
	return localAttachment
}

func RegisterAttachment(i IAttachment) {
	localAttachment = i
}
