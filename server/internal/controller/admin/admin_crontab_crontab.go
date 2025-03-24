package admin

import (
	"context"
	"server/internal/service"

	"server/api/admin/crontab"
)

func (c *ControllerCrontab) Create(ctx context.Context, req *crontab.CreateReq) (res *crontab.CreateRes, err error) {
	err = service.Crontab().Create(ctx, req)
	return
}
func (c *ControllerCrontab) Delete(ctx context.Context, req *crontab.DeleteReq) (res *crontab.DeleteRes, err error) {
	err = service.Crontab().Delete(ctx, req)
	return
}
func (c *ControllerCrontab) Update(ctx context.Context, req *crontab.UpdateReq) (res *crontab.UpdateRes, err error) {
	err = service.Crontab().Update(ctx, req)
	return
}
func (c *ControllerCrontab) GetList(ctx context.Context, req *crontab.GetListReq) (res *crontab.GetListRes, err error) {
	res, err = service.Crontab().GetList(ctx, req)
	return
}
func (c *ControllerCrontab) ChangeStatus(ctx context.Context, req *crontab.ChangeStatusReq) (res *crontab.ChangeStatusRes, err error) {
	err = service.Crontab().ChangeStatus(ctx, req)
	return
}
