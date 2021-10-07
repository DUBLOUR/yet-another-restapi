package server

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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
	AddService(string, *thirdPartyPayments.IPaymentService) error
	CreateBill(product string, serviceName string) (IPayment, error)
}

type Server struct {
	Model     IModel
	Presenter IPresenter
	Log       ILog
	Port      string
}

const ProductFilename string = "data/products.json"
const PayPalToken string = "-"
const QiwiToken string = "-"

func (s *Server) Init() error {
	s.Log.Info("Init server...")
	s.Log.Debug("Load file", ProductFilename)
	err := s.Model.LoadProducts(ProductFilename)
	if err != nil {
		s.Log.Warn(err)
		return fmt.Errorf("cannot load products")
	}
	s.Log.Debug("Successfully loaded all products from", ProductFilename)

	s.Log.Debug("Load payment services...")
	m := map[string]thirdPartyPayments.IPaymentService{
		"PayPal": thirdPartyPayments.DefaultPayPal(PayPalToken),
		"QIWI":   thirdPartyPayments.DefaultQiwi(QiwiToken),
		"VP":     thirdPartyPayments.DefaultVirtualPay(),
	}

	for name, service := range m {
		s.Log.Debug("Load", name)
		tmpS := service //Manually create entity for avoid problems with shallow copying
		err = s.Model.AddService(name, &tmpS)
		if err != nil {
			s.Log.Warn(err)
			return fmt.Errorf("cannot add payment service")
		}
	}

	s.Log.Info("Initialize is successful")
	return nil
}

func (s *Server) Run() {
	if err := s.Init(); err != nil {
		s.Log.Warn(err)
		panic("cannot init server")
	}

	h := Handlers(s.Model, s.Presenter, s.Log)
	server := &http.Server{
		Addr:    s.Port,
		Handler: h,
	}

	s.Log.Info("Start server at port", s.Port)

	go func() {
		if err := server.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			s.Log.Warn("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	s.Log.Info("Shutting down server...")

	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	//
	//if err := server.Shutdown(ctx); err != nil {
	//	log.Fatal("Server forced to shutdown:", err)
	//}

}
