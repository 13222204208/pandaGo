package common

// PageReq 公共请求参数
type PageReq struct {
	CurrentPage int    `p:"currentPage" d:"1" dc:"当前页码"` //当前页码
	PageSize    int    `p:"pageSize" d:"10" dc:"页大小"`    //每页数
	OrderBy     string //排序方式
}

type Author struct {
	Authorization string `p:"Authorization" in:"header" dc:"Bearer {{token}}"`
}
