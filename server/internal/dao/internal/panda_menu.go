// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PandaMenuDao is the data access object for the table panda_menu.
type PandaMenuDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  PandaMenuColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// PandaMenuColumns defines and stores column names for the table panda_menu.
type PandaMenuColumns struct {
	Id              string // 菜单ID
	ParentId        string // 父级ID
	Title           string // 菜单标题
	Name            string // 路由名称
	Path            string // 路由地址
	Component       string // 组件路径
	Rank            string // 排序号
	Redirect        string // 重定向地址
	Icon            string // 图标
	ExtraIcon       string // 额外图标
	EnterTransition string // 进场动画
	LeaveTransition string // 离场动画
	ActivePath      string // 激活路径
	Auths           string // 权限标识
	MenuType        string // 菜单类型（0菜单 1iframe 2外链 3按钮）
	FrameSrc        string // 内嵌iframe地址
	FrameLoading    string // iframe加载状态(0关闭 1开启)
	KeepAlive       string // 缓存路由(0关闭 1开启)
	HiddenTag       string // 隐藏标签(0显示 1隐藏)
	FixedTag        string // 固定标签(0不固定 1固定)
	ShowLink        string // 显示链接(0隐藏 1显示)
	ShowParent      string // 显示父级(0隐藏 1显示)
	Status          string // 状态(0停用 1正常)
	CreatedAt       string // 创建时间
	UpdatedAt       string // 更新时间
}

// pandaMenuColumns holds the columns for the table panda_menu.
var pandaMenuColumns = PandaMenuColumns{
	Id:              "id",
	ParentId:        "parent_id",
	Title:           "title",
	Name:            "name",
	Path:            "path",
	Component:       "component",
	Rank:            "rank",
	Redirect:        "redirect",
	Icon:            "icon",
	ExtraIcon:       "extra_icon",
	EnterTransition: "enter_transition",
	LeaveTransition: "leave_transition",
	ActivePath:      "active_path",
	Auths:           "auths",
	MenuType:        "menu_type",
	FrameSrc:        "frame_src",
	FrameLoading:    "frame_loading",
	KeepAlive:       "keep_alive",
	HiddenTag:       "hidden_tag",
	FixedTag:        "fixed_tag",
	ShowLink:        "show_link",
	ShowParent:      "show_parent",
	Status:          "status",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
}

// NewPandaMenuDao creates and returns a new DAO object for table data access.
func NewPandaMenuDao(handlers ...gdb.ModelHandler) *PandaMenuDao {
	return &PandaMenuDao{
		group:    "default",
		table:    "panda_menu",
		columns:  pandaMenuColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PandaMenuDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PandaMenuDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PandaMenuDao) Columns() PandaMenuColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PandaMenuDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PandaMenuDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *PandaMenuDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
