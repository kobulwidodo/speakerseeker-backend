package postgresql

import (
	"speakerseeker-backend/domain"

	"gorm.io/gorm"
)

type RatingRepository struct {
	db *gorm.DB
}

func NewRatingRepository(db *gorm.DB) domain.RatingRepository {
	return &RatingRepository{db}
}

func (r *RatingRepository) Create(rating domain.Rating) error {
	if err := r.db.Create(&rating).Error; err != nil {
		return err
	}
	return nil
}
