package snap

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"github.com/mrido10/payment/domain/midtransPay"
)

type payment struct {
	snap.Client
}

func New(m midtransPay.Config) payment {
	env := midtrans.Sandbox
	if m.IsProduction {
		env = midtrans.Production
	}
	var client snap.Client
	client.New(m.ServerKey, env)

	return payment{
		Client: client,
	}
}

func (p payment) GenerateURL(req *snap.Request) (string, error) {
	return p.CreateTransactionUrl(req)
}

func (p payment) GenerateToken(req *snap.Request) (string, error) {
	return p.CreateTransactionToken(req)
}

func (p payment) GenerateTransaction(req *snap.Request) (*snap.Response, error) {
	return p.CreateTransaction(req)
}
