package usecase

import (
	"speakerseeker-backend/domain"
	"speakerseeker-backend/infrastructure"
	"strconv"
)

type MidtransTransactionUsecase struct {
	midtransTransactionRepository domain.MidtransTransactionRepository
	midtransCoreClient            *infrastructure.CoreApi
}

func NewMidtransTransactionUsecase(mtr domain.MidtransTransactionRepository, mcc *infrastructure.CoreApi) domain.MidtransTransactionUsecase {
	return &MidtransTransactionUsecase{midtransTransactionRepository: mtr, midtransCoreClient: mcc}
}

func (u *MidtransTransactionUsecase) Handler(id string) error {
	midtransReport, err := u.midtransCoreClient.HandleNotification(id)
	if err != nil {
		return err
	}
	idInt, _ := strconv.Atoi(id)
	mTrx, err := u.midtransTransactionRepository.GetById(uint(idInt))
	if err != nil {
		return err
	}
	if midtransReport != nil {
		if midtransReport.TransactionStatus == "capture" {
			if midtransReport.FraudStatus == "challenge" {
				mTrx.Status = "challange"
			} else if midtransReport.FraudStatus == "accept" {
				mTrx.Status = "success"
			}
		} else if midtransReport.TransactionStatus == "settlement" {
			mTrx.Status = "success"
		} else if midtransReport.TransactionStatus == "deny" {
			mTrx.Status = "deny"
		} else if midtransReport.TransactionStatus == "cancel" || midtransReport.TransactionStatus == "expire" {
			mTrx.Status = "failure"
		} else if midtransReport.TransactionStatus == "pending" {
			mTrx.Status = "pending"
		}
	}

	err = u.midtransTransactionRepository.Update(mTrx)
	if err != nil {
		return err
	}
	return nil
}
