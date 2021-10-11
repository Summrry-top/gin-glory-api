package response

// 文章分类Json
type ArticleSortJson struct {
	Id          int         `json:"id"`
	Name        string      `json:"name"`
	Alias       string      `json:"alias"`
	Description string      `json:"description"`
	IsShow      bool        `json:"is_show"`
	Expand      interface{} `json:"expand"`
	CreateTime  string      `json:"create_time"`
	UpdateTime  string      `json:"update_time"`
}

// 多条文章分类拓展 获取一个分类下的文章数量
type ArticleSortsExpand struct {
	Count int64 `json:"count"`
}

// 文章分类数据格式化 时间 => 2006-01-02 15:04:05
type ArticleSortFormat struct {
	CreateTime string
	UpdateTime string
}
