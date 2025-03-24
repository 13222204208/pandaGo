package admin

import (
	"context"
	"server/internal/service"

	"server/api/admin/user"
)

func (c *ControllerUser) Create(ctx context.Context, req *user.CreateReq) (res *user.CreateRes, err error) {
	err = service.User().Create(ctx, req)
	return
}
func (c *ControllerUser) Delete(ctx context.Context, req *user.DeleteReq) (res *user.DeleteRes, err error) {
	err = service.User().Delete(ctx, req)
	return
}
func (c *ControllerUser) Update(ctx context.Context, req *user.UpdateReq) (res *user.UpdateRes, err error) {
	err = service.User().Update(ctx, req)
	return
}
func (c *ControllerUser) GetList(ctx context.Context, req *user.GetListReq) (res *user.GetListRes, err error) {
	res, err = service.User().GetList(ctx, req)
	return
}
func (c *ControllerUser) Reset(ctx context.Context, req *user.ResetReq) (res *user.ResetRes, err error) {
	err = service.User().Reset(ctx, req)
	return
}
func (c *ControllerUser) GetUserRoleIds(ctx context.Context, req *user.GetUserRoleIdsReq) (res *user.GetUserRoleIdsRes, err error) {
	res, err = service.User().GetUserRoleIds(ctx, req)
	return
}
func (c *ControllerUser) GetUserMenuTree(ctx context.Context, req *user.GetUserMenuTreeReq) (res *user.GetUserMenuTreeRes, err error) {
	res, err = service.User().GetUserMenuTree(ctx, req)
	return
}
