package domain

import "gorm.io/gorm"

type SpeakerSkill struct {
	gorm.Model
	Title     string `json:"title"`
	SpeakerId uint   `json:"speaker_id"`
}

type SpeakerSkillRepository interface {
	FindBySpeakerId(id uint) ([]SpeakerSkill, error)
}

type SpeakerSkillUsecase interface {
	GetBySpeakerId(id uint) ([]SpeakerSkill, error)
}

type SkillByIdUri struct {
	Id uint `uri:"id" binding:"required"`
}
