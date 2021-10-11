package response

// 定义返回消息体
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// 登录成功返回的数据
type LoginData struct {
	LoginToken string      `json:"login_token"`
	User       interface{} `json:"user"`
}

// 分页数据
type PageData struct {
	TotalPage   int64       `json:"totalPage"`
	CurrentPage int         `json:"currentPage"`
	PageSize    int         `json:"pageSize"`
	Count       int64       `json:"count"`
	Data        interface{} `json:"data"`
}

// 名称 用于拓展 获取标签,文章分类或友链分类名称
type Name struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
