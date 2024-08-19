package coreAPI

import (
	errz "github.com/mrido10/error"
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

func (p payment) GenerateTransaction(charge *coreapi.ChargeReq) (*coreapi.ChargeResponse, *errz.Error) {
	var mtErr *midtrans.Error
	resp, err := p.ChargeTransaction(charge)
	if err != nil && err != mtErr {
		return nil, errz.InternalServerErrorBadRequest(err.Message)
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
func (p payment) Notification(req map[string]interface{}) (bool, *errz.Error) {
	orderId, exists := req["order_id"].(string)
	if !exists {
		return false, errz.NotFound("order_id doesn't exists")
	}

	transactionStatusResp, err := p.CheckTransaction(orderId)
	if err != nil {
		return false, errz.InternalServerErrorBadRequest(err.Error())
	}

	if transactionStatusResp == nil {
		return false, errz.BadRequest("no response status")
	}

	switch transactionStatusResp.TransactionStatus {
	case "settlement":
		return true, nil
	case "capture":
		if transactionStatusResp.FraudStatus == "accept" {
			return true, nil
		} else {
			return false, errz.BadRequest(transactionStatusResp.FraudStatus)
		}
	default:
		return false, errz.BadRequest(transactionStatusResp.TransactionStatus)
	}
}
