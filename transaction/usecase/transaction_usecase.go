package usecase

import (
	"encoding/json"
	"speakerseeker-backend/domain"
	"speakerseeker-backend/infrastructure"
	"time"
)

type TransactionUsecase struct {
	transactionRepository         domain.TransactionRepository
	midtransTransactionRepository domain.MidtransTransactionRepository
	speakerRepository             domain.SpeakerRepository
	midtransCoreClient            *infrastructure.CoreApi
}

func NewTransactionUsecase(tr domain.TransactionRepository, mtr domain.MidtransTransactionRepository, sr domain.SpeakerRepository, mcc *infrastructure.CoreApi) domain.TransactionUseCase {
	return &TransactionUsecase{transactionRepository: tr, midtransTransactionRepository: mtr, speakerRepository: sr, midtransCoreClient: mcc}
}

func (u *TransactionUsecase) Order(speakerId uint, userId uint, input domain.CreateTransaction) (uint, error) {
	speaker, err := u.speakerRepository.FindOne(speakerId)
	if err != nil {
		return 0, err
	}
	layoutFormat := "2006-01-02T15:04 MST"
	date, _ := time.Parse(layoutFormat, input.EventDate+" WIB")

	transaction := domain.Transaction{
		OrganizationName: input.OrganizationName,
		Phone:            input.Phone,
		Email:            input.Email,
		Type:             input.Type,
		Topic:            input.Topic,
		Goals:            input.Goals,
		EventName:        input.EventName,
		EventDate:        &date,
		Location:         input.Location,
		Address:          input.Address,
		Budget:           input.Budget,
		FilePath:         "",
		UserId:           userId,
		SpeakerId:        speakerId,
	}

	id, err := u.transactionRepository.Create(transaction)
	if err != nil {
		return id, err
	}

	coreApiRes, err := u.midtransCoreClient.CreateOrder(id, speaker.ID, speaker.Name, int64(input.Budget), input.OrganizationName, input.Email, input.PaymentType)
	if err != nil {
		return id, err
	}

	type paymentDataStruct struct {
		Key string `json:"key"`
		Qr  string `json:"qr"`
	}

	paymentData := &paymentDataStruct{}
	if input.PaymentType == "gopay" {
		paymentData.Key = coreApiRes.Actions[1].URL
		paymentData.Qr = coreApiRes.Actions[0].URL
	} else if input.PaymentType == "va-bni" {
		paymentData.Key = coreApiRes.VaNumbers[0].VANumber
	}

	paymentDataJson, err := json.Marshal(paymentData)
	if err != nil {
		return 0, err
	}

	mtTrx := domain.MidtransTransaction{
		MidtransId:    coreApiRes.TransactionID,
		PaymentType:   input.PaymentType,
		Amount:        input.Budget,
		Status:        "pending",
		PaymentData:   paymentDataJson,
		TransactionId: id,
	}

	err = u.midtransTransactionRepository.Create(mtTrx)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (u *TransactionUsecase) FindOne(id uint, userId uint) (domain.Transaction, error) {
	transaction, err := u.transactionRepository.GetById(id)
	if err != nil {
		return transaction, err
	}
	if userId != transaction.ID {
		return transaction, err
	}
	return transaction, nil
}

func (u *TransactionUsecase) FindByUserId(id uint) ([]domain.Transaction, error) {
	transaction, err := u.transactionRepository.GetAllByUserId(id)
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}
