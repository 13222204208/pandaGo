package admin

import (
	"context"
	"server/internal/service"

	"server/api/admin/dictdata"
)

func (c *ControllerDictdata) Create(ctx context.Context, req *dictdata.CreateReq) (res *dictdata.CreateRes, err error) {
	res, err = service.DictData().Create(ctx, req)
	return
}
func (c *ControllerDictdata) Delete(ctx context.Context, req *dictdata.DeleteReq) (res *dictdata.DeleteRes, err error) {
	err = service.DictData().Delete(ctx, req)
	return
}
func (c *ControllerDictdata) Update(ctx context.Context, req *dictdata.UpdateReq) (res *dictdata.UpdateRes, err error) {
	err = service.DictData().Update(ctx, req)
	return
}
func (c *ControllerDictdata) GetList(ctx context.Context, req *dictdata.GetListReq) (res *dictdata.GetListRes, err error) {
	res, err = service.DictData().GetList(ctx, req)
	return
}
