// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PandaRoleMenuDao is the data access object for the table panda_role_menu.
type PandaRoleMenuDao struct {
	table    string               // table is the underlying table name of the DAO.
	group    string               // group is the database configuration group name of the current DAO.
	columns  PandaRoleMenuColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler   // handlers for customized model modification.
}

// PandaRoleMenuColumns defines and stores column names for the table panda_role_menu.
type PandaRoleMenuColumns struct {
	Id        string // 角色菜单关联ID
	RoleId    string // 角色ID
	MenuId    string // 菜单ID
	CreatedAt string // 创建时间
}

// pandaRoleMenuColumns holds the columns for the table panda_role_menu.
var pandaRoleMenuColumns = PandaRoleMenuColumns{
	Id:        "id",
	RoleId:    "role_id",
	MenuId:    "menu_id",
	CreatedAt: "created_at",
}

// NewPandaRoleMenuDao creates and returns a new DAO object for table data access.
func NewPandaRoleMenuDao(handlers ...gdb.ModelHandler) *PandaRoleMenuDao {
	return &PandaRoleMenuDao{
		group:    "default",
		table:    "panda_role_menu",
		columns:  pandaRoleMenuColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PandaRoleMenuDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PandaRoleMenuDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PandaRoleMenuDao) Columns() PandaRoleMenuColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PandaRoleMenuDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PandaRoleMenuDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *PandaRoleMenuDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
