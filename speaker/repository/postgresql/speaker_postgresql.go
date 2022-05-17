package postgresql

import (
	"speakerseeker-backend/domain"

	"gorm.io/gorm"
)

type SpeakerRepository struct {
	db *gorm.DB
}

func NewSpeakerRepository(db *gorm.DB) domain.SpekaerRepository {
	return &SpeakerRepository{db}
}

func (r *SpeakerRepository) FindAll() ([]domain.Speaker, error) {
	var speakers []domain.Speaker
	if err := r.db.Find(&speakers).Error; err != nil {
		return speakers, err
	}
	return speakers, nil
}

func (r *SpeakerRepository) FindOne(id uint) (domain.Speaker, error) {
	var speaker domain.Speaker
	if err := r.db.Preload("SpeakerSkills").Where("id = ?", id).First(&speaker).Error; err != nil {
		return speaker, err
	}
	return speaker, nil
}
