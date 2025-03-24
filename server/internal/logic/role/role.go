package role

import (
	"context"
	"fmt"
	"server/api/admin/role"
	"server/api/common"
	"server/internal/dao"
	"server/internal/model/entity"
	"server/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
)

type sRole struct{}

func init() {
	service.RegisterRole(New())
}

func New() *sRole {
	return &sRole{}
}

// Create 创建角色
func (s *sRole) Create(ctx context.Context, req *role.CreateReq) (res *role.CreateRes, err error) {
	// 检查角色编码是否已存在
	count, err := dao.PandaRole.Ctx(ctx).Where(dao.PandaRole.Columns().Code, req.Code).Count()
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, gerror.New("角色编码已存在")
	}

	// 创建角色
	data := &entity.PandaRole{
		Name:   req.Name,
		Code:   req.Code,
		Remark: req.Remark,
		Status: req.Status,
	}
	result, err := dao.PandaRole.Ctx(ctx).Insert(data)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &role.CreateRes{Id: id}, nil
}

// Delete 删除角色
func (s *sRole) Delete(ctx context.Context, req *role.DeleteReq) error {
	// 检查角色是否存在
	info, err := dao.PandaRole.Ctx(ctx).Where(dao.PandaRole.Columns().Id, req.Id).One()
	if err != nil {
		return err
	}
	if info == nil {
		return gerror.New("角色不存在")
	}

	// 检查角色是否被用户使用
	count, err := dao.PandaUserRole.Ctx(ctx).Where(dao.PandaUserRole.Columns().RoleId, req.Id).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return gerror.New("角色已被用户使用，无法删除")
	}

	_, err = dao.PandaRole.Ctx(ctx).Where(dao.PandaRole.Columns().Id, req.Id).Delete()
	return err
}

// Update 更新角色
func (s *sRole) Update(ctx context.Context, req *role.UpdateReq) error {
	// 检查角色是否存在
	info, err := dao.PandaRole.Ctx(ctx).Where(dao.PandaRole.Columns().Id, req.Id).One()
	if err != nil {
		return err
	}
	if info == nil {
		return gerror.New("角色不存在")
	}

	// 如果修改了角色编码，检查是否与其他角色冲突
	if req.Code != "" && req.Code != info["code"].String() {
		count, err := dao.PandaRole.Ctx(ctx).
			Where(dao.PandaRole.Columns().Code, req.Code).
			WhereNot(dao.PandaRole.Columns().Id, req.Id).
			Count()
		if err != nil {
			return err
		}
		if count > 0 {
			return gerror.New("角色编码已存在")
		}
	}

	// 更新角色信息
	data := gconv.Map(req)

	_, err = dao.PandaRole.Ctx(ctx).Where(dao.PandaRole.Columns().Id, req.Id).Update(data)
	return err
}

// SaveMenuIds 角色保存菜单id
func (s *sRole) SaveMenuIds(ctx context.Context, req *role.CreateRoleMenuReq) error {
	// 检查角色是否存在
	info, err := dao.PandaRole.Ctx(ctx).Where(dao.PandaRole.Columns().Id, req.Id).One()
	if err != nil {
		return err
	}
	if info == nil {
		return gerror.New("角色不存在")
	}

	// 开启事务
	err = dao.PandaRoleMenu.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除原有的角色菜单关联
		_, err = tx.Model(dao.PandaRoleMenu.Table()).Where(dao.PandaRoleMenu.Columns().RoleId, req.Id).Delete()
		if err != nil {
			return err
		}
		err := service.Permission().RemoveRolePermission(ctx, gconv.String(req.Id))
		if err != nil {
			return err
		}

		var menuIds []string
		// 批量插入新的角色菜单关联
		if len(req.MenuIds) > 0 {
			var roleMenus []*entity.PandaRoleMenu
			for _, menuId := range req.MenuIds {
				roleMenus = append(roleMenus, &entity.PandaRoleMenu{
					RoleId: req.Id,
					MenuId: menuId,
				})
				menuIds = append(menuIds, gconv.String(menuId))
			}
			_, err = tx.Model(dao.PandaRoleMenu.Table()).Insert(roleMenus)
			if err != nil {
				return err
			}

		}
		err = service.Permission().AddRolePermission(ctx, gconv.String(req.Id), menuIds)
		if err != nil {
			return err
		}
		return nil
	})

	return err
}

// GetList 获取角色列表
func (s *sRole) GetList(ctx context.Context, req *role.GetListReq) (res *role.GetListRes, err error) {
	m := dao.PandaRole.Ctx(ctx)

	// 条件查询
	if req.Name != "" {
		m = m.WhereLike(dao.PandaRole.Columns().Name, "%"+req.Name+"%")
	}
	if req.Code != "" {
		m = m.WhereLike(dao.PandaRole.Columns().Code, "%"+req.Code+"%")
	}
	if req.Status != -1 {
		m = m.Where(dao.PandaRole.Columns().Status, req.Status)
	}
	fmt.Println("status", req.Status)
	// 获取总数
	total, err := m.Count()
	if err != nil {
		return nil, err
	}

	// 分页查询
	var list []*entity.PandaRole
	err = m.Page(req.CurrentPage, req.PageSize).Scan(&list)
	if err != nil {
		return nil, err
	}

	// 转换为响应结构
	var roleList []*entity.PandaRole
	for _, item := range list {
		roleList = append(roleList, &entity.PandaRole{
			Id:     item.Id,
			Name:   item.Name,
			Code:   item.Code,
			Remark: item.Remark,
			Status: item.Status,
		})
	}
	fmt.Println("roleList", roleList)
	return &role.GetListRes{
		ListRes: common.ListRes{
			Total:       total,
			CurrentPage: req.CurrentPage,
			PageSize:    req.PageSize,
		},
		List: roleList,
	}, nil
}

// GetRoleMenuIds 获取角色菜单ID列表
func (s *sRole) GetRoleMenuIds(ctx context.Context, req *role.GetRoleMenuIdsReq) (res *role.GetRoleMenuIdsRes, err error) {
	// 检查角色是否存在
	info, err := dao.PandaRole.Ctx(ctx).Where(dao.PandaRole.Columns().Id, req.Id).One()
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, gerror.New("角色不存在")
	}

	// 查询角色关联的菜单ID
	var roleMenus []*entity.PandaRoleMenu
	err = dao.PandaRoleMenu.Ctx(ctx).
		Where(dao.PandaRoleMenu.Columns().RoleId, req.Id).
		Scan(&roleMenus)
	if err != nil {
		return nil, err
	}

	// 提取菜单ID
	var menuIds []int64
	for _, rm := range roleMenus {
		menuIds = append(menuIds, rm.MenuId)
	}

	return &role.GetRoleMenuIdsRes{
		MenuIds: menuIds,
	}, nil
}
