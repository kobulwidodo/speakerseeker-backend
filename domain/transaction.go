package domain

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	OrganizationName    string     `json:"organization_name"`
	Phone               string     `json:"phone"`
	Email               string     `json:"email"`
	Type                string     `json:"type"`
	Topic               string     `json:"topic"`
	Goals               string     `json:"goals"`
	EventName           string     `json:"event_name"`
	EventDate           *time.Time `json:"event_date"`
	Location            string     `json:"location"`
	Address             string     `json:"address"`
	Budget              int        `json:"budget"`
	FilePath            string     `json:"file_path"`
	UserId              uint       `json:"user_id"`
	SpeakerId           uint       `json:"speaker_id"`
	Speaker             Speaker`json:"speaker"`
	MidtransTransaction MidtransTransaction `json:"midtrans_transaction"`
}

type TransactionRepository interface {
	GetAllByUserId(userId uint) ([]Transaction, error)
	Create(transaction Transaction) (uint, error)
	GetById(id uint) (Transaction, error)
}

type TransactionUseCase interface {
	Order(speakerId uint, userId uint, input CreateTransaction) (uint, error)
	FindOne(id uint, userId uint) (Transaction, error)
	FindByUserId(id uint) ([]Transaction, error)
}

type CreateTransaction struct {
	OrganizationName string `json:"organization_name" binding:"required"`
	Phone            string `json:"phone"  binding:"required"`
	Email            string `json:"email"  binding:"required"`
	Type             string `json:"type"  binding:"required"`
	Topic            string `json:"topic"  binding:"required"`
	Goals            string `json:"goals"  binding:"required"`
	EventName        string `json:"event_name"  binding:"required"`
	EventDate        string `json:"event_date"  binding:"required"`
	Location         string `json:"location"`
	Address          string `json:"address"  binding:"required"`
	Budget           int    `json:"budget"  binding:"required"`
	PaymentType      string `json:"payment_type" binding:"required"`
}

type TransactionSpeakerIdUri struct {
	SpeakerId uint `uri:"speaker_id" binding:"required"`
}

type TransactionIdUri struct {
	Id uint `uri:"id" binding:"required"`
}
