package response

type OptionJson struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Keywords    string `json:"keywords"`
	Copy        string `json:"copy"`
	SiteIco     string `json:"site_ico"`
	SiteImg     string `json:"site_img"`
	SiteUrl     string `json:"site_url"`
	CreateTime  string `json:"create_time"`
	UpdateTime  string `json:"update_time"`
}
