// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PandaUser is the golang structure for table panda_user.
type PandaUser struct {
	Id            int64       `json:"id"            orm:"id"              description:"用户ID"`               // 用户ID
	Username      string      `json:"username"      orm:"username"        description:"用户名"`                // 用户名
	Password      string      `json:"password"      orm:"password"        description:"密码"`                 // 密码
	Nickname      string      `json:"nickname"      orm:"nickname"        description:"昵称"`                 // 昵称
	Avatar        string      `json:"avatar"        orm:"avatar"          description:"头像"`                 // 头像
	Email         string      `json:"email"         orm:"email"           description:"电子邮件"`               // 电子邮件
	Phone         string      `json:"phone"         orm:"phone"           description:"手机号码"`               // 手机号码
	IsDevelop     int         `json:"isDevelop"     orm:"is_develop"      description:"是否开发者账号 1 是  0   否"` // 是否开发者账号 1 是  0   否
	Status        int         `json:"status"        orm:"status"          description:"用户状态（1-正常 0-停用）"`    // 用户状态（1-正常 0-停用）
	LastLoginIp   string      `json:"lastLoginIp"   orm:"last_login_ip"   description:"最后登录ip"`             // 最后登录ip
	LastLoginTime *gtime.Time `json:"lastLoginTime" orm:"last_login_time" description:"最后登录时间"`             // 最后登录时间
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"      description:"创建时间"`               // 创建时间
	UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"      description:"更新时间"`               // 更新时间
	DeletedAt     *gtime.Time `json:"deletedAt"     orm:"deleted_at"      description:"删除时间"`               // 删除时间
}
