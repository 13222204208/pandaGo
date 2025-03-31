// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"server/api/admin/generate"
)

type (
	IGenerate interface {
		// GetTables 获取数据表列表
		GetTables(ctx context.Context, req *generate.GetTablesReq) (res *generate.GetTablesRes, err error)
		// GenerateSql 根据描述生成建表SQL
		GenerateSql(ctx context.Context, req *generate.GenerateSqlReq) (res *generate.GenerateSqlRes, err error)
		// ExecuteSql 执行SQL语句
		ExecuteSql(ctx context.Context, req *generate.ExecuteSqlReq) (err error)
	}
)

var (
	localGenerate IGenerate
)

func Generate() IGenerate {
	if localGenerate == nil {
		panic("implement not found for interface IGenerate, forgot register?")
	}
	return localGenerate
}

func RegisterGenerate(i IGenerate) {
	localGenerate = i
}
