package role

import (
	"server/api/common"
	"server/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

// Role 角色基础信息
type Role struct {
	Name   string `p:"name" v:"required" dc:"角色名称"`
	Code   string `p:"code" v:"required" dc:"角色编码"`
	Remark string `p:"remark" dc:"备注"`
	Status int    `p:"status" v:"required|in:0,1" dc:"状态（0停用 1正常）"`
}

// CreateReq 创建角色
type CreateReq struct {
	g.Meta `path:"/role" method:"post" tags:"Role" summary:"创建角色"`
	Role
}
type CreateRes struct {
	Id int64 `json:"id" dc:"角色ID"`
}

// CreateRoleMenuReq 角色保存菜单id
type CreateRoleMenuReq struct {
	g.Meta  `path:"/role/{id}/menu" method:"put" tags:"Role" summary:"角色保存菜单id"`
	Id      int64   `p:"id" v:"required" dc:"角色ID"`
	MenuIds []int64 `p:"menuIds" v:"required" dc:"菜单ID列表"`
}
type CreateRoleMenuRes struct{}

// DeleteReq 删除角色
type DeleteReq struct {
	g.Meta `path:"/role/{id}" method:"delete" tags:"Role" summary:"删除角色"`
	Id     int64 `p:"id" v:"required" dc:"角色ID"`
}
type DeleteRes struct{}

// UpdateReq 更新角色
type UpdateReq struct {
	g.Meta `path:"/role/{id}" method:"put" tags:"Role" summary:"更新角色"`
	Id     int64 `p:"id" v:"required" dc:"角色ID"`
	Role
}
type UpdateRes struct{}

// GetListReq 获取角色列表
type GetListReq struct {
	g.Meta `path:"/role" method:"get" tags:"Role" summary:"获取角色列表"`
	Name   string `p:"name" dc:"角色名称"`
	Code   string `p:"code" dc:"角色编码"`
	Status int    `p:"status" d:"-1" v:"in:-1,0,1" dc:"状态（-1:全部 0:停用 1:启用）"`
	common.PageReq
}
type GetListRes struct {
	common.ListRes
	List []*entity.PandaRole `json:"list" dc:"角色列表"`
}

// GetRoleMenuIdsReq 获取角色菜单ID列表
type GetRoleMenuIdsReq struct {
	g.Meta `path:"/role/{id}/menu" method:"get" tags:"Role" summary:"获取角色菜单ID列表"`
	Id     int64 `p:"id" v:"required" dc:"角色ID"`
}

// GetRoleMenuIdsRes 获取角色菜单ID列表响应
type GetRoleMenuIdsRes struct {
	MenuIds []int64 `json:"menuIds" dc:"菜单ID列表"`
}
