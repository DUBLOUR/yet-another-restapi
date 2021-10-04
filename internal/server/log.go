package server

import (
	"fmt"
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
