package model

import "time"

type SpeechModel struct {
	ID        uint   `gorm:"primarykey"`
	CityId    uint   `gorm:"not nul"`
	SpeakerId uint   `gorm:"not nul"`
	Info      string `gorm:"not nul"`
	Like      uint   `gorm:"not nul"`
	Boo       uint   `gorm:"not nul"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SpeechLogModel struct {
	ID        uint `gorm:"primarykey"`
	SpeakerId uint `gorm:"not nul"`
	VoterId   uint `gorm:"not nul"`
	IsGood    bool `gorm:"not nul"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
