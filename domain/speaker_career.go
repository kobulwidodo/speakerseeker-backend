package domain

import (
	"gorm.io/gorm"
)

type SpeakerCareer struct {
	gorm.Model
	Title      string `json:"title"`
	MonthStart string `json:"month_start"`
	YearStart  uint16 `json:"year_start"`
	SpeakerId  uint   `json:"speaker_id"`
}
