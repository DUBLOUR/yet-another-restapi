package server

import (
	"net/http"
)

type ILog interface {
	Debug(...interface{})
	Info(...interface{})
	Warn(...interface{})
}

type IPresenter interface {
	Format(interface{}) string
}

type IMoney interface{}

type IPayment interface{}

type IPaymentService interface{
	Name() string
	CreatePay(IMoney) (IPayment, error)
}

type IPaymentServiceRepository interface {
	Add(name string, service IPaymentService) error
	All() map[string]IPaymentService
	ByName(name string) (IPaymentService, bool)
}

type IModel interface {
	Price(string) (IMoney, error)
	CreatePay(money IMoney, serviceName string) (IPayment, error)
}

type Server struct {
	Model     IModel
	Presenter IPresenter
	Log       ILog
	Port      string
}

func (s *Server) Run() {
	h := Handlers(s.Model, s.Presenter, s.Log)
	server := &http.Server{
		Addr:    s.Port,
		Handler: h,
	}

	s.Log.Warn(server.ListenAndServe())
	//TODO: graceful shutdown
}

