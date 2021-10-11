package response

type LinkJson struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	Url         string   `json:"url"`
	HeadImg     string   `json:"head_img"`
	Description string   `json:"description"`
	LinkSort    []string `json:"link_sort"`
	Expand      []Name   `json:"expand"`
	IsShow      bool     `json:"is_show"`
	CreateTime  string   `json:"create_time"`
	UpdateTime  string   `json:"update_time"`
}

type LinkFormat struct {
	Sort       []string
	Expand     []Name
	CreateTime string
	UpdateTime string
}
