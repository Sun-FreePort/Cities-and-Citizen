package handler

import "github.com/Sun-FreePort/Cities-and-Citizen/model"

type SquareInfoResp struct {
	City        model.CityModel
	SpeechList  []model.SpeechModel
	SpeechCount int16
}

type SpeechListReq struct {
	PageLimit int16
	PageNow   int16
}

type SpeechListResp struct {
	SpeechList []model.SpeechModel
}

type PublishSpeechReq struct {
	Info string
}

type PublishFeelReq struct {
	Id     uint
	IsGood bool
}
