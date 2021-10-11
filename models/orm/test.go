package orm

import "github.com/lib/pq"

type Test struct {
	Id   int
	Sort []ArticleSort `gorm:"foreignKey:Alias"`
	Tag  pq.Int64Array `gorm:"type:string"`
}
