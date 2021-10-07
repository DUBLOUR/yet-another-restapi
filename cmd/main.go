package main

import "yet-another-restapi/internal/server"

const logFile string = "data/log"
const servePort string = ":8999"

type xxx struct {
	m map[int]int
}

func (x xxx) Add(a, b int) {
	x.m[a] = b
}

func (x xxx) Get(a int) int {
	return x.m[a]
}

func main() {
	//
	//s := &xxx{}
	//s.m = make(map[int]int)
	//
	//s.Add(1,100)
	//s.Add(2,200)
	//fmt.Println(s.Get(2))

	s := &server.Server{
		server.DefaultModel(),
		&server.JsonPresenter{},
		//server.NewFileLogger("data/log"),
		server.NewMultiLogger("data/log"),
		servePort,
	}
	s.Run()

}
