package orm

import "github.com/Summrry-top/gin-glory-api/models/response"

type LinkSort struct {
	ArticleSort
}

func (s LinkSort) GetLinkSortJson(expand interface{}) *response.LinkSortJson {
	return &response.LinkSortJson{
		Id:          s.Id,
		Name:        s.Name,
		Alias:       s.Alias,
		Description: s.Description,
		IsShow:      s.IsShow,
		Expand:      expand,
		CreateTime:  "",
		UpdateTime:  "",
	}
}

func (s LinkSort) GetLinkSortName() response.Name {
	return response.Name{
		Id:   s.Id,
		Name: s.Name,
	}
}
