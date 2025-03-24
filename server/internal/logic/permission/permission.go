package permission

import (
	"context"
	"server/internal/dao"
	"server/internal/library/casbin"
	"server/internal/service"
)

type sPermission struct{}

func init() {
	service.RegisterPermission(New())
}

func New() *sPermission {
	return &sPermission{}
}

// AddRolePermission 添加角色权限
func (s *sPermission) AddRolePermission(ctx context.Context, roleCode string, permissions []string) error {
	// 获取 Casbin 执行器
	enforcer, err := casbin.CasbinEnforcer(ctx)
	if err != nil {
		return err
	}

	// 先清除该角色的所有权限
	err = s.RemoveRolePermission(ctx, roleCode)
	if err != nil {
		return err
	}

	// 添加新的权限
	for _, permission := range permissions {
		// 格式: p, roleCode, resource, action
		parts := []string{roleCode, permission, "allow"}
		_, err = enforcer.AddPolicy(parts)
		if err != nil {
			return err
		}
	}

	// 保存策略
	return enforcer.SavePolicy()
}

// RemoveRolePermission 删除角色权限
func (s *sPermission) RemoveRolePermission(ctx context.Context, roleCode string) error {
	enforcer, err := casbin.CasbinEnforcer(ctx)
	if err != nil {
		return err
	}

	// 删除角色的所有权限
	_, err = enforcer.RemoveFilteredPolicy(0, roleCode)
	if err != nil {
		return err
	}

	// 保存策略
	return enforcer.SavePolicy()
}

// CheckPermission 检查用户是否有权限
func (s *sPermission) CheckPermission(ctx context.Context, userId int64, permission string) (bool, error) {
	// 获取用户角色
	userRoles, err := s.getUserRoles(ctx, userId)
	if err != nil {
		return false, err
	}

	enforcer, err := casbin.CasbinEnforcer(ctx)
	if err != nil {
		return false, err
	}

	// 检查每个角色是否有权限
	for _, role := range userRoles {
		hasPermission, err := enforcer.Enforce(role, permission, "allow")
		if err != nil {
			return false, err
		}
		if hasPermission {
			return true, nil
		}
	}

	return false, nil
}

// getUserRoles 获取用户角色
func (s *sPermission) getUserRoles(ctx context.Context, userId int64) ([]string, error) {
	var roles []string

	// 查询用户角色关联
	userRoles, err := dao.PandaUserRole.Ctx(ctx).
		Where(dao.PandaUserRole.Columns().UserId, userId).
		All()
	if err != nil {
		return nil, err
	}

	if len(userRoles) > 0 {
		var roleIds []int64
		for _, record := range userRoles {
			roleIds = append(roleIds, record["role_id"].Int64())
		}

		// 查询角色编码
		roleRecords, err := dao.PandaRole.Ctx(ctx).
			Fields(dao.PandaRole.Columns().Code).
			WhereIn(dao.PandaRole.Columns().Id, roleIds).
			All()
		if err != nil {
			return nil, err
		}

		for _, role := range roleRecords {
			roles = append(roles, role["code"].String())
		}
	}

	return roles, nil
}
