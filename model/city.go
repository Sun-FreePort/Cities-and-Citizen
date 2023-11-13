package model

import (
	"time"
)

type CityModel struct {
	ID           uint    `gorm:"primarykey"`
	Name         string  `gorm:"not nul"`
	HumanCount   uint    `gorm:"not nul"`
	LandHas      uint    `gorm:"not nul"`
	LandHardBase float32 `gorm:"not nul"`

	PastureLevel float32 `gorm:"not nul"`
	FarmLevel    float32 `gorm:"not nul"`
	FishLevel    float32 `gorm:"not nul"`
	WoodLevel    float32 `gorm:"not nul"`
	StoneLevel   float32 `gorm:"not nul"`
	CopperLevel  float32 `gorm:"not nul"`
	IronLevel    float32 `gorm:"not nul"`
	SilverLevel  float32 `gorm:"not nul"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

type CityHistoryModel struct {
	ID        uint   `gorm:"primarykey"`
	CityId    uint   `gorm:"index;not nul"`
	Info      string `gorm:"not nul"`
	CreatedAt time.Time
}
