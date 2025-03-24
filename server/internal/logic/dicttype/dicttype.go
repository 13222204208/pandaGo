package dicttype

import (
	"context"
	"server/api/admin/dicttype"
	"server/internal/dao"
	"server/internal/model/entity"
	"server/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sDictType struct{}

func init() {
	service.RegisterDictType(New())
}

func New() *sDictType {
	return &sDictType{}
}

// Create 创建字典类型
func (s *sDictType) Create(ctx context.Context, req *dicttype.CreateReq) (res *dicttype.CreateRes, err error) {
	// 检查字典类型是否已存在
	count, err := dao.PandaDictType.Ctx(ctx).Where(dao.PandaDictType.Columns().Type, req.Type).Count()
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, gerror.New("字典类型已存在")
	}

	dictType := &entity.PandaDictType{
		Pid:    req.Pid,
		Name:   req.Name,
		Type:   req.Type,
		Status: req.Status,
		Remark: req.Remark,
	}
	_, err = dao.PandaDictType.Ctx(ctx).Insert(dictType)
	return
}

// Delete 删除字典类型
func (s *sDictType) Delete(ctx context.Context, req *dicttype.DeleteReq) error {
	// 检查字典类型是否存在
	info, err := dao.PandaDictType.Ctx(ctx).Where(dao.PandaDictType.Columns().Id, req.Id).One()
	if err != nil {
		return err
	}
	if info == nil {
		return gerror.New("字典类型不存在")
	}

	// 检查是否存在关联的字典数据
	count, err := dao.PandaDictData.Ctx(ctx).Where(dao.PandaDictData.Columns().Type, info["type"]).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return gerror.New("存在关联的字典数据，不能删除")
	}

	_, err = dao.PandaDictType.Ctx(ctx).Where(dao.PandaDictType.Columns().Id, req.Id).Delete()
	return err
}

// Update 更新字典类型
func (s *sDictType) Update(ctx context.Context, req *dicttype.UpdateReq) error {
	// 检查字典类型是否存在
	info, err := dao.PandaDictType.Ctx(ctx).Where(dao.PandaDictType.Columns().Id, req.Id).One()
	if err != nil {
		return err
	}
	if info == nil {
		return gerror.New("字典类型不存在")
	}

	data := g.Map{}
	if req.Name != "" {
		data[dao.PandaDictType.Columns().Name] = req.Name
	}
	if req.Type != "" {
		// 检查新的字典类型是否与其他记录冲突
		count, err := dao.PandaDictType.Ctx(ctx).
			Where(dao.PandaDictType.Columns().Type, req.Type).
			WhereNot(dao.PandaDictType.Columns().Id, req.Id).
			Count()
		if err != nil {
			return err
		}
		if count > 0 {
			return gerror.New("字典类型已存在")
		}
		data[dao.PandaDictType.Columns().Type] = req.Type
	}
	if req.Status != 0 {
		data[dao.PandaDictType.Columns().Status] = int(req.Status)
	}
	if req.Remark != "" {
		data[dao.PandaDictType.Columns().Remark] = req.Remark
	}

	_, err = dao.PandaDictType.Ctx(ctx).Where(dao.PandaDictType.Columns().Id, req.Id).Update(data)
	return err
}

// GetList 获取字典类型列表
func (s *sDictType) GetList(ctx context.Context, req *dicttype.GetListReq) (res *dicttype.GetListRes, err error) {
	m := dao.PandaDictType.Ctx(ctx)

	var list []*entity.PandaDictType
	err = m.OrderAsc(dao.PandaDictType.Columns().Sort).Scan(&list)
	if err != nil {
		return nil, err
	}

	return &dicttype.GetListRes{List: list}, nil
}
