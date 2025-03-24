// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PandaCrontabDao is the data access object for the table panda_crontab.
type PandaCrontabDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  PandaCrontabColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// PandaCrontabColumns defines and stores column names for the table panda_crontab.
type PandaCrontabColumns struct {
	Id        string // ID
	Title     string // 任务标题
	Name      string // 任务方法
	Params    string // 函数参数
	Pattern   string // 表达式
	Policy    string // 策略
	Count     string // 执行次数
	Sort      string // 排序
	Remark    string // 备注
	Status    string // 任务状态（1-正常 0-停用）
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// pandaCrontabColumns holds the columns for the table panda_crontab.
var pandaCrontabColumns = PandaCrontabColumns{
	Id:        "id",
	Title:     "title",
	Name:      "name",
	Params:    "params",
	Pattern:   "pattern",
	Policy:    "policy",
	Count:     "count",
	Sort:      "sort",
	Remark:    "remark",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewPandaCrontabDao creates and returns a new DAO object for table data access.
func NewPandaCrontabDao(handlers ...gdb.ModelHandler) *PandaCrontabDao {
	return &PandaCrontabDao{
		group:    "default",
		table:    "panda_crontab",
		columns:  pandaCrontabColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PandaCrontabDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PandaCrontabDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PandaCrontabDao) Columns() PandaCrontabColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PandaCrontabDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PandaCrontabDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *PandaCrontabDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
