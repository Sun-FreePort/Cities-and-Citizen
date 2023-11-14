package model

import (
	"gorm.io/gorm"
	"time"
)

type UserModel struct {
	ID               uint   `gorm:"primarykey"`
	Name             string `gorm:"index;not nul"`
	AvatarPath       string
	CityId           uint
	Age              uint // 天数
	Credit           float32
	Money            float32
	Energy           float32
	Hungry           float32
	Happy            float32
	Health           float32
	EquipHandId      uint    // 手持物品ID
	EquipBodyId      uint    // 上身物品ID
	EquipTrouserId   uint    // 下身物品ID
	SkillFarm        uint    // 农作
	SkillLumber      uint    // 伐木
	SkillMine        uint    // 采矿
	SkillHandwork    uint    // 手工
	SkillFight       uint    // 战斗
	SkillFarmExp     float32 // 农作经验
	SkillLumberExp   float32 // 伐木经验
	SkillMineExp     float32 // 采矿经验
	SkillHandworkExp float32 // 手工经验
	SkillFightExp    float32 // 战斗经验
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type UserLogModel struct {
	gorm.Model
	ID         uint   `gorm:"primarykey"`
	Name       string `gorm:"index;not nul"`
	AvatarPath string
	Action     string
	CreatedAt  time.Time
}
