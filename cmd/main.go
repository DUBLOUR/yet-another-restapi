package main

import (
	"time"
	"yet-another-restapi/internal/server"
	"yet-another-restapi/pkg/advancedLog"
)

const logFile string = "data/log"
const servePort string = ":8999"

func main() {

	s := &server.Server{
		server.DefaultModel(),
		&server.JsonPresenter{},
		advancedLog.NewMultiLogger("data/log"),
		servePort,
		5 * time.Second,
	}
	s.Run()

}
