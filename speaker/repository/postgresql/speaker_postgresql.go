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

func (r *SpeakerRepository) FindAll(query string) ([]domain.Speaker, error) {
	var speakers []domain.Speaker
	queryFix := "%" + query + "%"
	if err := r.db.Preload("SpeakerSkills").Where("name ILIKE ?", queryFix).Find(&speakers).Error; err != nil {
		return speakers, err
	}
	return speakers, nil
}

func (r *SpeakerRepository) FindOne(id uint) (domain.Speaker, error) {
	var speaker domain.Speaker
	if err := r.db.Preload("SpeakerSkills").Preload("SpeakerCareers").Preload("SpeakerExperiences").Where("id = ?", id).First(&speaker).Error; err != nil {
		return speaker, err
	}
	return speaker, nil
}
