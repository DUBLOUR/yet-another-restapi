package server

import (
	"yet-another-restapi/pkg/thirdPartyPayments"
)

type ILog interface {
	Debug(...interface{})
	Info(...interface{})
	Warn(...interface{})
}

type IPresenter interface {
	Format(interface{}) (string, error)
}

type IPayment interface{}

type IModel interface {
	LoadProducts(string) error
	AddGateway(string, *thirdPartyPayments.IPaymentGateway) error
	CreateBill(product string, gatewayName string) (IPayment, error)
}
