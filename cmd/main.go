package main

import (
	"yet-another-restapi/internal/server"
)

const servePort string = ":8999"

func main() {

	l, _ := server.NewMultiLogger("data/log2")
	l.Debug("Ahaha1", "lol1")
	l.Info("Ahaha2", "lol2")
	l.Warn("Ahaha3", "lol3")


	//
	//s := &server.Server{
	//	server.DefaultModel(),
	//	&server.JsonPresenter{},
	//	server.NewFileLogger("data/log"),
	//	servePort,
	//}
	//s.Run()

}
