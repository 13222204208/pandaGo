// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"server/api/admin/dicttype"
)

type (
	IDictType interface {
		// Create 创建字典类型
		Create(ctx context.Context, req *dicttype.CreateReq) (res *dicttype.CreateRes, err error)
		// Delete 删除字典类型
		Delete(ctx context.Context, req *dicttype.DeleteReq) error
		// Update 更新字典类型
		Update(ctx context.Context, req *dicttype.UpdateReq) error
		// GetList 获取字典类型列表
		GetList(ctx context.Context, req *dicttype.GetListReq) (res *dicttype.GetListRes, err error)
	}
)

var (
	localDictType IDictType
)

func DictType() IDictType {
	if localDictType == nil {
		panic("implement not found for interface IDictType, forgot register?")
	}
	return localDictType
}

func RegisterDictType(i IDictType) {
	localDictType = i
}
