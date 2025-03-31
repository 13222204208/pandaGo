package generate

import (
	"server/api/common"

	"github.com/gogf/gf/v2/frame/g"
)

// TableInfo 表信息结构
type TableInfo struct {
	TableName    string `json:"tableName" dc:"表名"`
	TableComment string `json:"tableComment" dc:"表描述"`
	CreateTime   string `json:"createTime" dc:"创建时间"`
	UpdateTime   string `json:"updateTime" dc:"修改时间"`
}

// GetTablesReq 获取数据表列表请求
type GetTablesReq struct {
	g.Meta `path:"/table" method:"get" tags:"Generate" summary:"获取数据表列表"`
	common.PageReq
}

// GetTablesRes 获取数据表列表响应
type GetTablesRes struct {
	List []*TableInfo `json:"list" dc:"数据表列表"`
	common.ListRes
}

// TableColumn 表字段结构
type TableColumn struct {
	ColumnName    string `json:"columnName" dc:"字段名"`
	DataType      string `json:"dataType" dc:"数据类型"`
	ColumnComment string `json:"columnComment" dc:"字段描述"`
	IsNullable    string `json:"isNullable" dc:"是否允许为空"`
	ColumnKey     string `json:"columnKey" dc:"字段键"`
	ColumnDefault string `json:"columnDefault" dc:"默认值"`
	Extra         string `json:"extra" dc:"额外信息"`
}

// GetTableColumnsReq 获取数据表字段请求
type GetTableColumnsReq struct {
	g.Meta    `path:"/table/columns" method:"get" tags:"Generate" summary:"获取数据表字段"`
	TableName string `p:"tableName" v:"required" dc:"表名"`
}

// GetTableColumnsRes 获取数据表字段响应
type GetTableColumnsRes struct {
	List []*TableColumn `json:"list" dc:"字段列表"`
}

// GenerateSqlReq 根据一句话生成创建表的sql语句
type GenerateSqlReq struct {
	g.Meta `path:"/generate/sql" method:"post" tags:"Generate" summary:"根据一句话生成创建表的sql语句"`
	Prompt string `p:"prompt" v:"required" dc:"一句话描述"`
}

// GenerateSqlRes 根据一句话生成创建表的sql语句
type GenerateSqlRes struct {
	Sql string `json:"sql" dc:"创建表的sql语句"`
}

// ExecuteSqlReq 执行sql语句
type ExecuteSqlReq struct {
	g.Meta `path:"/generate/execute" method:"post" tags:"Generate" summary:"执行sql语句"`
	Sql    string `p:"sql" v:"required" dc:"sql语句"`
}

// 执行sql语句
type ExecuteSqlRes struct {
}
