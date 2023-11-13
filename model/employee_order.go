package model

import (
	"time"
)

type EmployeeOrderModel struct {
	ID             uint    `gorm:"primarykey"`
	CityId         uint    `gorm:"index;not nul"`
	OwnerId        uint    `gorm:"index;not nul"`
	BuildingId     uint    `gorm:"index;not nul"`
	WorkIndex      uint    `gorm:"not nul"`
	WorkHour       uint    `gorm:"not nul"` // 需要工作时间
	UnitProductPay float32 `gorm:"not nul"` // 生产一单位货物的薪酬
	Has            uint    `gorm:"not nul"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type EmployeeContractModel struct {
	ID         uint `gorm:"primarykey"`
	CityId     uint `gorm:"not nul"`
	EmployerId uint `gorm:"not nul"`
	BuildingId uint `gorm:"not nul"`
	EmployeeId uint `gorm:"not nul"`

	WorkIndex      uint      `gorm:"not nul"`
	WorkHour       uint      `gorm:"not nul"` // 需要工作时间
	WorkNeed       uint      `gorm:"not nul"`
	UnitProductPay float32   `gorm:"not nul"` // 生产一单位货物的薪酬
	EndAt          time.Time // 合同结束时间

	CreatedAt time.Time
	UpdatedAt time.Time
}

type EmployeeContractLogModel struct {
	ID             uint      `gorm:"primarykey"`
	CityId         uint      `gorm:"not nul"`
	EmployerId     uint      `gorm:"not nul"`
	BuildingId     uint      `gorm:"not nul"`
	EmployeeId     uint      `gorm:"not nul"`
	WorkIndex      uint      `gorm:"not nul"`
	WorkHour       uint      `gorm:"not nul"` // 需要工作时间
	UnitProductPay float32   `gorm:"not nul"` // 生产一单位货物的薪酬
	EndAt          time.Time // 合同结束时间
	CreatedAt      time.Time
}
