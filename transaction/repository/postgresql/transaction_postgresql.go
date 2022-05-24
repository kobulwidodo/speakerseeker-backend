package postgresql

import (
	"speakerseeker-backend/domain"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) domain.TransactionRepository {
	return &TransactionRepository{db}
}

func (r *TransactionRepository) GetAllByUserId(userId uint) ([]domain.Transaction, error) {
	var transactions []domain.Transaction
	if err := r.db.Where("user_id = ?", userId).Preload("MidtransTransaction").Preload("Speaker").Find(&transactions).Error; err != nil {
		return transactions, err
	}
	return transactions, nil
}

func (r *TransactionRepository) Create(transaction domain.Transaction) (uint, error) {
	if err := r.db.Create(&transaction).Error; err != nil {
		return 0, err
	}
	return transaction.ID, nil
}

func (r *TransactionRepository) GetById(id uint) (domain.Transaction, error) {
	var transaction domain.Transaction
	if err := r.db.Where("id = ?", id).Preload("MidtransTransaction").Find(&transaction).Error; err != nil {
		return transaction, err
	}
	return transaction, nil
}
