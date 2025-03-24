// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"server/api/admin/menu"
)

type (
	IMenu interface {
		// Create 创建菜单
		Create(ctx context.Context, req *menu.CreateReq) error
		// Update 更新菜单
		Update(ctx context.Context, req *menu.UpdateReq) error
		// GetList 获取菜单列表
		GetList(ctx context.Context, req *menu.GetListReq) (res *menu.GetListRes, err error)
		// GetInfo 获取菜单详情
		GetInfo(ctx context.Context, req *menu.GetInfoReq) (res *menu.GetInfoRes, err error)
		// Delete 删除菜单
		Delete(ctx context.Context, req *menu.DeleteReq) error
		// GetSimpleList 获取简化菜单列表
		GetSimpleList(ctx context.Context, req *menu.GetSimpleListReq) (res *menu.GetSimpleListRes, err error)
	}
)

var (
	localMenu IMenu
)

func Menu() IMenu {
	if localMenu == nil {
		panic("implement not found for interface IMenu, forgot register?")
	}
	return localMenu
}

func RegisterMenu(i IMenu) {
	localMenu = i
}
