package server

import (
	"net/http"
)

func Handlers(model IModel, presenter IPresenter, logger ILog) http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		responder := InitResponder(w)
		defer responder.WriteResponse()
		responder.SetLogger(logger)
		responder.Set("Content-Type", presenter.ContentType())

		var product string = r.URL.Query().Get("product_id")
		var gateway string = r.URL.Query().Get("service")
		logger.Info("Query:", r.URL.Query())

		if gateway == "" || product == "" {
			responder.SetStatus(http.StatusBadRequest)
			answer, _ := presenter.Format(struct {
				Status string   `json:"status"`
				Error  string   `json:"error"`
				Href   IPayment `json:"href"`
			}{
				"fail",
				"`product_id` or `service` is empty",
				model.DefaultResponse(),
			})
			responder.SetBody(answer)
			return
		}

		payment, err := model.CreateBill(product, gateway)
		if err != nil {
			responder.SetStatus(http.StatusInternalServerError)
			answer, _ := presenter.Format(struct {
				Status string   `json:"status"`
				Error  string   `json:"error"`
				Href   IPayment `json:"href"`
			}{
				"fail",
				err.Error(),
				model.DefaultResponse(),
			})
			responder.SetBody(answer)
			return
		}

		answer, _ := presenter.Format(struct {
			Status string   `json:"status"`
			Href   IPayment `json:"href"`
		}{
			"ok",
			payment,
		})
		responder.SetBody(answer)
		responder.SetStatus(http.StatusOK)
	})

	return r
}
