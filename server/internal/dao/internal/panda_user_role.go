// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PandaUserRoleDao is the data access object for the table panda_user_role.
type PandaUserRoleDao struct {
	table    string               // table is the underlying table name of the DAO.
	group    string               // group is the database configuration group name of the current DAO.
	columns  PandaUserRoleColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler   // handlers for customized model modification.
}

// PandaUserRoleColumns defines and stores column names for the table panda_user_role.
type PandaUserRoleColumns struct {
	Id        string // 用户角色关联ID
	UserId    string // 用户ID
	RoleId    string // 角色ID
	CreatedAt string // 创建时间
}

// pandaUserRoleColumns holds the columns for the table panda_user_role.
var pandaUserRoleColumns = PandaUserRoleColumns{
	Id:        "id",
	UserId:    "user_id",
	RoleId:    "role_id",
	CreatedAt: "created_at",
}

// NewPandaUserRoleDao creates and returns a new DAO object for table data access.
func NewPandaUserRoleDao(handlers ...gdb.ModelHandler) *PandaUserRoleDao {
	return &PandaUserRoleDao{
		group:    "default",
		table:    "panda_user_role",
		columns:  pandaUserRoleColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PandaUserRoleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PandaUserRoleDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PandaUserRoleDao) Columns() PandaUserRoleColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PandaUserRoleDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PandaUserRoleDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *PandaUserRoleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
