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
	midtrans.DefaultLoggerLevel = &midtrans.LoggerImplementation{LogLevel: midtrans.LogError}
	QrisTransaction(midtrans.TransactionDetails{
		OrderID:  "ORDER-test1",
		GrossAmt: 10000,
	})
}
