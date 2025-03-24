package storage

// FileMeta 文件元数据
type FileMeta struct {
	Filename  string `json:"filename"`  // 文件名称
	Size      int64  `json:"size"`      // 文件大小
	Kind      string `json:"kind"`      // 文件上传类型
	MimeType  string `json:"mimeType"`  // 文件扩展类型
	NaiveType string `json:"naiveType"` // NaiveUI类型
	Ext       string `json:"ext"`       // 文件扩展名
	Md5       string `json:"md5"`       // 文件hash
}
