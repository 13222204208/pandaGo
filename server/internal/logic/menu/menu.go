package menu

import (
	"context"
	"server/api/admin/menu"
	"server/internal/dao"
	"server/internal/model/entity"
	"server/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sMenu struct{}

func init() {
	service.RegisterMenu(New())
}

func New() *sMenu {
	return &sMenu{}
}

// Create 创建菜单
func (s *sMenu) Create(ctx context.Context, req *menu.CreateReq) error {
	// 检查菜单名称是否存在
	count, err := dao.PandaMenu.Ctx(ctx).Where(dao.PandaMenu.Columns().Title, req.Title).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return gerror.New("菜单名称已存在")
	}

	// 如果有父级菜单，检查父级菜单是否存在
	if req.ParentId > 0 {
		count, err = dao.PandaMenu.Ctx(ctx).Where(dao.PandaMenu.Columns().Id, req.ParentId).Count()
		if err != nil {
			return err
		}
		if count == 0 {
			return gerror.New("父级菜单不存在")
		}
	}

	// 创建菜单
	data := &entity.PandaMenu{
		ParentId:        req.ParentId,
		Title:           req.Title,
		Name:            req.Name,
		Path:            req.Path,
		Component:       req.Component,
		Rank:            req.Rank,
		Redirect:        req.Redirect,
		Icon:            req.Icon,
		ExtraIcon:       req.ExtraIcon,
		EnterTransition: req.EnterTransition,
		LeaveTransition: req.LeaveTransition,
		ActivePath:      req.ActivePath,
		Auths:           req.Auths,
		FrameSrc:        req.FrameSrc,
		MenuType:        req.MenuType,
		FrameLoading:    req.FrameLoading,
		KeepAlive:       req.KeepAlive,
		HiddenTag:       req.HiddenTag,
		FixedTag:        req.FixedTag,
		ShowLink:        req.ShowLink,
		ShowParent:      req.ShowParent,
	}

	_, err = dao.PandaMenu.Ctx(ctx).Insert(data)
	return err
}

// Update 更新菜单
func (s *sMenu) Update(ctx context.Context, req *menu.UpdateReq) error {
	// 检查菜单是否存在
	info, err := dao.PandaMenu.Ctx(ctx).Where(dao.PandaMenu.Columns().Id, req.Id).One()
	if err != nil {
		return err
	}
	if info == nil {
		return gerror.New("菜单不存在")
	}

	// 如果修改了菜单名称，检查名称是否存在
	if req.Title != info["title"].String() {
		count, err := dao.PandaMenu.Ctx(ctx).Where(dao.PandaMenu.Columns().Title, req.Title).
			WhereNot(dao.PandaMenu.Columns().Id, req.Id).Count()
		if err != nil {
			return err
		}
		if count > 0 {
			return gerror.New("菜单名称已存在")
		}
	}

	// 如果有父级菜单，检查父级菜单是否存在
	if req.ParentId > 0 {
		count, err := dao.PandaMenu.Ctx(ctx).Where(dao.PandaMenu.Columns().Id, req.ParentId).Count()
		if err != nil {
			return err
		}
		if count == 0 {
			return gerror.New("父级菜单不存在")
		}
	}

	// 不能将自己设为自己的父级
	if req.ParentId == req.Id {
		return gerror.New("不能将自己设为自己的父级")
	}

	// 更新菜单
	data := g.Map{
		dao.PandaMenu.Columns().ParentId:        req.ParentId,
		dao.PandaMenu.Columns().Title:           req.Title,
		dao.PandaMenu.Columns().Name:            req.Name,
		dao.PandaMenu.Columns().Path:            req.Path,
		dao.PandaMenu.Columns().Component:       req.Component,
		dao.PandaMenu.Columns().Rank:            req.Rank,
		dao.PandaMenu.Columns().Redirect:        req.Redirect,
		dao.PandaMenu.Columns().Icon:            req.Icon,
		dao.PandaMenu.Columns().ExtraIcon:       req.ExtraIcon,
		dao.PandaMenu.Columns().EnterTransition: req.EnterTransition,
		dao.PandaMenu.Columns().LeaveTransition: req.LeaveTransition,
		dao.PandaMenu.Columns().ActivePath:      req.ActivePath,
		dao.PandaMenu.Columns().Auths:           req.Auths,
		dao.PandaMenu.Columns().FrameSrc:        req.FrameSrc,
		dao.PandaMenu.Columns().MenuType:        req.MenuType,
		dao.PandaMenu.Columns().FrameLoading:    req.FrameLoading,
		dao.PandaMenu.Columns().KeepAlive:       req.KeepAlive,
		dao.PandaMenu.Columns().HiddenTag:       req.HiddenTag,
		dao.PandaMenu.Columns().FixedTag:        req.FixedTag,
		dao.PandaMenu.Columns().ShowLink:        req.ShowLink,
		dao.PandaMenu.Columns().ShowParent:      req.ShowParent,
	}

	_, err = dao.PandaMenu.Ctx(ctx).Where(dao.PandaMenu.Columns().Id, req.Id).Data(data).Update()
	return err
}

// GetList 获取菜单列表
func (s *sMenu) GetList(ctx context.Context, req *menu.GetListReq) (res *menu.GetListRes, err error) {
	m := dao.PandaMenu.Ctx(ctx)

	// 条件查询
	if req.Title != "" {
		m = m.WhereLike(dao.PandaMenu.Columns().Title, "%"+req.Title+"%")
	}
	if req.MenuType != -1 {
		m = m.Where(dao.PandaMenu.Columns().MenuType, req.MenuType)
	}

	// 查询所有菜单
	var list []*entity.PandaMenu
	err = m.OrderAsc(dao.PandaMenu.Columns().Rank).Scan(&list)
	if err != nil {
		return nil, err
	}
	return &menu.GetListRes{
		List: list,
	}, nil
}

// GetInfo 获取菜单详情
func (s *sMenu) GetInfo(ctx context.Context, req *menu.GetInfoReq) (res *menu.GetInfoRes, err error) {
	info, err := dao.PandaMenu.Ctx(ctx).Where(dao.PandaMenu.Columns().Id, req.Id).One()
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, gerror.New("菜单不存在")
	}

	return &menu.GetInfoRes{
		Info: info,
	}, nil
}

// Delete 删除菜单
func (s *sMenu) Delete(ctx context.Context, req *menu.DeleteReq) error {
	// 检查菜单是否存在
	info, err := dao.PandaMenu.Ctx(ctx).Where(dao.PandaMenu.Columns().Id, req.Id).One()
	if err != nil {
		return err
	}
	if info == nil {
		return gerror.New("菜单不存在")
	}

	// 递归删除子菜单
	err = s.deleteChildren(ctx, req.Id)
	if err != nil {
		return err
	}

	// 删除当前菜单
	_, err = dao.PandaMenu.Ctx(ctx).Where(dao.PandaMenu.Columns().Id, req.Id).Delete()
	return err
}

// deleteChildren 递归删除子菜单
func (s *sMenu) deleteChildren(ctx context.Context, parentId int64) error {
	// 查询所有子菜单
	var children []entity.PandaMenu
	err := dao.PandaMenu.Ctx(ctx).Where(dao.PandaMenu.Columns().ParentId, parentId).Scan(&children)
	if err != nil {
		return err
	}

	// 递归删除每个子菜单的子菜单
	for _, child := range children {
		err = s.deleteChildren(ctx, child.Id)
		if err != nil {
			return err
		}
	}

	// 删除当前层级的子菜单
	if len(children) > 0 {
		_, err = dao.PandaMenu.Ctx(ctx).Where(dao.PandaMenu.Columns().ParentId, parentId).Delete()
		if err != nil {
			return err
		}
	}

	return nil
}

// GetSimpleList 获取简化菜单列表
func (s *sMenu) GetSimpleList(ctx context.Context, req *menu.GetSimpleListReq) (res *menu.GetSimpleListRes, err error) {
	var list []*menu.SimpleMenu
	err = dao.PandaMenu.Ctx(ctx).
		Fields("id,parent_id,menu_type,title").
		OrderAsc(dao.PandaMenu.Columns().Rank).
		Scan(&list)
	if err != nil {
		return nil, err
	}
	return &menu.GetSimpleListRes{
		List: list,
	}, nil
}
