package server

import (
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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
	//Price(string) (IMoney, error)
	//CreatePay(money IMoney, serviceName string) (IPayment, error)
	CreateBill(product string, serviceName string) (IPayment, error)
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
