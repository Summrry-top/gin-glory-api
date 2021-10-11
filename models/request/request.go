package request

// get 参数
type GetParam struct {
	Search string `form:"search"`
	Id     int    `form:"id"`
	Limit  int    `form:"limit"`
	Page   int    `form:"page"`
	Order  string `form:"order"`
	Cache  bool   `form:"cache"`
}

// 权限参数
type Param struct {
	LoginToken string `form:"login-token"`
	Mode       string `form:"mode"`
}

// post users接口参数
type PostUserParam struct {
	Mode     string      `form:"mode" binding:"required"`
	Account  string      `form:"account"`
	Email    string      `form:"email"`
	Password string      `form:"password"`
	Code     string      `form:"code"`
	Id       interface{} `form:"id"`
}
