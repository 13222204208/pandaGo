// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"server/api/admin/crontab"
)

type (
	ICrontab interface {
		// Create 创建定时任务
		Create(ctx context.Context, req *crontab.CreateReq) error
		// Delete 删除定时任务
		Delete(ctx context.Context, req *crontab.DeleteReq) error
		// Update 更新定时任务
		Update(ctx context.Context, req *crontab.UpdateReq) error
		// GetList 获取定时任务列表
		GetList(ctx context.Context, req *crontab.GetListReq) (res *crontab.GetListRes, err error)
		// ChangeStatus 修改任务状态
		ChangeStatus(ctx context.Context, req *crontab.ChangeStatusReq) error
	}
)

var (
	localCrontab ICrontab
)

func Crontab() ICrontab {
	if localCrontab == nil {
		panic("implement not found for interface ICrontab, forgot register?")
	}
	return localCrontab
}

func RegisterCrontab(i ICrontab) {
	localCrontab = i
}
