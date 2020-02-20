package logger

import (
	"io"
	"log"
	"os"
)

// Via https://www.ardanlabs.com/blog/2013/11/using-log-package-in-go.html
var (
	Info  *log.Logger
	Error *log.Logger
)

func initWriter(
	infoHandle io.Writer,
	errorHandle io.Writer) {

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func InitLog() {
	bpfLog, err := os.OpenFile("./checkbpf.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	initWriter(bpfLog, bpfLog)
}
