package common

import "github.com/gogf/gf/v2/frame/g"

// EmptyRes 不响应任何数据
type EmptyRes struct {
	g.Meta `mime:"application/json"`
}

// ListRes 列表公共返回
type ListRes struct {
	CurrentPage int `json:"currentPage"` //当前页
	PageSize    int `json:"pageSize"`    //每页数量
	Total       int `json:"total"`       //总数量
}
