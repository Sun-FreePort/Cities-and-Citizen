package model

import (
	"time"
)

type MarketOrderModel struct {
	ID             uint    `gorm:"primarykey"`
	CityId         uint    `gorm:"index;not nul"`
	SellerId       uint    `gorm:"not nul"`
	SellGoodsIndex uint    `gorm:"index;not nul"`
	UnitPrice      float32 `gorm:"not nul"` // 一货物单位的价格
	Has            uint    `gorm:"not nul"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type MarketOrderLogModel struct {
	ID             uint    `gorm:"primarykey"`
	CityId         uint    `gorm:"not nul"`
	SellerId       uint    `gorm:"not nul"`
	SellGoodsIndex uint    `gorm:"index;not nul"`
	UnitPrice      float32 `gorm:"not nul"`
	Count          uint    `gorm:"not nul"`
	CreatedAt      time.Time
}
