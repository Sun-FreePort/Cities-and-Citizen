package model

import (
	"time"
)

type StorageGoodsModel struct {
	ID         uint `gorm:"primarykey"`
	CityId     uint `gorm:"index;not nul"`
	OwnerId    uint `gorm:"index;not nul"`
	GoodsIndex uint `gorm:"not nul"`
	Idle       uint `gorm:"not nul"`
	Occupy     uint `gorm:"not nul"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
