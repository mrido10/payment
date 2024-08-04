package main

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	coreAPIPayment "github.com/mrido10/payment/useCase/midtrans/coreAPI"
	"log"
)

func QrisTransaction(transaction midtrans.TransactionDetails) {
	req := coreapi.ChargeReq{
		PaymentType:        coreapi.PaymentTypeQris,
		TransactionDetails: transaction,
		Qris: &coreapi.QrisDetails{
			Acquirer: "gopay",
		},
		CustomerDetails: &midtrans.CustomerDetails{
			Email: "testing@mail.mail.com",
		},
	}
	coreAPI, err := coreAPIPayment.New(config).GenerateTransaction(&req)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println(coreAPI.Actions[0].URL)
}
