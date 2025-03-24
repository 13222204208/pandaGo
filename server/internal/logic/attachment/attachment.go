package attachment

import (
	"context"
	"fmt"
	"os"
	"server/api/admin/attachment"
	"server/api/common"
	"server/internal/dao"
	"server/internal/library/storage"
	"server/internal/model/entity"
	"server/internal/service"
	"time"

	"github.com/gogf/gf/v2/net/ghttp"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sAttachment struct{}

func init() {
	service.RegisterAttachment(New())
}

func New() *sAttachment {
	return &sAttachment{}
}

// Upload 上传附件
func (s *sAttachment) Upload(ctx context.Context, file []*ghttp.UploadFile) (res *attachment.UploadRes, err error) {
	err = storage.DoAttachmentUpload(ctx, file)
	return
}

// Delete 删除附件
func (s *sAttachment) Delete(ctx context.Context, req *attachment.DeleteReq) error {
	// 检查附件是否存在
	info, err := dao.PandaAttachment.Ctx(ctx).Where(dao.PandaAttachment.Columns().Id, req.Id).One()
	if err != nil {
		return err
	}
	if info == nil {
		return gerror.New("附件不存在")
	}

	// 删除数据库记录
	_, err = dao.PandaAttachment.Ctx(ctx).Where(dao.PandaAttachment.Columns().Id, req.Id).Delete()
	if err == nil {
		if err = os.Remove(info["path"].String()); err != nil {
			return gerror.New("文件删除失败")
		}
	}
	return err
}

// Update 更新附件名称
func (s *sAttachment) Update(ctx context.Context, req *attachment.UpdateReq) error {
	// 检查附件是否存在
	info, err := dao.PandaAttachment.Ctx(ctx).Where(dao.PandaAttachment.Columns().Id, req.Id).One()
	if err != nil {
		return err
	}
	if info == nil {
		return gerror.New("附件不存在")
	}

	// 构建更新数据
	data := g.Map{}

	// 更新名称
	if req.Name != "" {
		data[dao.PandaAttachment.Columns().Name] = req.Name
	}
	fmt.Print("状态", req.Status)
	// 更新状态
	if req.Status != -1 {
		data[dao.PandaAttachment.Columns().Status] = req.Status
	}

	// 执行更新
	_, err = dao.PandaAttachment.Ctx(ctx).
		Where(dao.PandaAttachment.Columns().Id, req.Id).
		Data(data).
		Update()
	return err
}

// GetList 获取附件列表
func (s *sAttachment) GetList(ctx context.Context, req *attachment.GetListReq) (res *attachment.GetListRes, err error) {
	m := dao.PandaAttachment.Ctx(ctx)

	// 条件查询
	if req.Name != "" {
		m = m.WhereLike(dao.PandaAttachment.Columns().Name, "%"+req.Name+"%")
	}
	if req.Type != "" {
		m = m.Where(dao.PandaAttachment.Columns().Type, req.Type)
	}
	if req.Status != -1 {
		m = m.Where(dao.PandaAttachment.Columns().Status, req.Status)
	}
	if req.Drive != 0 {
		m = m.Where(dao.PandaAttachment.Columns().Drive, req.Drive)
	}

	// 查询日期范围
	if req.StartTime != "" || req.EndTime != "" {
		timeLayout := "2006-01-02"
		if req.StartTime != "" && req.EndTime != "" {
			startTime, _ := time.Parse(timeLayout, req.StartTime)
			endTime, _ := time.Parse(timeLayout, req.EndTime)
			m = m.Where("DATE(created_at) BETWEEN ? AND ?",
				startTime.Format(timeLayout),
				endTime.Format(timeLayout))
		} else if req.StartTime != "" {
			startTime, _ := time.Parse(timeLayout, req.StartTime)
			m = m.Where("DATE(created_at) >= ?", startTime.Format(timeLayout))
		} else if req.EndTime != "" {
			endTime, _ := time.Parse(timeLayout, req.EndTime)
			m = m.Where("DATE(created_at) <= ?", endTime.Format(timeLayout))
		}
	}

	// 获取总数
	total, err := m.Count()
	if err != nil {
		return nil, err
	}

	// 分页查询
	var list []*entity.PandaAttachment
	
	// 使用原始查询构建器，确保所有条件都被应用
	err = m.Page(req.CurrentPage, req.PageSize).
		OrderDesc(dao.PandaAttachment.Columns().Id).
		Scan(&list)
	if err != nil {
		return nil, err
	}

	// 直接使用查询结果，不需要再次转换
	return &attachment.GetListRes{
		ListRes: common.ListRes{
			Total:       total,
			CurrentPage: req.CurrentPage,
			PageSize:    req.PageSize,
		},
		List: list,
	}, nil
}
