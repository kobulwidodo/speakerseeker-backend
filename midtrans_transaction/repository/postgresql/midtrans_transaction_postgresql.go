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

func (r *MidtransTransactionRepository) GetById(id uint) (domain.MidtransTransaction, error) {
	var mTrx domain.MidtransTransaction
	if err := r.db.Where("id = ?", id).Find(&mTrx).Error; err != nil {
		return mTrx, err
	}
	return mTrx, nil
}

func (r *MidtransTransactionRepository) Update(mtrx domain.MidtransTransaction) error {
	if err := r.db.Save(&mtrx).Error; err != nil {
		return err
	}
	return nil
}
