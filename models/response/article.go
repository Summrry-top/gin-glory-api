package response

// 多条文章json数据 不包含文章内容
type ArticlesJson struct {
	Id            int           `json:"id"`
	Title         string        `json:"title"`
	Description   string        `json:"description"`
	Top           int           `json:"top"`
	IsShow        bool          `json:"is_show"`
	ImgSrc        string        `json:"img_src"`
	Views         int           `json:"views"`
	FontCount     int           `json:"font_count"`
	ArticleSortId []string      `json:"article_sort_id"`
	ArticleTagId  []string      `json:"article_tag_id"`
	UserId        int           `json:"user_id"`
	Expand        ArticleExpand `json:"expand"`
	CreateTime    string        `json:"create_time"`
	UpdateTime    string        `json:"update_time"`
}

// 一条文章json数据 文章详情页
type ArticleJson struct {
	Id            int           `json:"id"`
	Title         string        `json:"title"`
	Description   string        `json:"description"`
	Content       string        `json:"content"`
	Top           int           `json:"top"`
	IsShow        bool          `json:"is_show"`
	Views         int           `json:"views"`
	FontCount     int           `json:"font_count"`
	ArticleSortId []string      `json:"article_sort_id"`
	ArticleTagId  []string      `json:"article_tag_id"`
	UserId        int           `json:"user_id"`
	Expand        ArticleExpand `json:"expand"`
	CreateTime    string        `json:"create_time"`
	UpdateTime    string        `json:"update_time"`
}

type ArticleExpand struct {
	ArticleSort []Name        `json:"article_sort"`
	ArticleTag  []Name        `json:"article_tag"`
	Author      ArticleAuthor `json:"author"`
	Comments    int           `json:"comments"`
}

type ArticleAuthor struct {
	Nickname    string `json:"nickname"`
	HeadImg     string `json:"head_img"`
	Email       string `json:"email"`
	AddressUrl  string `json:"address_url"`
	Description string `json:"description"`
}

type ArticleFormat struct {
	Sort       []string
	Tag        []string
	Expand     ArticleExpand
	CreateTime string
	UpdateTime string
}
