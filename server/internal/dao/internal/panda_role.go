// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PandaRoleDao is the data access object for the table panda_role.
type PandaRoleDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  PandaRoleColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// PandaRoleColumns defines and stores column names for the table panda_role.
type PandaRoleColumns struct {
	Id        string // 角色ID
	Name      string // 角色名
	Code      string // 角色标识
	Remark    string // 备注
	Status    string // 用户状态（1-正常 0-停用）
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// pandaRoleColumns holds the columns for the table panda_role.
var pandaRoleColumns = PandaRoleColumns{
	Id:        "id",
	Name:      "name",
	Code:      "code",
	Remark:    "remark",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewPandaRoleDao creates and returns a new DAO object for table data access.
func NewPandaRoleDao(handlers ...gdb.ModelHandler) *PandaRoleDao {
	return &PandaRoleDao{
		group:    "default",
		table:    "panda_role",
		columns:  pandaRoleColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PandaRoleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PandaRoleDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PandaRoleDao) Columns() PandaRoleColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PandaRoleDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PandaRoleDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *PandaRoleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
