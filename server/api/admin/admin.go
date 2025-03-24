// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package admin

import (
	"context"

	"server/api/admin/attachment"
	"server/api/admin/crontab"
	"server/api/admin/dictdata"
	"server/api/admin/dicttype"
	"server/api/admin/login"
	"server/api/admin/menu"
	"server/api/admin/role"
	"server/api/admin/user"
)

type IAdminAttachment interface {
	Upload(ctx context.Context, req *attachment.UploadReq) (res *attachment.UploadRes, err error)
	Delete(ctx context.Context, req *attachment.DeleteReq) (res *attachment.DeleteRes, err error)
	Update(ctx context.Context, req *attachment.UpdateReq) (res *attachment.UpdateRes, err error)
	GetList(ctx context.Context, req *attachment.GetListReq) (res *attachment.GetListRes, err error)
}

type IAdminCrontab interface {
	Create(ctx context.Context, req *crontab.CreateReq) (res *crontab.CreateRes, err error)
	Delete(ctx context.Context, req *crontab.DeleteReq) (res *crontab.DeleteRes, err error)
	Update(ctx context.Context, req *crontab.UpdateReq) (res *crontab.UpdateRes, err error)
	GetList(ctx context.Context, req *crontab.GetListReq) (res *crontab.GetListRes, err error)
	ChangeStatus(ctx context.Context, req *crontab.ChangeStatusReq) (res *crontab.ChangeStatusRes, err error)
}

type IAdminDictdata interface {
	Create(ctx context.Context, req *dictdata.CreateReq) (res *dictdata.CreateRes, err error)
	Delete(ctx context.Context, req *dictdata.DeleteReq) (res *dictdata.DeleteRes, err error)
	Update(ctx context.Context, req *dictdata.UpdateReq) (res *dictdata.UpdateRes, err error)
	GetList(ctx context.Context, req *dictdata.GetListReq) (res *dictdata.GetListRes, err error)
}

type IAdminDicttype interface {
	Create(ctx context.Context, req *dicttype.CreateReq) (res *dicttype.CreateRes, err error)
	Delete(ctx context.Context, req *dicttype.DeleteReq) (res *dicttype.DeleteRes, err error)
	Update(ctx context.Context, req *dicttype.UpdateReq) (res *dicttype.UpdateRes, err error)
	GetList(ctx context.Context, req *dicttype.GetListReq) (res *dicttype.GetListRes, err error)
}

type IAdminLogin interface {
	Login(ctx context.Context, req *login.LoginReq) (res *login.LoginRes, err error)
	Captcha(ctx context.Context, req *login.CaptchaReq) (res *login.CaptchaRes, err error)
	RefreshToken(ctx context.Context, req *login.RefreshTokenReq) (res *login.RefreshTokenRes, err error)
}

type IAdminMenu interface {
	Create(ctx context.Context, req *menu.CreateReq) (res *menu.CreateRes, err error)
	Delete(ctx context.Context, req *menu.DeleteReq) (res *menu.DeleteRes, err error)
	Update(ctx context.Context, req *menu.UpdateReq) (res *menu.UpdateRes, err error)
	GetList(ctx context.Context, req *menu.GetListReq) (res *menu.GetListRes, err error)
	GetInfo(ctx context.Context, req *menu.GetInfoReq) (res *menu.GetInfoRes, err error)
	GetSimpleList(ctx context.Context, req *menu.GetSimpleListReq) (res *menu.GetSimpleListRes, err error)
}

type IAdminRole interface {
	Create(ctx context.Context, req *role.CreateReq) (res *role.CreateRes, err error)
	CreateRoleMenu(ctx context.Context, req *role.CreateRoleMenuReq) (res *role.CreateRoleMenuRes, err error)
	Delete(ctx context.Context, req *role.DeleteReq) (res *role.DeleteRes, err error)
	Update(ctx context.Context, req *role.UpdateReq) (res *role.UpdateRes, err error)
	GetList(ctx context.Context, req *role.GetListReq) (res *role.GetListRes, err error)
	GetRoleMenuIds(ctx context.Context, req *role.GetRoleMenuIdsReq) (res *role.GetRoleMenuIdsRes, err error)
}

type IAdminUser interface {
	Create(ctx context.Context, req *user.CreateReq) (res *user.CreateRes, err error)
	Delete(ctx context.Context, req *user.DeleteReq) (res *user.DeleteRes, err error)
	Update(ctx context.Context, req *user.UpdateReq) (res *user.UpdateRes, err error)
	Reset(ctx context.Context, req *user.ResetReq) (res *user.ResetRes, err error)
	GetList(ctx context.Context, req *user.GetListReq) (res *user.GetListRes, err error)
	GetUserRoleIds(ctx context.Context, req *user.GetUserRoleIdsReq) (res *user.GetUserRoleIdsRes, err error)
	GetUserMenuTree(ctx context.Context, req *user.GetUserMenuTreeReq) (res *user.GetUserMenuTreeRes, err error)
}
