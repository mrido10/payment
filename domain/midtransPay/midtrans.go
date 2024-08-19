package midtransPay

import (
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
	errz "github.com/mrido10/error"
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
	GenerateTransaction(*coreapi.ChargeReq) (*coreapi.ChargeResponse, *errz.Error)
	Notification(map[string]interface{}) (bool, *errz.Error)
}
