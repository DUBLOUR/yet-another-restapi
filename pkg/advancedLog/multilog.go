package advancedLog

import (
	"fmt"
	"io"
	"log"
	"os"
)

type MultiLogger struct {
	files map[string]io.Writer
}

func NewMultiLogger(filename string) *MultiLogger {
	mainLogFile, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	return &MultiLogger{
		map[string]io.Writer{
			"main": mainLogFile,
		},
	}
}

func (l MultiLogger) print(out []io.Writer, v ...interface{}) {
	for _, flow := range out {
		_, err := fmt.Fprint(flow, v...)
		if err != nil {
			log.Println("(WW) Error while logging:", err)
		}
	}
}

func (l MultiLogger) Debug(v ...interface{}) {
	flows := []io.Writer{os.Stdout}
	l.print(flows, "(--) ", fmt.Sprintln(v...))
}

func (l MultiLogger) Info(v ...interface{}) {
	flows := []io.Writer{
		os.Stdout,
		l.files["main"],
	}
	l.print(flows, "(II) ", fmt.Sprintln(v...))
}

func (l MultiLogger) Warn(v ...interface{}) {
	flows := []io.Writer{
		os.Stdout,
		l.files["main"],
	}
	l.print(flows, "(WW) ", fmt.Sprintln(v...))
}
