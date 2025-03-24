package storage

import (
	"context"
	"server/internal/dao"
	"server/internal/model/entity"
	"server/utility"
	"strconv"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
)

// DoAttachmentUpload DoUpload 上传入口
func DoAttachmentUpload(ctx context.Context, files []*ghttp.UploadFile) (err error) {
	if len(files) == 0 {
		err = gerror.New("文件必须!")
		return
	}

	for _, file := range files {
		meta, err := GetFileMeta(file)
		if err != nil {
			return err
		}

		// 生成文件保存路径
		basePath := "uploads/" + gtime.Date()
		file.Filename = strconv.FormatInt(gtime.TimestampNano(), 36) + grand.S(6) + "." + meta.Ext
		_, err = file.Save(basePath)
		if err != nil {
			return err
		}

		data := &entity.PandaAttachment{
			Type:   GetFileKind(meta.Ext),
			Name:   meta.Filename,
			Path:   basePath + "/" + file.Filename,
			Size:   utility.FormatFileSize(meta.Size),
			Ext:    meta.Ext,
			Drive:  1,
			Status: 1,
		}
		_, err = dao.PandaAttachment.Ctx(ctx).Insert(data)
		if err != nil {
			return err
		}
	}

	return nil
}

// GetFileMeta 获取上传文件元数据
func GetFileMeta(file *ghttp.UploadFile) (meta *FileMeta, err error) {
	meta = new(FileMeta)
	meta.Filename = file.Filename
	meta.Size = file.Size
	meta.Ext = Ext(file.Filename)
	meta.Kind = GetFileKind(meta.Ext)
	meta.MimeType, err = GetFileMimeType(meta.Ext)
	if err != nil {
		return
	}

	// 兼容naiveUI
	naiveType := meta.MimeType
	if len(naiveType) == 0 {
		naiveType = "text/plain"
	}
	meta.NaiveType = naiveType

	// 计算md5值
	meta.Md5, err = CalcFileMd5(file)
	return
}
