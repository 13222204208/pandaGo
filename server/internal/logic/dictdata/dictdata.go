package dictdata

import (
	"context"
	"server/api/admin/dictdata"
	"server/api/common"
	"server/internal/dao"
	"server/internal/model/entity"
	"server/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
)

type sDictData struct{}

func init() {
	service.RegisterDictData(New())
}

func New() *sDictData {
	return &sDictData{}
}

// Create 创建字典数据
func (s *sDictData) Create(ctx context.Context, req *dictdata.CreateReq) (res *dictdata.CreateRes, err error) {
	// 检查字典类型是否存在
	count, err := dao.PandaDictType.Ctx(ctx).Where(dao.PandaDictType.Columns().Type, req.Type).Count()
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, gerror.New("字典类型不存在")
	}

	dictData := &entity.PandaDictData{
		Label:     req.Label,
		Value:     req.Value,
		ValueType: req.ValueType,
		Type:      req.Type,
		ListClass: req.ListClass,
		IsDefault: req.IsDefault,
		Sort:      req.Sort,
		Status:    req.Status,
		Remark:    req.Remark,
	}
	_, err = dao.PandaDictData.Ctx(ctx).Insert(dictData)
	return
}

// Delete 删除字典数据
func (s *sDictData) Delete(ctx context.Context, req *dictdata.DeleteReq) error {
	// 检查字典数据是否存在
	info, err := dao.PandaDictData.Ctx(ctx).Where(dao.PandaDictData.Columns().Id, req.Id).One()
	if err != nil {
		return err
	}
	if info == nil {
		return gerror.New("字典数据不存在")
	}

	_, err = dao.PandaDictData.Ctx(ctx).Where(dao.PandaDictData.Columns().Id, req.Id).Delete()
	return err
}

// Update 更新字典数据
func (s *sDictData) Update(ctx context.Context, req *dictdata.UpdateReq) error {
	// 检查字典数据是否存在
	info, err := dao.PandaDictData.Ctx(ctx).Where(dao.PandaDictData.Columns().Id, req.Id).One()
	if err != nil {
		return err
	}
	if info == nil {
		return gerror.New("字典数据不存在")
	}

	// 检查字典类型是否存在
	count, err := dao.PandaDictType.Ctx(ctx).Where(dao.PandaDictType.Columns().Type, req.Type).Count()
	if err != nil {
		return err
	}
	if count == 0 {
		return gerror.New("字典类型不存在")
	}

	// 直接使用请求数据更新
	_, err = dao.PandaDictData.Ctx(ctx).
		Where(dao.PandaDictData.Columns().Id, req.Id).
		Data(req).
		Update()
	return err
}

// GetList 获取字典数据列表
func (s *sDictData) GetList(ctx context.Context, req *dictdata.GetListReq) (res *dictdata.GetListRes, err error) {
	m := dao.PandaDictData.Ctx(ctx)

	// 条件查询
	if req.Type != "" {
		// 根据类型查询字典类型
		dictType, err := dao.PandaDictType.Ctx(ctx).Where(dao.PandaDictType.Columns().Type, req.Type).One()
		if err != nil {
			return nil, err
		}
		if dictType != nil {
			// 如果是父类型(pid=0)，查询所有子类型
			if dictType["pid"].Int64() == 0 {
				// 查询所有子类型
				types, err := dao.PandaDictType.Ctx(ctx).
					Where(dao.PandaDictType.Columns().Pid, dictType["id"]).
					Array(dao.PandaDictType.Columns().Type)
				if err != nil {
					return nil, err
				}
				// 将父类型加入到类型列表
				types = append(types, dictType["type"])
				// 查询这些类型的所有字典数据
				m = m.WhereIn(dao.PandaDictData.Columns().Type, types)
			} else {
				// 如果是子类型，直接查询该类型
				m = m.Where(dao.PandaDictData.Columns().Type, req.Type)
			}
		}
	}
	if req.Label != "" {
		m = m.WhereLike(dao.PandaDictData.Columns().Label, "%"+req.Label+"%")
	}
	if req.Value != "" {
		m = m.WhereLike(dao.PandaDictData.Columns().Value, "%"+req.Value+"%")
	}
	if req.Status != -1 {
		m = m.Where(dao.PandaDictData.Columns().Status, req.Status)
	}

	// 获取总数
	total, err := m.Count()
	if err != nil {
		return nil, err
	}

	// 分页查询
	var list []*entity.PandaDictData
	err = m.Page(req.CurrentPage, req.PageSize).
		OrderAsc(dao.PandaDictData.Columns().Sort).
		Scan(&list)
	if err != nil {
		return nil, err
	}

	return &dictdata.GetListRes{
		ListRes: common.ListRes{
			Total:       total,
			CurrentPage: req.CurrentPage,
			PageSize:    req.PageSize,
		},
		List: list,
	}, nil
}
