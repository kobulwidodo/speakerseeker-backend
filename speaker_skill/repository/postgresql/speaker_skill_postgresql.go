package postgresql

import (
	"speakerseeker-backend/domain"

	"gorm.io/gorm"
)

type SpeakerSkillRepository struct {
	db *gorm.DB
}

func NewSpeakerSkillRepository(db *gorm.DB) domain.SpeakerSkillRepository {
	return &SpeakerSkillRepository{db}
}

func (r *SpeakerSkillRepository) FindBySpeakerId(id uint) ([]domain.SpeakerSkill, error) {
	var skills []domain.SpeakerSkill
	if err := r.db.Where("speaker_id = ?", id).Find(&skills).Error; err != nil {
		return skills, err
	}
	return skills, nil
}

func (r *SpeakerSkillRepository) FindPluck(id uint) ([]string, error) {
	var skills []string
	if err := r.db.Where("speaker_id = ?", id).Pluck("title", &skills).Error; err != nil {
		return skills, err
	}
	return skills, nil
}
