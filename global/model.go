package global

import (
	"gorm.io/gorm"
	"time"
)

type GRA_MODEL struct {
	ID       uint           `gorm:"primarykey"` // 主键ID
	CreateAt time.Time      // 创建时间
	UpdateAt time.Time      // 更新时间
	DeleteAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}
