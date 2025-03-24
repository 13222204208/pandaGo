package dicttype

import (
	"server/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type DictType struct {
	Pid    int64  `p:"pid"  dc:"父字典类型ID"`
	Name   string `p:"name" v:"required" dc:"字典类型名称"`
	Type   string `p:"type" v:"required" dc:"字典类型标识（唯一）"`
	Sort   int    `p:"sort" dc:"字典排序"`
	Remark string `p:"remark" dc:"备注"`
	Status int    `p:"status" d:"1" v:"required|in:0,1#状态不能为空|状态只能为0或1" dc:"状态（1正常 0停用）"`
}

type CreateReq struct {
	g.Meta `path:"/dicttype" method:"post" tags:"DictType" summary:"新增字典类型" notes:"新增字典类型"`
	DictType
}
type CreateRes struct {
	g.Meta `mime:"application/json"`
}

type DeleteReq struct {
	g.Meta `path:"/dicttype/{id}" method:"delete" tags:"DictType" summary:"删除字典类型" notes:"删除字典类型"`
	Id     int64 `p:"id" v:"required" dc:"字典类型ID"`
}
type DeleteRes struct{}

type UpdateReq struct {
	g.Meta `path:"/dicttype/{id}" method:"put" tags:"DictType" summary:"更新字典类型" notes:"更新字典类型"`
	Id     int64 `p:"id" v:"required" dc:"字典类型ID"`
	DictType
}
type UpdateRes struct{}

type GetListReq struct {
	g.Meta `path:"/dicttype" method:"get" tags:"DictType" summary:"获取字典类型列表" notes:"获取字典类型列表"`
}
type GetListRes struct {
	List []*entity.PandaDictType `json:"list" dc:"字典类型列表"`
}
