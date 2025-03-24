package admin

import (
	"context"
	"server/internal/service"

	"server/api/admin/role"
)

func (c *ControllerRole) Create(ctx context.Context, req *role.CreateReq) (res *role.CreateRes, err error) {
	res, err = service.Role().Create(ctx, req)
	return
}
func (c *ControllerRole) Delete(ctx context.Context, req *role.DeleteReq) (res *role.DeleteRes, err error) {
	err = service.Role().Delete(ctx, req)
	return
}
func (c *ControllerRole) Update(ctx context.Context, req *role.UpdateReq) (res *role.UpdateRes, err error) {
	err = service.Role().Update(ctx, req)
	return
}
func (c *ControllerRole) GetList(ctx context.Context, req *role.GetListReq) (res *role.GetListRes, err error) {
	res, err = service.Role().GetList(ctx, req)
	return
}
func (c *ControllerRole) CreateRoleMenu(ctx context.Context, req *role.CreateRoleMenuReq) (res *role.CreateRoleMenuRes, err error) {
	err = service.Role().SaveMenuIds(ctx, req)
	return
}
func (c *ControllerRole) GetRoleMenuIds(ctx context.Context, req *role.GetRoleMenuIdsReq) (res *role.GetRoleMenuIdsRes, err error) {
	res, err = service.Role().GetRoleMenuIds(ctx, req)
	return
}
