package admin

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"server/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"

	"server/api/admin/attachment"
)

func (c *ControllerAttachment) Upload(ctx context.Context, req *attachment.UploadReq) (res *attachment.UploadRes, err error) {
	r := g.RequestFromCtx(ctx)
	file := r.GetUploadFiles("file")
	if file == nil {
		err = gerror.New("没有找到上传的文件")
		return
	}

	res, err = service.Attachment().Upload(ctx, file)
	return
}
func (c *ControllerAttachment) Delete(ctx context.Context, req *attachment.DeleteReq) (res *attachment.DeleteRes, err error) {
	err = service.Attachment().Delete(ctx, req)
	return
}
func (c *ControllerAttachment) Update(ctx context.Context, req *attachment.UpdateReq) (res *attachment.UpdateRes, err error) {
	err = service.Attachment().Update(ctx, req)
	return
}
func (c *ControllerAttachment) GetList(ctx context.Context, req *attachment.GetListReq) (res *attachment.GetListRes, err error) {
	res, err = service.Attachment().GetList(ctx, req)
	return
}
