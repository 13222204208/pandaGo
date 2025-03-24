package admin

import (
	"context"
	"server/internal/service"

	"server/api/admin/dicttype"
)

func (c *ControllerDicttype) Create(ctx context.Context, req *dicttype.CreateReq) (res *dicttype.CreateRes, err error) {
	res, err = service.DictType().Create(ctx, req)
	return
}
func (c *ControllerDicttype) Delete(ctx context.Context, req *dicttype.DeleteReq) (res *dicttype.DeleteRes, err error) {
	err = service.DictType().Delete(ctx, req)
	return
}
func (c *ControllerDicttype) Update(ctx context.Context, req *dicttype.UpdateReq) (res *dicttype.UpdateRes, err error) {
	err = service.DictType().Update(ctx, req)
	return
}
func (c *ControllerDicttype) GetList(ctx context.Context, req *dicttype.GetListReq) (res *dicttype.GetListRes, err error) {
	res, err = service.DictType().GetList(ctx, req)
	return
}
