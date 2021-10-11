package response

type LinkSortJson struct {
	Id          int         `json:"id"`
	Name        string      `json:"name"`
	Alias       string      `json:"alias"`
	Description string      `json:"description"`
	IsShow      bool        `json:"is_show"`
	Expand      interface{} `json:"expand"`
	CreateTime  string      `json:"create_time"`
	UpdateTime  string      `json:"update_time"`
}

// 全部友链分类拓展
type LinkSortsExpand struct {
	Count int64  `json:"count"`
	Color string `json:"color"`
}
