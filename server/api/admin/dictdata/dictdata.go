package dictdata

import (
	"server/api/common"
	"server/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type DictData struct {
	Label     string `p:"label" v:"required" dc:"字典标签"`
	Value     string `p:"value" v:"required" dc:"字典键值"`
	ValueType string `p:"valueType" v:"required" dc:"键值数据类型：string,int,uint,bool,datetime,date"`
	Type      string `p:"type" v:"required" dc:"字典数据标识（唯一）"`
	ListClass string `p:"listClass" dc:"列表渲染样式"`
	IsDefault int    `p:"isDefault" v:"required|in:0,1#状态不能为空|状态只能为0或1" dc:"是否默认（0否 1是）"`
	Sort      int    `p:"sort" dc:"字典排序"`
	Remark    string `p:"remark" dc:"备注"`
	Status    int    `p:"status" v:"required|in:0,1#状态不能为空|状态只能为0或1" dc:"状态（1正常 0停用）"`
}

type CreateReq struct {
	g.Meta `path:"/dictdata" method:"post" tags:"dictdata" summary:"新增字典数据" notes:"新增字典数据"`
	DictData
}
type CreateRes struct {
	g.Meta `mime:"application/json"`
}

type DeleteReq struct {
	g.Meta `path:"/dictdata/{id}" method:"delete" tags:"dictdata" summary:"删除字典数据" notes:"删除字典数据"`
	Id     int64 `p:"id" v:"required" dc:"字典数据ID"`
}
type DeleteRes struct{}

type UpdateReq struct {
	g.Meta `path:"/dictdata/{id}" method:"put" tags:"dictdata" summary:"更新字典数据" notes:"更新字典数据"`
	Id     int64 `p:"id" v:"required" dc:"字典数据ID"`
	DictData
}
type UpdateRes struct{}

type GetListReq struct {
	g.Meta `path:"/dictdata" method:"get" tags:"dictdata" summary:"获取字典数据列表" notes:"获取字典数据列表"`
	Label  string `p:"label" dc:"字典标签"`
	Value  string `p:"value" dc:"字典键值"`
	Type   string `p:"type" dc:"字典类型"`
	Status int    `p:"status" d:"-1" v:"in:-1,0,1" dc:"字典数据状态"`
	common.PageReq
}
type GetListRes struct {
	List []*entity.PandaDictData `json:"list"`
	common.ListRes
}
