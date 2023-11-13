package model

import "time"

type DictBuildingModel struct {
	Index             uint   `gorm:"primarykey"`
	Name              string `gorm:"not nul"`
	ProductGoodsIndex uint   `gorm:"not nul"`
	NeedGoodsIndex1   uint   `gorm:"not nul"` // 生产一件物品需要的素材
	NeedGoodsIndex2   uint   `gorm:"default:0"`
	NeedGoodsIndex3   uint   `gorm:"default:0"`
	NeedGoodsIndex4   uint   `gorm:"default:0"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type DictGoodsModel struct {
	Index     uint   `gorm:"primarykey"`
	Name      string `gorm:"not nul"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
