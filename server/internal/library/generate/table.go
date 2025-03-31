package generate

import (
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

// TableMeta 表元数据信息结构体
type TableMeta struct {
	Name       string `json:"name"`        // 表名
	Comment    string `json:"comment"`     // 表注释/描述
	CreateTime string `json:"create_time"` // 创建时间
	UpdateTime string `json:"update_time"` // 更新时间
}

// TableMetaResult 表元数据查询结果
type TableMetaResult struct {
	List  []TableMeta `json:"list"`  // 表列表
	Total int         `json:"total"` // 总数量
}

// GetTablesMeta 获取表元数据信息(支持分页)
// dbType: mysql, pgsql, sqlite, mssql
// page: 当前页码(从1开始)
// pageSize: 每页条数
func GetTablesMeta(dbType string, page, pageSize int) (TableMetaResult, error) {
	ctx := gctx.New()
	db := g.DB()
	result := TableMetaResult{
		List:  make([]TableMeta, 0),
		Total: 0,
	}

	// 获取当前数据库名称
	currentDb, err := db.GetValue(ctx, "SELECT DATABASE()")
	if err != nil {
		fmt.Println("获取当前数据库名称出错:", err)
		return result, err
	}

	// 计算总数量
	countQuery := `
        SELECT 
            COUNT(*) as count
        FROM
            information_schema.TABLES
        WHERE
            TABLE_SCHEMA = ?
    `
	countResult, err := db.GetValue(ctx, countQuery, currentDb.String())
	if err != nil {
		fmt.Println("查询总数出错:", err)
		return result, err
	}
	result.Total = countResult.Int()

	// 计算分页参数
	offset := (page - 1) * pageSize
	if offset < 0 {
		offset = 0
	}

	// 查询当前数据库的表信息（带分页）
	query := `
        SELECT
            TABLE_NAME AS table_name,
            TABLE_COMMENT AS table_comment,
            CREATE_TIME AS create_time,
            UPDATE_TIME AS update_time
        FROM
            information_schema.TABLES
        WHERE
            TABLE_SCHEMA = ?
        ORDER BY
            CREATE_TIME DESC, TABLE_NAME ASC
        LIMIT ?, ?
    `
	rows, err := db.GetAll(ctx, query, currentDb.String(), offset, pageSize)
	if err != nil {
		fmt.Println("查询出错:", err)
		return result, err
	}

	// 处理查询结果
	for _, row := range rows {
		tableMeta := TableMeta{
			Name:       row["table_name"].String(),
			Comment:    row["table_comment"].String(),
			CreateTime: row["create_time"].Time().String(),
			UpdateTime: row["update_time"].Time().String(),
		}
		result.List = append(result.List, tableMeta)
	}

	return result, nil
}
