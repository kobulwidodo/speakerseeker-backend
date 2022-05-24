package postgresql

import (
	"speakerseeker-backend/domain"

	"gorm.io/gorm"
)

type MidtransTransactionRepository struct {
	db *gorm.DB
}

func NewMidtransTransactionRepository(db *gorm.DB) domain.MidtransTransactionRepository {
	return &MidtransTransactionRepository{db}
}

func (r *MidtransTransactionRepository) Create(midtransTrx domain.MidtransTransaction) error {
	if err := r.db.Create(&midtransTrx).Error; err != nil {
		return err
	}
	return nil
}
