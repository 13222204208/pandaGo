// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"server/api/admin/role"
)

type (
	IRole interface {
		// Create 创建角色
		Create(ctx context.Context, req *role.CreateReq) (res *role.CreateRes, err error)
		// Delete 删除角色
		Delete(ctx context.Context, req *role.DeleteReq) error
		// Update 更新角色
		Update(ctx context.Context, req *role.UpdateReq) error
		// SaveMenuIds 角色保存菜单id
		SaveMenuIds(ctx context.Context, req *role.CreateRoleMenuReq) error
		// GetList 获取角色列表
		GetList(ctx context.Context, req *role.GetListReq) (res *role.GetListRes, err error)
		// GetRoleMenuIds 获取角色菜单ID列表
		GetRoleMenuIds(ctx context.Context, req *role.GetRoleMenuIdsReq) (res *role.GetRoleMenuIdsRes, err error)
	}
)

var (
	localRole IRole
)

func Role() IRole {
	if localRole == nil {
		panic("implement not found for interface IRole, forgot register?")
	}
	return localRole
}

func RegisterRole(i IRole) {
	localRole = i
}
