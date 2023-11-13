package model

import (
	"time"
)

type BuildingModel struct {
	ID             uint    `gorm:"primarykey"`
	CityId         uint    `gorm:"not nul"`
	OwnerId        uint    `gorm:"index;not nul"`
	WorkIndex      uint    `gorm:"not nul"`
	BuildingCount  uint    `gorm:"not nul"`
	EmployeeCount  uint    `gorm:"not nul"`
	UnitProductPay float32 `gorm:"not nul"` // 生产一单位货物的薪酬
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type BuildWorkModel struct {
	ID        uint `gorm:"primarykey"`
	CityId    uint `gorm:"index;not nul"`
	OwnerId   uint `gorm:"not nul"`
	WorkIndex uint `gorm:"index;not nul"`
}
