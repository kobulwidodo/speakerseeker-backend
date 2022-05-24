package infrastructure

import (
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type CoreApi struct {
	ca coreapi.Client
}

func NewMidtransDriver() CoreApi {
	return CoreApi{ca: coreapi.Client{}}
}

func (c *CoreApi) CreateOrder(orderId uint, speakerId uint, speakerName string, price int64, organizationName string, email string, paymentType string) (*coreapi.ChargeResponse, error) {
	c.ca.New(os.Getenv("MIDTRANS_SERVER_KEY"), midtrans.Sandbox)
	chargeReq := &coreapi.ChargeReq{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(int(orderId)),
			GrossAmt: price,
		},
		Items: &[]midtrans.ItemDetails{
			midtrans.ItemDetails{
				ID:    strconv.Itoa(int(speakerId)),
				Price: price,
				Qty:   1,
				Name:  speakerName,
			},
		},
		CustomerDetails: &midtrans.CustomerDetails{
			FName: organizationName,
			Email: email,
		},
	}

	if paymentType == "gopay" {
		chargeReq.PaymentType = coreapi.PaymentTypeGopay
	} else if strings.HasPrefix(paymentType, "va-") {
		chargeReq.PaymentType = coreapi.PaymentTypeBankTransfer
		payment := strings.Split(paymentType, "-")
		if payment[1] == "bni" {
			chargeReq.BankTransfer = &coreapi.BankTransferDetails{
				Bank: midtrans.BankBni,
			}
		} else {
			return &coreapi.ChargeResponse{}, errors.New("payment type not found")
		}
	} else {
		return &coreapi.ChargeResponse{}, errors.New("payment type not found")
	}

	coreApiRes, _ := c.ca.ChargeTransaction(chargeReq)
	return coreApiRes, nil
}
