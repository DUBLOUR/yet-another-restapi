package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"yet-another-restapi/pkg/thirdPartyPayments"
)

type Server struct {
	Model           IModel
	Presenter       IPresenter
	Log             ILog
	Port            string
	ShutdownWaiting time.Duration
}

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
	m := map[string]thirdPartyPayments.IPaymentGateway{
		"PayPal": thirdPartyPayments.DefaultPayPal(PayPalToken),
		"QIWI":   thirdPartyPayments.DefaultQiwi(QiwiToken),
		"VP":     thirdPartyPayments.DefaultVirtualPay(),
	}

	for name, gateway := range m {
		s.Log.Debug("Load", name)
		tmpG := gateway //Manually create entity for avoid problems with shallow copying
		err = s.Model.AddGateway(name, &tmpG)
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
			s.Log.Warn("Stop listening:", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	s.Log.Info("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), s.ShutdownWaiting)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		s.Log.Warn("Server forced to shutdown:", err)
	}

}
