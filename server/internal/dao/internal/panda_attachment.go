// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PandaAttachmentDao is the data access object for the table panda_attachment.
type PandaAttachmentDao struct {
	table    string                 // table is the underlying table name of the DAO.
	group    string                 // group is the database configuration group name of the current DAO.
	columns  PandaAttachmentColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler     // handlers for customized model modification.
}

// PandaAttachmentColumns defines and stores column names for the table panda_attachment.
type PandaAttachmentColumns struct {
	Id        string // ID
	Type      string // 文件类型（image doc audio video zip other）
	Name      string // 文件原始名
	Path      string // 文件路径
	Size      string // 文件大小
	Ext       string // 扩展名
	Drive     string // 上传驱动类型 1 本地 2 阿里云 3 腾讯云 4 七牛云
	Status    string // 文件状态（1-正常 0-停用）
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// pandaAttachmentColumns holds the columns for the table panda_attachment.
var pandaAttachmentColumns = PandaAttachmentColumns{
	Id:        "id",
	Type:      "type",
	Name:      "name",
	Path:      "path",
	Size:      "size",
	Ext:       "ext",
	Drive:     "drive",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewPandaAttachmentDao creates and returns a new DAO object for table data access.
func NewPandaAttachmentDao(handlers ...gdb.ModelHandler) *PandaAttachmentDao {
	return &PandaAttachmentDao{
		group:    "default",
		table:    "panda_attachment",
		columns:  pandaAttachmentColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PandaAttachmentDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PandaAttachmentDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PandaAttachmentDao) Columns() PandaAttachmentColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PandaAttachmentDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PandaAttachmentDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *PandaAttachmentDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
