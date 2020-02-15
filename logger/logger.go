package logger

import (
	"fmt"
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
	bpfLog, err := os.OpenFile("/var/log/checkbpf.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		fmt.Println("Unable to write logfile in /var/log", err)
		home, userHomeErr := os.UserHomeDir()
		if userHomeErr != nil {
			log.Fatal("Unable to determine user's dir", userHomeErr)
		}
		bpfLog2, err2 := os.Create(home + "/bpfcheck.log")
		if err2 != nil {
			log.Fatal("Unable to write logfile in ~", err2)
		}
		bpfLog = bpfLog2
	}
	initWriter(bpfLog, bpfLog)
}
