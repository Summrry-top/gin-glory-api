package response

type BannerJson struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	Img         string `json:"img"`
	CreateTime  string `json:"create_time"`
	UpdateTime  string `json:"update_time"`
}
