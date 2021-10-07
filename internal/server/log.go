package server

import (
	"fmt"
	"io"
	"log"
	"os"
)

type FileLogger struct {
	filename string
}

func NewFileLogger(filename string) *FileLogger {
	return &FileLogger{filename}
}

func (l FileLogger) Print(v ...interface{}) {
	f, err := os.OpenFile(l.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}

	defer func() {
		_ = f.Close()
	}()

	if _, err := f.WriteString(fmt.Sprint(v...)); err != nil {
		log.Println(err)
	}
}

func (l FileLogger) Debug(v ...interface{}) {
	l.Print("(--) ", fmt.Sprintln(v...))
}

func (l FileLogger) Info(v ...interface{}) {
	l.Print("(II) ", fmt.Sprintln(v...))
}

func (l FileLogger) Warn(v ...interface{}) {
	l.Print("(WW) ", fmt.Sprintln(v...))
}

type MultiLogger struct {
	files map[string]io.Writer
}

func NewMultiLogger(filename string) (*MultiLogger, error) {
	mainLogFile, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return &MultiLogger{}, err
	}

	return &MultiLogger{
		map[string]io.Writer{
			"main": mainLogFile,
		},
	}, nil
}


func (l MultiLogger) Print(out []io.Writer, v ...interface{}) {
	for _, flow := range out {
		_, err := fmt.Fprint(flow, v...)
		if err != nil {
			log.Println("(WW) Error while logging:", err)
		}
	}
}

func (l MultiLogger) Debug(v ...interface{}) {
	flows := []io.Writer{os.Stdout}
	l.Print(flows, "(--) ", fmt.Sprintln(v...))
}


func (l MultiLogger) Info(v ...interface{}) {
	flows := []io.Writer{
		os.Stdout,
		l.files["main"],
	}
	l.Print(flows, "(II) ", fmt.Sprintln(v...))
}


func (l MultiLogger) Warn(v ...interface{}) {
	flows := []io.Writer{
		os.Stdout,
		l.files["main"],
	}
	l.Print(flows, "(WW) ", fmt.Sprintln(v...))
}
