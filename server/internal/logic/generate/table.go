package generate

import (
	"context"
	"fmt"
	"server/api/admin/generate"
	"server/api/common"
	generateLibrary "server/internal/library/generate"
	"server/internal/service"
	"server/utility"

	"github.com/gogf/gf/v2/frame/g"
)

type sGenerate struct{}

func init() {
	service.RegisterGenerate(New())
}

func New() *sGenerate {
	return &sGenerate{}
}

// GetTables 获取数据表列表
func (s *sGenerate) GetTables(ctx context.Context, req *generate.GetTablesReq) (res *generate.GetTablesRes, err error) {
	// 调用库函数获取表元数据
	result, err := generateLibrary.GetTablesMeta("mysql", req.CurrentPage, req.PageSize)
	if err != nil {
		return nil, err
	}
	// r, err := generateLibrary.CreateTableSql("帮我创建一个商品表,里面有要有商品进货时间,商品名称,商品价格,商品库存,商品分类,商品图片,商品描述,商品状态,商品属性,商品标签,商品销量,商品收藏量,商品评价量,商品评价分数,商品评价内容,商品评价时间,商品评价用户,商品评价回复,商品评价回复时间,商品评价回复用户,商品评价回复内容")
	// if err != nil {
	// 	return nil, err
	// }
	// fmt.Println("结果", r)
	// 转换为API响应格式
	res = &generate.GetTablesRes{
		List: make([]*generate.TableInfo, 0, len(result.List)),
		ListRes: common.ListRes{
			Total:       result.Total,
			CurrentPage: req.CurrentPage,
			PageSize:    req.PageSize,
		},
	}

	for _, table := range result.List {
		res.List = append(res.List, &generate.TableInfo{
			TableName:    table.Name,
			TableComment: table.Comment,
			CreateTime:   utility.StrToFormatTime(table.CreateTime),
			UpdateTime:   utility.StrToFormatTime(table.UpdateTime),
		})
	}

	return
}

// GenerateSql 根据描述生成建表SQL
func (s *sGenerate) GenerateSql(ctx context.Context, req *generate.GenerateSqlReq) (res *generate.GenerateSqlRes, err error) {
	sql, err := generateLibrary.CreateTableSql(req.Prompt)
	if err != nil {
		return nil, err
	}
	fmt.Println("结果", sql)

	res = &generate.GenerateSqlRes{
		Sql: sql,
	}
	return
}

// ExecuteSql 执行SQL语句
func (s *sGenerate) ExecuteSql(ctx context.Context, req *generate.ExecuteSqlReq) (err error) {
	// 执行SQL语句
	db := g.DB()
	if _, err := db.Exec(ctx, req.Sql); err != nil {
		return err
	}
	return nil
}
