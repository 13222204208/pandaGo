// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"server/api/admin/dictdata"
)

type (
	IDictData interface {
		// Create 创建字典数据
		Create(ctx context.Context, req *dictdata.CreateReq) (res *dictdata.CreateRes, err error)
		// Delete 删除字典数据
		Delete(ctx context.Context, req *dictdata.DeleteReq) error
		// Update 更新字典数据
		Update(ctx context.Context, req *dictdata.UpdateReq) error
		// GetList 获取字典数据列表
		GetList(ctx context.Context, req *dictdata.GetListReq) (res *dictdata.GetListRes, err error)
	}
)

var (
	localDictData IDictData
)

func DictData() IDictData {
	if localDictData == nil {
		panic("implement not found for interface IDictData, forgot register?")
	}
	return localDictData
}

func RegisterDictData(i IDictData) {
	localDictData = i
}
