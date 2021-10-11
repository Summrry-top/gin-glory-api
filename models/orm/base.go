package orm

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	Id        int            `gorm:"comment:主键;primaryKey;" form:"id" `
	CreatedAt time.Time      `gorm:"comment:创建时间;"`
	UpdatedAt time.Time      `gorm:"comment:更新时间;"`
	DeletedAt gorm.DeletedAt `gorm:"comment:软删除;index;"`
}
