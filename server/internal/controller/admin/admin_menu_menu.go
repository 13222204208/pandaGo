package admin

import (
	"context"
	"server/internal/service"

	"server/api/admin/menu"
)

func (c *ControllerMenu) Create(ctx context.Context, req *menu.CreateReq) (res *menu.CreateRes, err error) {
	err = service.Menu().Create(ctx, req)
	return
}
func (c *ControllerMenu) Delete(ctx context.Context, req *menu.DeleteReq) (res *menu.DeleteRes, err error) {
	err = service.Menu().Delete(ctx, req)
	return
}
func (c *ControllerMenu) Update(ctx context.Context, req *menu.UpdateReq) (res *menu.UpdateRes, err error) {
	err = service.Menu().Update(ctx, req)
	return
}
func (c *ControllerMenu) GetList(ctx context.Context, req *menu.GetListReq) (res *menu.GetListRes, err error) {
	res, err = service.Menu().GetList(ctx, req)
	return
}
func (c *ControllerMenu) GetInfo(ctx context.Context, req *menu.GetInfoReq) (res *menu.GetInfoRes, err error) {
	res, err = service.Menu().GetInfo(ctx, req)
	return
}
func (c *ControllerMenu) GetSimpleList(ctx context.Context, req *menu.GetSimpleListReq) (res *menu.GetSimpleListRes, err error) {
	res, err = service.Menu().GetSimpleList(ctx, req)
	return
}
