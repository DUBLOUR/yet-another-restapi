package server

import (
	"fmt"
	"net/http"
)

func Handlers(model IModel, presenter IPresenter, logger ILog) http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var product string = r.URL.Query().Get("product_id")
		var service string = r.URL.Query().Get("service")
		logger.Info("Query:", r.URL.Query())

		if service == "" || product == "" {
			logger.Debug("Empty fields")
			logger.Info("Respond with status", http.StatusBadRequest)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		payment, err := model.CreateBill(product, service)
		if err != nil {
			logger.Debug(err)
			logger.Info("Respond with status", http.StatusInternalServerError)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		response := struct {
			Status string   `json:"status"`
			Href   IPayment `json:"href"`
		}{
			"ok",
			payment,
		}

		responseStr, _ := presenter.Format(response)

		logger.Debug(responseStr)
		logger.Info("Respond with status", http.StatusOK)

		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprintf(w, responseStr)
	})

	return r
}
