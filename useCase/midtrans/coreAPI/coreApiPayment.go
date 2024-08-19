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
		return nil, errors.New(err.Error())
	}
	return resp, nil
}

/*
Notification is to checking status on midtrans by order_id

settlement:	succes
pending: 	waiting
cancel: 	fail
expire: 	fail
*/
func (p payment) Notification(req map[string]interface{}) (bool, error) {
	orderId, exists := req["order_id"].(string)
	if !exists {
		return false, errors.New("order_id doesn't exists")
	}

	transactionStatusResp, err := p.CheckTransaction(orderId)
	if err != nil {
		return false, errors.New(err.Error())
	}

	if transactionStatusResp == nil {
		return false, errors.New("no response status")
	}

	switch transactionStatusResp.TransactionStatus {
	case "settlement":
		return true, nil
	case "capture":
		if transactionStatusResp.FraudStatus == "accept" {
			return true, nil
		} else {
			return true, errors.New(transactionStatusResp.FraudStatus)
		}
	default:
		return false, errors.New(transactionStatusResp.TransactionStatus)
	}
}
