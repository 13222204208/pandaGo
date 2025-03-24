package crontab

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
	"server/api/admin/crontab"
	"server/api/common"
	"server/internal/consts"
	"server/internal/dao"
	crontabLib "server/internal/library/crontab"
	"server/internal/model/entity"
	"server/internal/service"
	"server/utility"
)

type sCrontab struct{}

func init() {
	service.RegisterCrontab(New())
}

func New() *sCrontab {
	return &sCrontab{}
}

// Create 创建定时任务
func (s *sCrontab) Create(ctx context.Context, req *crontab.CreateReq) error {
	// 检查任务名称是否存在
	count, err := dao.PandaCrontab.Ctx(ctx).Where(dao.PandaCrontab.Columns().Name, req.Name).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return gerror.New("任务方法名称已存在")
	}

	// 创建任务
	data := &entity.PandaCrontab{
		Title:   req.Title,
		Name:    req.Name,
		Params:  req.Params,
		Pattern: req.Pattern,
		Policy:  req.Policy,
		Sort:    req.Sort,
		Status:  req.Status,
		Remark:  req.Remark,
	}

	_, err = dao.PandaCrontab.Ctx(ctx).Insert(data)
	if err != nil {
		return err
	}

	// 如果状态为启用，则添加到定时任务
	if req.Status == consts.StatusEnabled {
		utility.SafeGo(ctx, func(ctx context.Context) {
			_ = crontabLib.Start(data)
		})
	}

	return nil
}

// Delete 删除定时任务
func (s *sCrontab) Delete(ctx context.Context, req *crontab.DeleteReq) error {
	// 检查任务是否存在
	info, err := dao.PandaCrontab.Ctx(ctx).Where(dao.PandaCrontab.Columns().Id, req.Id).One()
	if err != nil {
		return err
	}
	if info == nil {
		return gerror.New("任务不存在")
	}

	// 如果任务正在运行，先停止
	if info["status"].Int() == 1 {
		gcron.Remove(info["name"].String())
	}

	// 删除任务
	_, err = dao.PandaCrontab.Ctx(ctx).Where(dao.PandaCrontab.Columns().Id, req.Id).Delete()
	return err
}

// Update 更新定时任务
func (s *sCrontab) Update(ctx context.Context, req *crontab.UpdateReq) error {
	// 检查任务是否存在
	info, err := dao.PandaCrontab.Ctx(ctx).Where(dao.PandaCrontab.Columns().Id, req.Id).One()
	if err != nil {
		return err
	}
	if info == nil {
		return gerror.New("任务不存在")
	}

	// 如果修改了名称，检查名称是否存在
	if req.Name != info["name"].String() {
		count, err := dao.PandaCrontab.Ctx(ctx).Where(dao.PandaCrontab.Columns().Name, req.Name).
			WhereNot(dao.PandaCrontab.Columns().Id, req.Id).Count()
		if err != nil {
			return err
		}
		if count > 0 {
			return gerror.New("任务名称已存在")
		}
	}

	// 如果任务正在运行，先停止
	if info["status"].Int() == 1 {
		gcron.Remove(info["name"].String())
	}

	// 更新任务
	data := g.Map{
		dao.PandaCrontab.Columns().Title:   req.Title,
		dao.PandaCrontab.Columns().Name:    req.Name,
		dao.PandaCrontab.Columns().Params:  req.Params,
		dao.PandaCrontab.Columns().Pattern: req.Pattern,
		dao.PandaCrontab.Columns().Policy:  req.Policy,
		dao.PandaCrontab.Columns().Sort:    req.Sort,
		dao.PandaCrontab.Columns().Status:  req.Status,
		dao.PandaCrontab.Columns().Remark:  req.Remark,
	}

	_, err = dao.PandaCrontab.Ctx(ctx).Where(dao.PandaCrontab.Columns().Id, req.Id).Data(data).Update()
	if err != nil {
		return err
	}

	// 如果状态为启用，则添加到定时任务
	if req.Status == 1 {
		task := &entity.PandaCrontab{
			Id:      req.Id,
			Title:   req.Title,
			Name:    req.Name,
			Params:  req.Params,
			Pattern: req.Pattern,
			Policy:  req.Policy,
			Sort:    req.Sort,
			Status:  req.Status,
			Remark:  req.Remark,
		}
		utility.SafeGo(ctx, func(ctx context.Context) {
			_ = crontabLib.Start(task)
		})
	}

	return nil
}

// GetList 获取定时任务列表
func (s *sCrontab) GetList(ctx context.Context, req *crontab.GetListReq) (res *crontab.GetListRes, err error) {
	m := dao.PandaCrontab.Ctx(ctx)

	// 条件查询
	if req.Name != "" {
		m = m.WhereLike(dao.PandaCrontab.Columns().Name, "%"+req.Name+"%")
	}
	if req.Status != -1 {
		m = m.Where(dao.PandaCrontab.Columns().Status, req.Status)
	}

	// 获取总数
	total, err := m.Count()
	if err != nil {
		return nil, err
	}

	// 分页查询
	var list []*entity.PandaCrontab
	err = m.Page(req.CurrentPage, req.PageSize).
		OrderDesc(dao.PandaCrontab.Columns().Id).
		Scan(&list)
	if err != nil {
		return nil, err
	}

	return &crontab.GetListRes{
		ListRes: common.ListRes{
			Total:       total,
			CurrentPage: req.CurrentPage,
			PageSize:    req.PageSize,
		},
		List: list,
	}, nil
}

// ChangeStatus 修改任务状态
func (s *sCrontab) ChangeStatus(ctx context.Context, req *crontab.ChangeStatusReq) error {
	// 检查任务是否存在
	info, err := dao.PandaCrontab.Ctx(ctx).Where(dao.PandaCrontab.Columns().Id, req.Id).One()
	if err != nil {
		return err
	}
	if info == nil {
		return gerror.New("任务不存在")
	}

	// 如果状态相同，不做处理
	if info["status"].Int() == req.Status {
		return nil
	}

	// 更新状态
	_, err = dao.PandaCrontab.Ctx(ctx).Where(dao.PandaCrontab.Columns().Id, req.Id).
		Data(g.Map{dao.PandaCrontab.Columns().Status: req.Status}).Update()
	if err != nil {
		return err
	}

	// 根据状态处理定时任务
	if req.Status == consts.StatusEnabled {
		// 启用任务
		task := &entity.PandaCrontab{}
		err = dao.PandaCrontab.Ctx(ctx).Where(dao.PandaCrontab.Columns().Id, req.Id).Scan(task)
		if err != nil {
			return err
		}
		utility.SafeGo(ctx, func(ctx context.Context) {
			_ = crontabLib.Start(task)
		})
	} else {
		// 停用任务
		gcron.Remove(info["name"].String())
	}

	return nil
}
