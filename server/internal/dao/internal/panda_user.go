// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PandaUserDao is the data access object for the table panda_user.
type PandaUserDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  PandaUserColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// PandaUserColumns defines and stores column names for the table panda_user.
type PandaUserColumns struct {
	Id            string // 用户ID
	Username      string // 用户名
	Password      string // 密码
	Nickname      string // 昵称
	Avatar        string // 头像
	Email         string // 电子邮件
	Phone         string // 手机号码
	IsDevelop     string // 是否开发者账号 1 是  0   否
	Status        string // 用户状态（1-正常 0-停用）
	LastLoginIp   string // 最后登录ip
	LastLoginTime string // 最后登录时间
	CreatedAt     string // 创建时间
	UpdatedAt     string // 更新时间
	DeletedAt     string // 删除时间
}

// pandaUserColumns holds the columns for the table panda_user.
var pandaUserColumns = PandaUserColumns{
	Id:            "id",
	Username:      "username",
	Password:      "password",
	Nickname:      "nickname",
	Avatar:        "avatar",
	Email:         "email",
	Phone:         "phone",
	IsDevelop:     "is_develop",
	Status:        "status",
	LastLoginIp:   "last_login_ip",
	LastLoginTime: "last_login_time",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
	DeletedAt:     "deleted_at",
}

// NewPandaUserDao creates and returns a new DAO object for table data access.
func NewPandaUserDao(handlers ...gdb.ModelHandler) *PandaUserDao {
	return &PandaUserDao{
		group:    "default",
		table:    "panda_user",
		columns:  pandaUserColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PandaUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PandaUserDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PandaUserDao) Columns() PandaUserColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PandaUserDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PandaUserDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *PandaUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
