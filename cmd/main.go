package main

import "yet-another-restapi/internal/server"

const servePort string = ":8999"

func main() {
	s := &server.Server{
		server.DefaultModel(),
		&server.JsonPresenter{},
		server.NewFileLogger("data/log"),
		servePort,
	}
	s.Run()

}
