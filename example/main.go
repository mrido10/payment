package main

import (
	"github.com/midtrans/midtrans-go"
	"github.com/mrido10/payment/domain/midtransPay"
	"os"
)

var config = midtransPay.Config{
	IsProduction: false,
	ServerKey:    os.Getenv("MIDTRANS_SERVER_KEY"),
}

func main() {
	QrisTransaction(midtrans.TransactionDetails{
		OrderID:  "ORDER-test4",
		GrossAmt: 10000,
	})
}
