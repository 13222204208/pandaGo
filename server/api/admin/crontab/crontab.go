package crontab

import (
	"server/api/common"

	"github.com/gogf/gf/v2/frame/g"
)

type Info struct {
	Title   string `p:"title" v:"required" dc:"任务标题"`
	Name    string `p:"name" v:"required" dc:"任务方法名称"`
	Params  string `p:"name"  dc:"任务方法传入的参数"`
	Pattern string `p:"pattern" v:"required" dc:"Cron表达式"`
	Policy  int64  `p:"policy" v:"required" dc:"策略"`
	Sort    int    `p:"sort"  dc:"排序"`
	Status  int    `p:"status" v:"required|in:0,1" dc:"状态(0:禁用 1:启用)"`
	Remark  string `p:"remark" dc:"任务描述"`
}

// CreateReq 创建定时任务请求
type CreateReq struct {
	g.Meta `path:"/crontab" method:"post" tags:"Crontab" summary:"创建定时任务"`
	Info
}

// CreateRes 创建定时任务响应
type CreateRes struct{}

// DeleteReq 删除定时任务请求
type DeleteReq struct {
	g.Meta `path:"/crontab/{id}" method:"delete" tags:"Crontab" summary:"删除定时任务"`
	Id     int64 `p:"id" v:"required" dc:"任务ID"`
}

// DeleteRes 删除定时任务响应
type DeleteRes struct{}

// UpdateReq 更新定时任务请求
type UpdateReq struct {
	g.Meta `path:"/crontab/{id}" method:"put" tags:"Crontab" summary:"更新定时任务"`
	Id     int64 `p:"id" v:"required" dc:"任务ID"`
	Info
}

// UpdateRes 更新定时任务响应
type UpdateRes struct{}

// GetListReq 获取定时任务列表请求
type GetListReq struct {
	g.Meta `path:"/crontab" method:"get" tags:"Crontab" summary:"获取定时任务列表"`
	Name   string `p:"name" dc:"任务名称"`
	Status int    `p:"status" d:"-1" dc:"状态(0:禁用 1:启用 -1:全部)"`
	common.PageReq
}

// GetListRes 获取定时任务列表响应
type GetListRes struct {
	common.ListRes
	List interface{} `json:"list" dc:"定时任务列表"`
}

// ChangeStatusReq 修改任务状态请求
type ChangeStatusReq struct {
	g.Meta `path:"/crontab/{id}/status" method:"put" tags:"Crontab" summary:"修改任务状态"`
	Id     int64 `p:"id" v:"required" dc:"任务ID"`
	Status int   `p:"status" v:"required|in:0,1" dc:"状态(0:禁用 1:启用)"`
}

// ChangeStatusRes 修改任务状态响应
type ChangeStatusRes struct{}
