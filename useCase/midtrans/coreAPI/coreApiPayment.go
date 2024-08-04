package coreAPI

import (
	"errors"
	"github.com/mrido10/payment/domain/midtransPay"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type payment struct {
	coreapi.Client
}

func New(m midtransPay.Config) midtransPay.CoreApi {
	env := midtrans.Sandbox
	if m.IsProduction {
		env = midtrans.Production
	}

	var client coreapi.Client
	client.New(m.ServerKey, env)

	return payment{
		Client: client,
	}
}

func (p payment) GenerateTransaction(charge *coreapi.ChargeReq) (*coreapi.ChargeResponse, error) {
	var mtErr *midtrans.Error
	resp, err := p.ChargeTransaction(charge)
	if err != nil && err != mtErr {
		return nil, err.RawError
	}
	return resp, nil
}

func (p payment) Notification(req map[string]interface{}) (bool, error) {
	orderId, exists := req["order_id"].(string)
	if !exists {
		return false, errors.New("payment: order_id doesn't exists")
	}

	var mtErr *midtrans.Error
	transactionStatusResp, err := p.CheckTransaction(orderId)
	if err != nil && err != mtErr {
		return false, err.RawError
	}

	if transactionStatusResp == nil {
		return false, errors.New("payment: no response status")
	}

	if transactionStatusResp.TransactionStatus == "settlement" {
		return true, nil
	}

	if transactionStatusResp.TransactionStatus == "capture" && transactionStatusResp.FraudStatus == "accept" {
		return true, nil
	}

	return false, nil
}
