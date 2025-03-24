// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PandaUser is the golang structure of table panda_user for DAO operations like Where/Data.
type PandaUser struct {
	g.Meta        `orm:"table:panda_user, do:true"`
	Id            interface{} // 用户ID
	Username      interface{} // 用户名
	Password      interface{} // 密码
	Nickname      interface{} // 昵称
	Avatar        interface{} // 头像
	Email         interface{} // 电子邮件
	Phone         interface{} // 手机号码
	IsDevelop     interface{} // 是否开发者账号 1 是  0   否
	Status        interface{} // 用户状态（1-正常 0-停用）
	LastLoginIp   interface{} // 最后登录ip
	LastLoginTime *gtime.Time // 最后登录时间
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 更新时间
	DeletedAt     *gtime.Time // 删除时间
}
