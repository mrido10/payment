package midtransPay

import (
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
)

type Config struct {
	IsProduction bool
	ServerKey    string
}

type Snap interface {
	GenerateURL(*snap.Request) (string, error)
	GenerateToken(*snap.Request) (string, error)
	GenerateTransaction(*snap.Request) (*snap.Response, error)
}

type CoreApi interface {
	GenerateTransaction(*coreapi.ChargeReq) (*coreapi.ChargeResponse, error)
	Notification(map[string]interface{}) (bool, error)
}
