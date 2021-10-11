package response

// 全部文章标签Json
type ArticleTagJson struct {
	Id         int         `json:"id"`
	Name       string      `json:"name"`
	IsShow     bool        `json:"is_show"`
	Expand     interface{} `json:"expand"`
	CreateTime string      `json:"create_time"`
	UpdateTime string      `json:"update_time"`
}

// 全部文章标签拓展
type ArticleTagsExpand struct {
	Count int64  `json:"count"`
	Color string `json:"color"`
}

//// 一条文章标签拓展
//type ArticleTagExpand struct {
//	Count    int             `json:"count"`
//	//Articles []*ArticlesJson `json:"articles"`
//
//}

//type ArticleTag struct {
//	Id   int    `json:"id"`
//	Name string `json:"name"`
//}
