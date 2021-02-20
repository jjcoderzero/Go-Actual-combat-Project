package global

import (
	"gorm.io/gorm"
	"time"
)

type CRC_MODEL struct {
	ID       uint `gorm:"primarykey"`
	CreateAt time.Time
	UpdateAt time.Time
	DeleteAt gorm.DeletedAt `gorm:"index" json:"-"`
}
