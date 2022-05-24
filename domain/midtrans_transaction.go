package domain

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type MidtransTransaction struct {
	gorm.Model
	MidtransId    string         `json:"midtrans_id"`
	PaymentType   string         `json:"payment_type"`
	Amount        int            `json:"amount"`
	Status        string         `json:"status"`
	PaymentData   datatypes.JSON `json:"payment_data"`
	TransactionId uint           `json:"transaction_id"`
}

type MidtransTransactionRepository interface {
	Create(midtransTrx MidtransTransaction) error
	GetById(id uint) (MidtransTransaction, error)
	GetByTrxId(id uint) (MidtransTransaction, error)
	Update(mtrx MidtransTransaction) error
}

type MidtransTransactionUsecase interface {
	Handler(id string) error
}
