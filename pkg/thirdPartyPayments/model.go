package thirdPartyPayments

type IMoney interface{}

type IPayment interface{}

type IPaymentGateway interface {
	//Name() string
	CreatePay(IMoney) (IPayment, error)
}
