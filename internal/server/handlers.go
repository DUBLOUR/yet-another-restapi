package server

import (
	"fmt"
	"log"
	"net/http"
)

func Handlers(model IModel, presenter IPresenter, logger ILog) http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var product string = r.URL.Query().Get("product_id")
		var service string = r.URL.Query().Get("service")
		logger.Info(r.URL.Query())

		if service == "" || product == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		pCost, err := model.Price(product)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		payment, err := model.CreatePay(pCost, service)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		response := struct{
			status string
			href IPayment
		}{
			status: "ok",
			href: payment,
		}

		responseStr := presenter.Format(response)

		log.Println(responseStr)
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprintf(w, responseStr)
	})

	return r
}

