package domain

import "gorm.io/gorm"

type SpeakerExperience struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	SpeakerId   uint   `json:"speaker_id"`
}
