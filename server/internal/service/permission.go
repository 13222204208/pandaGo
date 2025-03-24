// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IPermission interface {
		// AddRolePermission 添加角色权限
		AddRolePermission(ctx context.Context, roleCode string, permissions []string) error
		// RemoveRolePermission 删除角色权限
		RemoveRolePermission(ctx context.Context, roleCode string) error
		// CheckPermission 检查用户是否有权限
		CheckPermission(ctx context.Context, userId int64, permission string) (bool, error)
	}
)

var (
	localPermission IPermission
)

func Permission() IPermission {
	if localPermission == nil {
		panic("implement not found for interface IPermission, forgot register?")
	}
	return localPermission
}

func RegisterPermission(i IPermission) {
	localPermission = i
}
