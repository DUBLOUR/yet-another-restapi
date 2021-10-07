package thirdPartyPayments

type IMoney interface{}

type IPayment interface{}

type IPaymentService interface {
	//Name() string
	CreatePay(IMoney) (IPayment, error)
}
