package user

import (
	"context"
	"fmt"
	"server/api/admin/user"
	"server/internal/dao"
	"server/internal/model/entity"
	"server/internal/service"
	"server/utility"

	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/v2/errors/gerror"
)

type sUser struct{}

func init() {
	service.RegisterUser(New())
}

func New() *sUser {
	return &sUser{}
}

// Create 创建用户
func (s *sUser) Create(ctx context.Context, req *user.CreateReq) (err error) {
	// 检查用户名是否已存在
	count, err := dao.PandaUser.Ctx(ctx).Where(dao.PandaUser.Columns().Username, req.Username).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return gerror.New("用户名已存在")
	}
	//给密码加密
	hashedPassword, err := utility.EncryptPassword(req.Password)
	if err != nil {
		return err
	}
	req.Password = hashedPassword

	//LastLoginIp:   g.RequestFromCtx(ctx).GetClientIp(),
	//LastLoginTime: utility.GetCurrentTime(),
	// 创建用户
	data := &entity.PandaUser{
		Username: req.Username,
		Password: req.Password,
		Nickname: req.Nickname,
		Email:    req.Email,
		Phone:    req.Phone,
		Status:   req.Status,
	}
	_, err = dao.PandaUser.Ctx(ctx).Insert(data)
	return err
}

// Delete 删除用户
func (s *sUser) Delete(ctx context.Context, req *user.DeleteReq) error {
	// 检查用户是否存在
	info, err := dao.PandaUser.Ctx(ctx).Where(dao.PandaUser.Columns().Id, req.Id).One()
	if err != nil {
		return err
	}
	if info == nil {
		return gerror.New("用户不存在")
	}

	_, err = dao.PandaUser.Ctx(ctx).Where(dao.PandaUser.Columns().Id, req.Id).Delete()
	return err
}

// Update 更新用户
func (s *sUser) Update(ctx context.Context, req *user.UpdateReq) error {
	// 检查用户是否存在
	info, err := dao.PandaUser.Ctx(ctx).Where(dao.PandaUser.Columns().Id, req.Id).One()
	if err != nil {
		return err
	}
	if info == nil {
		return gerror.New("用户不存在")
	}

	// 如果修改了用户名，检查是否与其他用户冲突
	if req.Username != "" && req.Username != info["username"].String() {
		count, err := dao.PandaUser.Ctx(ctx).
			Where(dao.PandaUser.Columns().Username, req.Username).
			WhereNot(dao.PandaUser.Columns().Id, req.Id).
			Count()
		if err != nil {
			return err
		}
		if count > 0 {
			return gerror.New("用户名已存在")
		}
	}

	// 更新用户信息
	data := gconv.Map(req)
	_, err = dao.PandaUser.Ctx(ctx).Where(dao.PandaUser.Columns().Id, req.Id).Update(data)
	return err
}

// Reset 重置用户状态,0禁用，1启用,头像,密码,分配角色
func (s *sUser) Reset(ctx context.Context, req *user.ResetReq) error {
	// 检查用户是否存在
	info, err := dao.PandaUser.Ctx(ctx).Where(dao.PandaUser.Columns().Id, req.Id).One()
	if err != nil {
		return err
	}
	if info == nil {
		return gerror.New("用户不存在")
	}

	// 构建更新数据
	data := make(map[string]interface{})

	// 更新状态
	if req.Status != nil {
		data[dao.PandaUser.Columns().Status] = *req.Status
	}

	// 更新头像
	if req.Avatar != "" {
		data[dao.PandaUser.Columns().Avatar] = req.Avatar
	}

	// 重置密码
	if req.Password != "" {
		hashedPassword, err := utility.EncryptPassword(req.Password)
		if err != nil {
			return err
		}
		data[dao.PandaUser.Columns().Password] = hashedPassword
	}

	// 如果有数据需要更新
	if len(data) > 0 {
		_, err = dao.PandaUser.Ctx(ctx).Where(dao.PandaUser.Columns().Id, req.Id).Update(data)
		if err != nil {
			return err
		}
	}

	// 重新分配角色
	if len(req.RoleId) > 0 {
		// 先删除原有角色关联
		_, err = dao.PandaUserRole.Ctx(ctx).Where(dao.PandaUserRole.Columns().UserId, req.Id).Delete()
		if err != nil {
			return err
		}

		// 删除用户现有的所有角色

		// 创建新的角色关联并添加到 Casbin
		var userRoles []*entity.PandaUserRole
		for _, roleId := range req.RoleId {
			// 添加到用户角色关联表
			userRoles = append(userRoles, &entity.PandaUserRole{
				UserId: req.Id,
				RoleId: roleId,
			})

		}

		// 保存用户角色关联
		_, err = dao.PandaUserRole.Ctx(ctx).Insert(userRoles)
		if err != nil {
			return err
		}
	}

	return nil
}

// GetList 获取用户列表
func (s *sUser) GetList(ctx context.Context, req *user.GetListReq) (res *user.GetListRes, err error) {
	m := dao.PandaUser.Ctx(ctx)

	// 条件查询
	if req.Username != "" {
		m = m.WhereLike(dao.PandaUser.Columns().Username, "%"+req.Username+"%")
	}
	if req.Phone != "" {
		m = m.WhereLike(dao.PandaUser.Columns().Phone, "%"+req.Phone+"%")
	}
	if req.Status != 0 {
		m = m.Where(dao.PandaUser.Columns().Status, req.Status)
	}

	// 获取总数
	total, err := m.Count()
	if err != nil {
		return nil, err
	}

	// 分页查询
	var list []*entity.PandaUser
	err = m.Page(req.CurrentPage, req.PageSize).Fields("id,username,nickname,avatar,email,phone,status,created_at,lastLoginTime").Scan(&list)
	if err != nil {
		return nil, err
	}

	return &user.GetListRes{
		List:  list,
		Total: total,
	}, nil
}

// GetUserRoleIds GetRoleIds 获取用户的角色ID
func (s *sUser) GetUserRoleIds(ctx context.Context, req *user.GetUserRoleIdsReq) (res *user.GetUserRoleIdsRes, err error) {
	// 检查用户是否存在
	info, err := dao.PandaUser.Ctx(ctx).Where(dao.PandaUser.Columns().Id, req.Id).One()
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, gerror.New("用户不存在")
	}

	// 查询用户角色关联表
	var userRoles []*entity.PandaUserRole
	err = dao.PandaUserRole.Ctx(ctx).
		Where(dao.PandaUserRole.Columns().UserId, req.Id).
		Scan(&userRoles)
	if err != nil {
		return nil, err
	}

	// 提取角色ID
	var roleIds []int64
	for _, ur := range userRoles {
		roleIds = append(roleIds, ur.RoleId)
	}

	return &user.GetUserRoleIdsRes{
		RoleIds: roleIds,
	}, nil
}

// GetMenuTree 根据角色ID获取菜单树
func (s *sUser) getMenuTree(ctx context.Context, roleIds []int64) ([]*user.MenuTreeItem, error) {
	// 查询角色关联的菜单ID
	var roleMenus []*entity.PandaRoleMenu
	err := dao.PandaRoleMenu.Ctx(ctx).
		WhereIn(dao.PandaRoleMenu.Columns().RoleId, roleIds).
		Scan(&roleMenus)
	if err != nil {
		return nil, err
	}

	// 提取菜单ID
	var menuIds []int64
	for _, rm := range roleMenus {
		menuIds = append(menuIds, rm.MenuId)
	}
	if len(menuIds) == 0 {
		return []*user.MenuTreeItem{}, nil
	}

	// 查询菜单信息
	var menus []*entity.PandaMenu
	err = dao.PandaMenu.Ctx(ctx).
		WhereIn(dao.PandaMenu.Columns().Id, menuIds).
		Order(dao.PandaMenu.Columns().Rank + " ASC").
		Scan(&menus)
	if err != nil {
		return nil, err
	}
	// 构建菜单树
	return s.buildMenuTree(menus, 0)
}

// GetUserMenuTree 根据用户ID获取菜单树
func (s *sUser) GetUserMenuTree(ctx context.Context, req *user.GetUserMenuTreeReq) (res *user.GetUserMenuTreeRes, err error) {
	// 获取当前用户ID
	userId := utility.GetUserId(ctx)
	fmt.Println("userId:", userId)
	if userId == 0 {
		return nil, gerror.New("用户未登录")
	}

	// 查询用户角色关联表
	var userRoles []*entity.PandaUserRole
	err = dao.PandaUserRole.Ctx(ctx).
		Where(dao.PandaUserRole.Columns().UserId, userId).
		Scan(&userRoles)
	if err != nil {
		return nil, err
	}

	// 提取角色ID
	var roleIds []int64
	for _, ur := range userRoles {
		roleIds = append(roleIds, ur.RoleId)
	}
	if len(roleIds) == 0 {
		return &user.GetUserMenuTreeRes{Menu: []*user.MenuTreeItem{}}, nil
	}

	// 根据角色ID获取菜单树
	menuTree, err := s.getMenuTree(ctx, roleIds)
	if err != nil {
		return nil, err
	}

	return &user.GetUserMenuTreeRes{Menu: menuTree}, nil
}

// buildMenuTree 构建菜单树
func (s *sUser) buildMenuTree(menus []*entity.PandaMenu, parentId int64) ([]*user.MenuTreeItem, error) {
	var tree []*user.MenuTreeItem

	for _, m := range menus {
		if m.ParentId == parentId {
			// 查询角色信息
			var roles []string
			// 这里可以根据菜单ID查询关联的角色代码
			// 简化处理，可以在菜单表中添加roles字段存储角色代码

			// 处理菜单显示状态 0:隐藏 1:显示
			showLink := m.ShowLink == 1

			item := &user.MenuTreeItem{
				Path:      m.Path,
				Component: m.Component,
				Name:      m.Name,
				Meta: user.MenuMeta{
					Icon:     m.Icon,
					Title:    m.Title,
					Rank:     m.Rank,
					Roles:    roles,
					ShowLink: showLink,
				},
			}
			// 递归查询子菜单
			children, err := s.buildMenuTree(menus, m.Id)
			if err != nil {
				return nil, err
			}

			if len(children) > 0 {
				item.Children = children
			}

			tree = append(tree, item)
		}
	}
	return tree, nil
}
