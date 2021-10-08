package server

import (
	"fmt"
	"net/http"
	"time"
	"yet-another-restapi/pkg/advancedLog"
)

type Responder struct {
	w       http.ResponseWriter
	startAt time.Time
	status  int
	values  map[string]string
	body    string
	log     ILog
}

func InitResponder(w http.ResponseWriter) *Responder {
	return &Responder{
		w,
		time.Now(),
		http.StatusOK,
		map[string]string{
			"X-Server-Name": Hostname,
		},
		"",
		advancedLog.NewEmptyLogger(),
	}
}

func (r *Responder) SetLogger(l ILog) {
	r.log = l
}

func (r *Responder) Set(key string, val interface{}) {
	r.values[key] = fmt.Sprintf("%v", val)
}

func (r *Responder) SetStatus(code int) {
	r.status = code
}

func (r *Responder) SetBody(body string) {
	r.body = body
	r.Set("Content-Length", len(body))
}

func (r *Responder) postHandling() {
	responseTime := time.Since(r.startAt).Microseconds()
	r.Set("X-Response-Time", fmt.Sprintf("%vms", responseTime))
}

func (r *Responder) WriteResponse() {
	r.postHandling()
	for key, val := range r.values {
		r.w.Header().Set(key, val)
	}
	r.w.WriteHeader(r.status)
	_, err := fmt.Fprintf(r.w, r.body)

	if err != nil {
		r.log.Warn("Unsuccessful respond:", err)
	}
	r.log.Info("Respond with status", r.status)
	r.log.Debug(r.body)
}
