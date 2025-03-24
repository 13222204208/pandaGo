// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"server/api/admin/user"
)

type (
	IUser interface {
		// Create 创建用户
		Create(ctx context.Context, req *user.CreateReq) (err error)
		// Delete 删除用户
		Delete(ctx context.Context, req *user.DeleteReq) error
		// Update 更新用户
		Update(ctx context.Context, req *user.UpdateReq) error
		// Reset 重置用户状态,0禁用，1启用,头像,密码,分配角色
		Reset(ctx context.Context, req *user.ResetReq) error
		// GetList 获取用户列表
		GetList(ctx context.Context, req *user.GetListReq) (res *user.GetListRes, err error)
		// GetUserRoleIds GetRoleIds 获取用户的角色ID
		GetUserRoleIds(ctx context.Context, req *user.GetUserRoleIdsReq) (res *user.GetUserRoleIdsRes, err error)
		// GetUserMenuTree 根据用户ID获取菜单树
		GetUserMenuTree(ctx context.Context, req *user.GetUserMenuTreeReq) (res *user.GetUserMenuTreeRes, err error)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
